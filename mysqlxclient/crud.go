package mysqlxclient

import (
	"fmt"
	"strings"
	"time"

	"github.com/ldeng7/go-mysqlx-client/mysqlxclient/internal/datatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxclient/mysqldatatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_crud"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_datatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_expr"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_notice"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_resultset"
)

func mysqlScalarFromColumnData(x interface{}) *mysqlxpb_datatypes.Scalar {
	switch y := x.(type) {
	case uint8: // BIT
		return datatypes.MysqlScalarFromUint64(uint64(y))
	case map[string]struct{}: // SET
		ar := make([]string, 0, len(y))
		for k, _ := range y {
			ar = append(ar, k)
		}
		return datatypes.MysqlScalarFromString(strings.Join(ar, ","))
	case *mysqldatatypes.Time: // TIME
		return datatypes.MysqlScalarFromString(y.String())
	case *time.Time: // DATE, DATETIME, TIMESTAMP
		return datatypes.MysqlScalarFromString(y.Format("2006-01-02 15:04:05"))
	case *mysqldatatypes.GenericGeometry: // GEOMETRY
		return datatypes.MysqlScalarFromBytes(y.Encode())
	default:
		return datatypes.MysqlScalarFromBasic(x)
	}
}

func (o *Order) toMsg() *mysqlxpb_crud.Order {
	msg := &mysqlxpb_crud.Order{
		Expr: o.By.toMsg(),
	}
	dir := mysqlxpb_crud.Order_ASC
	if !o.Asc {
		dir = mysqlxpb_crud.Order_DESC
	}
	msg.Direction = &dir
	return msg
}

func (c *Criteria) toMsg() (*mysqlxpb_expr.Expr, []*mysqlxpb_crud.Order, *mysqlxpb_crud.Limit) {
	orders := make([]*mysqlxpb_crud.Order, len(c.Orders))
	for i, o := range c.Orders {
		orders[i] = o.toMsg()
	}

	var limitOffset *mysqlxpb_crud.Limit
	if 0 != c.Limit || 0 != c.Offset {
		limitOffset = &mysqlxpb_crud.Limit{
			RowCount: &c.Limit,
			Offset:   &c.Offset,
		}
	}

	return c.Where.toMsg(), orders, limitOffset
}

func (fsi *FindSelectItem) toMsg() *mysqlxpb_crud.Projection {
	msg := &mysqlxpb_crud.Projection{
		Source: fsi.Expr.toMsg(),
	}
	if 0 != len(fsi.As) {
		msg.Alias = &fsi.As
	}
	return msg
}

func (ia *InsertArgs) toMsg() *mysqlxpb_crud.Insert {
	dataModel := mysqlxpb_crud.DataModel_TABLE
	nCols := len(ia.Columns)
	columns := make([]*mysqlxpb_crud.Column, nCols)
	for i, c := range ia.Columns {
		c := c
		columns[i] = &mysqlxpb_crud.Column{
			Name: &c,
		}
	}

	rows := make([]*mysqlxpb_crud.Insert_TypedRow, len(ia.Values))
	exprType := mysqlxpb_expr.Expr_LITERAL
	for i, ar := range ia.Values {
		if len(ar) != nCols {
			panic(fmt.Sprintf("len(values[%d]) != len(columns)", i))
		}
		fields := make([]*mysqlxpb_expr.Expr, nCols)
		for k, v := range ar {
			fields[k] = &mysqlxpb_expr.Expr{
				Type:    &exprType,
				Literal: mysqlScalarFromColumnData(v),
			}
		}
		rows[i] = &mysqlxpb_crud.Insert_TypedRow{
			Field: fields,
		}
	}

	msg := &mysqlxpb_crud.Insert{
		Collection: &mysqlxpb_crud.Collection{
			Name: &ia.TableName,
		},
		DataModel:  &dataModel,
		Projection: columns,
		Row:        rows,
	}
	return msg
}

func (pc *poolConn) insert(ia *InsertArgs) (uint64, uint64, error) {
	if err := pc.m.sendMsg(mysqlxpb.ClientMessages_CRUD_INSERT, ia.toMsg()); nil != err {
		return 0, 0, err
	}

	var rows, id uint64
	for level := 0; level < 3; {
		msg, t, err := pc.m.recvMsgUntilTypes(
			mysqlxpb.ServerMessages_NOTICE,
			mysqlxpb.ServerMessages_SQL_STMT_EXECUTE_OK,
		)
		if nil != err {
			return 0, 0, err
		}
		switch t {
		case mysqlxpb.ServerMessages_SQL_STMT_EXECUTE_OK:
			level++
		case mysqlxpb.ServerMessages_NOTICE:
			n, err := pc.m.parseSessionStateChangedNotice(msg)
			if nil != err {
				return 0, 0, err
			} else if nil == n {
				continue
			}

			var param *uint64
			switch n.GetParam() {
			case mysqlxpb_notice.SessionStateChanged_ROWS_AFFECTED:
				param = &rows
			case mysqlxpb_notice.SessionStateChanged_GENERATED_INSERT_ID:
				param = &id
			default:
				continue
			}
			level++
			if *param, err = pc.m.uint64ValueFromSessionStateChangedNotice(n); nil != err {
				return 0, 0, err
			}
		}
	}

	return rows, id, nil
}

