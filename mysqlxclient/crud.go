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
			n, nt := msg.(*mysqlxpb_notice.Frame), mysqlxpb_notice.Frame_SESSION_STATE_CHANGED
			if mysqlxpb_notice.Frame_Type(n.GetType()) != nt {
				continue
			}
			msg, err = pc.m.parseNotice(nt, n.Payload)
			if nil != err {
				return 0, 0, err
			}

			n1 := msg.(*mysqlxpb_notice.SessionStateChanged)
			var param *uint64
			switch n1.GetParam() {
			case mysqlxpb_notice.SessionStateChanged_ROWS_AFFECTED:
				param = &rows
			case mysqlxpb_notice.SessionStateChanged_GENERATED_INSERT_ID:
				param = &id
			default:
				continue
			}
			level++

			if len(n1.Value) < 1 {
				return 0, 0, errMysqlInvalidMessage
			} else if d, err := datatypes.MysqlDataFromScalar(n1.Value[0]); nil != err {
				return 0, 0, err
			} else if d.Type != datatypes.MYSQL_DATA_TYPE_UINT {
				return 0, 0, errMysqlInvalidMessage
			} else {
				*param = d.Value.(uint64)
			}
		}
	}

	return rows, id, nil
}

func (fa *FindArgs) toMsg() *mysqlxpb_crud.Find {
	dataModel := mysqlxpb_crud.DataModel_TABLE

	selects := make([]*mysqlxpb_crud.Projection, len(fa.Select))
	for i, s := range fa.Select {
		p := &mysqlxpb_crud.Projection{
			Source: s.Expr.toMsg(),
		}
		if 0 != len(s.As) {
			p.Alias = &s.As
		}
		selects[i] = p
	}

	groups := make([]*mysqlxpb_expr.Expr, len(fa.Groups))
	for i, g := range fa.Groups {
		groups[i] = g.toMsg()
	}

	orders := make([]*mysqlxpb_crud.Order, len(fa.Orders))
	for i, o := range fa.Orders {
		om := &mysqlxpb_crud.Order{
			Expr: o.By.toMsg(),
		}
		dir := mysqlxpb_crud.Order_ASC
		if !o.Asc {
			dir = mysqlxpb_crud.Order_DESC
		}
		om.Direction = &dir
		orders[i] = om
	}

	var limitOffset *mysqlxpb_crud.Limit
	if 0 != fa.Limit || 0 != fa.Offset {
		limitOffset = &mysqlxpb_crud.Limit{
			RowCount: &fa.Limit,
			Offset:   &fa.Offset,
		}
	}

	msg := &mysqlxpb_crud.Find{
		Collection: &mysqlxpb_crud.Collection{
			Name: &fa.TableName,
		},
		DataModel:        &dataModel,
		Projection:       selects,
		Criteria:         fa.Criteria.toMsg(),
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
