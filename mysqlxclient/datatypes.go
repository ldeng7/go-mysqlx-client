package mysqlxclient

import (
	"fmt"

	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_datatypes"
)

type MysqlError struct {
	Code uint32
	Msg  string
}

func (e MysqlError) Error() string {
	return fmt.Sprintf("%s (code %d)", e.Msg, e.Code)
}

type MysqlDataType uint8

const (
	MYSQL_DATA_TYPE_INVALID MysqlDataType = iota
	MYSQL_DATA_TYPE_SINT                  // int64
	MYSQL_DATA_TYPE_UINT                  // uint64
	MYSQL_DATA_TYPE_NULL                  // (nil)
	MYSQL_DATA_TYPE_OCTETS                // []byte
	MYSQL_DATA_TYPE_DOUBLE                // float64
	MYSQL_DATA_TYPE_FLOAT                 // float32
	MYSQL_DATA_TYPE_BOOL                  //bool
	MYSQL_DATA_TYPE_STRING                //string
	MYSQL_DATA_TYPE_ARRAY                 // []MysqlData
	MYSQL_DATA_TYPE_OBJECT                // map[string]MysqlData
)

type MysqlData struct {
	Value   interface{}
	Type    MysqlDataType
	SubType uint32
}

var mysqlDataInvalid = MysqlData{nil, MYSQL_DATA_TYPE_INVALID, 0}

func mysqlAnyToData(any *mysqlxpb_datatypes.Any) MysqlData {
	if nil != any {
		switch any.GetType() {
		case mysqlxpb_datatypes.Any_SCALAR:
			if s := any.Scalar; nil != s {
				switch s.GetType() {
				case mysqlxpb_datatypes.Scalar_V_SINT:
					return MysqlData{s.GetVSignedInt(), MYSQL_DATA_TYPE_SINT, 0}
				case mysqlxpb_datatypes.Scalar_V_UINT:
					return MysqlData{s.GetVUnsignedInt(), MYSQL_DATA_TYPE_UINT, 0}
				case mysqlxpb_datatypes.Scalar_V_NULL:
					return MysqlData{nil, MYSQL_DATA_TYPE_NULL, 0}
				case mysqlxpb_datatypes.Scalar_V_OCTETS:
					data := MysqlData{Type: MYSQL_DATA_TYPE_OCTETS}
					var v []byte
					if nil != s.VOctets {
						v, data.SubType = s.VOctets.Value, s.VOctets.GetContentType()
					}
					if 0 == len(v) {
						v = []byte{}
					}
					data.Value = v
					return data
				case mysqlxpb_datatypes.Scalar_V_DOUBLE:
					return MysqlData{s.GetVDouble(), MYSQL_DATA_TYPE_DOUBLE, 0}
				case mysqlxpb_datatypes.Scalar_V_FLOAT:
					return MysqlData{s.GetVFloat(), MYSQL_DATA_TYPE_FLOAT, 0}
				case mysqlxpb_datatypes.Scalar_V_BOOL:
					return MysqlData{s.GetVBool(), MYSQL_DATA_TYPE_BOOL, 0}
				case mysqlxpb_datatypes.Scalar_V_STRING:
					data := MysqlData{Type: MYSQL_DATA_TYPE_STRING}
					var v string
					if nil != s.VString {
						v, data.SubType = string(s.VString.Value), uint32(s.VString.GetCollation())
					}
					data.Value = v
					return data
				default:
					return mysqlDataInvalid
				}
			}
			return mysqlDataInvalid
		case mysqlxpb_datatypes.Any_ARRAY:
			if a := any.Array; nil != a {
				v := make([]MysqlData, len(a.Value))
				for i, e := range a.Value {
					v[i] = mysqlAnyToData(e)
				}
				return MysqlData{v, MYSQL_DATA_TYPE_ARRAY, 0}
			}
			return MysqlData{[]MysqlData{}, MYSQL_DATA_TYPE_ARRAY, 0}
		case mysqlxpb_datatypes.Any_OBJECT:
			if o := any.Obj; nil != o {
				v := make(map[string]MysqlData, len(o.Fld))
				for _, e := range o.Fld {
					if nil != e && nil != e.Key {
						v[*e.Key] = mysqlAnyToData(e.Value)
					}
				}
				return MysqlData{v, MYSQL_DATA_TYPE_OBJECT, 0}
			}
			return MysqlData{map[string]MysqlData{}, MYSQL_DATA_TYPE_OBJECT, 0}
		default:
			return mysqlDataInvalid
		}
	}
	return mysqlDataInvalid
}

func (d *MysqlData) ToStringArray() []string {
	if d.Type == MYSQL_DATA_TYPE_ARRAY {
		v := d.Value.([]MysqlData)
		arr := make([]string, 0, len(v))
		for _, e := range v {
			if e.Type == MYSQL_DATA_TYPE_STRING {
				arr = append(arr, e.Value.(string))
			}
		}
		return arr
	}
	return []string{}
}