func (fa *FindArgs) toMsg() *mysqlxpb_crud.Find {
	dataModel := mysqlxpb_crud.DataModel_TABLE

	selects := make([]*mysqlxpb_crud.Projection, len(fa.Select))
	for i, s := range fa.Select {
		selects[i] = s.toMsg()
	}

	groups := make([]*mysqlxpb_expr.Expr, len(fa.Groups))
	for i, g := range fa.Groups {
		groups[i] = g.toMsg()
	}

	where, orders, limitOffset := fa.Criteria.toMsg()

	msg := &mysqlxpb_crud.Find{
		Collection: &mysqlxpb_crud.Collection{
			Name: &fa.TableName,
		},
		DataModel:        &dataModel,
		Projection:       selects,
		Criteria:         where,
		Grouping:         groups,
		GroupingCriteria: fa.Having.toMsg(),
		Order:            orders,
		Limit:            limitOffset,
	}
	return msg
}

func (pc *poolConn) find(fa *FindArgs) (*FindResultSet, error) {
	if err := pc.m.sendMsg(mysqlxpb.ClientMessages_CRUD_FIND, fa.toMsg()); nil != err {
		return nil, err
	}

	ms := make([]*mysqlxpb_resultset.ColumnMetaData, 0, len(fa.Select))
	rs := make([]*mysqlxpb_resultset.Row, 0, 8)
loop:
	for {
		msg, t, err := pc.m.recvMsgUntilTypes(
			mysqlxpb.ServerMessages_RESULTSET_COLUMN_META_DATA,
			mysqlxpb.ServerMessages_RESULTSET_ROW,
			mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE,
		)
		if nil != err {
			return nil, err
		}
		switch t {
		case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE:
			break loop
		case mysqlxpb.ServerMessages_RESULTSET_COLUMN_META_DATA:
			ms = append(ms, msg.(*mysqlxpb_resultset.ColumnMetaData))
		case mysqlxpb.ServerMessages_RESULTSET_ROW:
			rs = append(rs, msg.(*mysqlxpb_resultset.Row))
		}
	}

	frs := &FindResultSet{}
	if err := frs.parse(ms, rs, pc.cfg); nil != err {
		return nil, err
	}
	return frs, nil
}

func (ua *UpdateArgs) toMsg() *mysqlxpb_crud.Update {
	dataModel := mysqlxpb_crud.DataModel_TABLE

	sets := make([]*mysqlxpb_crud.UpdateOperation, 0, len(ua.Values))
	setType, exprType := mysqlxpb_crud.UpdateOperation_SET, mysqlxpb_expr.Expr_LITERAL
	for k, v := range ua.Values {
		k := k
		sets = append(sets, &mysqlxpb_crud.UpdateOperation{
			Operation: &setType,
			Source: &mysqlxpb_expr.ColumnIdentifier{
				Name: &k,
			},
			Value: &mysqlxpb_expr.Expr{
				Type:    &exprType,
				Literal: mysqlScalarFromColumnData(v),
			},
		})
	}

	where, orders, limitOffset := ua.Criteria.toMsg()

	msg := &mysqlxpb_crud.Update{
		Collection: &mysqlxpb_crud.Collection{
			Name: &ua.TableName,
		},
		DataModel: &dataModel,
		Operation: sets,
		Criteria:  where,
		Order:     orders,
		Limit:     limitOffset,
	}
	return msg
}

func (pc *poolConn) updateDelete() (uint64, error) {
	var rows uint64
	for level := 0; level < 2; {
		msg, t, err := pc.m.recvMsgUntilTypes(
			mysqlxpb.ServerMessages_NOTICE,
			mysqlxpb.ServerMessages_SQL_STMT_EXECUTE_OK,
		)
		if nil != err {
			return 0, err
		}
		switch t {
		case mysqlxpb.ServerMessages_SQL_STMT_EXECUTE_OK:
			level++
		case mysqlxpb.ServerMessages_NOTICE:
			n, err := pc.m.parseSessionStateChangedNotice(msg)
			if nil != err {
				return 0, err
			} else if nil == n {
				continue
			}

			if mysqlxpb_notice.SessionStateChanged_ROWS_AFFECTED != n.GetParam() {
				continue
			}
			level++
			if rows, err = pc.m.uint64ValueFromSessionStateChangedNotice(n); nil != err {
				return 0, err
			}
		}
	}

	return rows, nil
}

func (pc *poolConn) update(ua *UpdateArgs) (uint64, error) {
	if err := pc.m.sendMsg(mysqlxpb.ClientMessages_CRUD_UPDATE, ua.toMsg()); nil != err {
		return 0, err
	}
	return pc.updateDelete()
}

func (da *DeleteArgs) toMsg() *mysqlxpb_crud.Delete {
	dataModel := mysqlxpb_crud.DataModel_TABLE

	where, orders, limitOffset := da.Criteria.toMsg()

	msg := &mysqlxpb_crud.Delete{
		Collection: &mysqlxpb_crud.Collection{
			Name: &da.TableName,
		},
		DataModel: &dataModel,
		Criteria:  where,
		Order:     orders,
		Limit:     limitOffset,
	}
	return msg
}

func (pc *poolConn) delete(da *DeleteArgs) (uint64, error) {
	if err := pc.m.sendMsg(mysqlxpb.ClientMessages_CRUD_DELETE, da.toMsg()); nil != err {
		return 0, err
	}
	return pc.updateDelete()
}
