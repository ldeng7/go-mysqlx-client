package mysqlxclient

import (
	"fmt"

	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_expr"
)

type MysqlError struct {
	Code uint32
	Msg  string
}

func (e MysqlError) Error() string {
	return fmt.Sprintf("%s (code %d)", e.Msg, e.Code)
}

type ExprType = mysqlxpb_expr.Expr_Type

const (
	EXPR_TYPE_COLUMN_NAME   = mysqlxpb_expr.Expr_IDENT
	EXPR_TYPE_LITERAL       = mysqlxpb_expr.Expr_LITERAL
	EXPR_TYPE_FUNCTION_CALL = mysqlxpb_expr.Expr_FUNC_CALL
	EXPR_TYPE_OPERATOR      = mysqlxpb_expr.Expr_OPERATOR
)

type Expr struct {
	Type  ExprType
	Name  string
	Value interface{}
	Args  []*Expr
}

func (e *Expr) toMsg() *mysqlxpb_expr.Expr {
	msg := &mysqlxpb_expr.Expr{
		Type: &e.Type,
	}
	switch e.Type {
	case EXPR_TYPE_COLUMN_NAME:
		msg.Identifier = &mysqlxpb_expr.ColumnIdentifier{
			Name: &e.Name,
		}
	case EXPR_TYPE_LITERAL:
		msg.Literal = mysqlScalarFromColumnData(e.Value)
	case EXPR_TYPE_FUNCTION_CALL:
		params := make([]*mysqlxpb_expr.Expr, len(e.Args))
		for i, a := range e.Args {
			params[i] = a.toMsg()
		}
		msg.FunctionCall = &mysqlxpb_expr.FunctionCall{
			Name: &mysqlxpb_expr.Identifier{
				Name: &e.Name,
			},
			Param: params,
		}
	case EXPR_TYPE_OPERATOR:
		params := make([]*mysqlxpb_expr.Expr, len(e.Args))
		for i, a := range e.Args {
			params[i] = a.toMsg()
		}
		msg.Operator = &mysqlxpb_expr.Operator{
			Name:  &e.Name,
			Param: params,
		}
	}
	return msg
}

type ColumnDataType uint8

const (
	_ ColumnDataType = iota
	COLUMN_DATA_TYPE_BIT
	COLUMN_DATA_TYPE_SINT // int64
	COLUMN_DATA_TYPE_UINT
	COLUMN_DATA_TYPE_FLOAT
	COLUMN_DATA_TYPE_DOUBLE
	COLUMN_DATA_TYPE_BYTES
	COLUMN_DATA_TYPE_ENUM
	COLUMN_DATA_TYPE_SET
	COLUMN_DATA_TYPE_TIME
	COLUMN_DATA_TYPE_DATETIME
	COLUMN_DATA_TYPE_DECIMAL
	COLUMN_DATA_TYPE_GEOMETRY
)

const (
	COLUMN_FLAG_NOT_NULL       uint32 = 0x0010
	COLUMN_FLAG_PRIMARY_KEY    uint32 = 0x0020
	COLUMN_FLAG_UNIQUE_KEY     uint32 = 0x0040
	COLUMN_FLAG_MULTIPLE_KEY   uint32 = 0x0080
	COLUMN_FLAG_AUTO_INCREMENT uint32 = 0x0100
)

type ColumnMeta struct {
	DataType ColumnDataType
	Name     string
	Flag     uint32
}
