package datatypes

import (
	"errors"

	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_datatypes"
)

func MysqlScalarFromInt64(x int64) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_SINT
	s := &mysqlxpb_datatypes.Scalar{
		Type:       &t,
		VSignedInt: &x,
	}
	return s
}

func MysqlScalarFromUint64(x uint64) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_UINT
	s := &mysqlxpb_datatypes.Scalar{
		Type:         &t,
		VUnsignedInt: &x,
	}
	return s
}

func mysqlScalarNull() *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_NULL
	s := &mysqlxpb_datatypes.Scalar{
		Type: &t,
	}
	return s
}

func MysqlScalarFromBytes(x []byte) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_OCTETS
	s := &mysqlxpb_datatypes.Scalar{
		Type: &t,
		VOctets: &mysqlxpb_datatypes.Scalar_Octets{
			Value: x,
		},
	}
	return s
}

func MysqlScalarFromFloat64(x float64) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_DOUBLE
	s := &mysqlxpb_datatypes.Scalar{
		Type:    &t,
		VDouble: &x,
	}
	return s
}

func MysqlScalarFromFloat32(x float32) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_FLOAT
	s := &mysqlxpb_datatypes.Scalar{
		Type:   &t,
		VFloat: &x,
	}
	return s
}

func MysqlScalarFromBool(x bool) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_BOOL
	s := &mysqlxpb_datatypes.Scalar{
		Type:  &t,
		VBool: &x,
	}
	return s
}

func MysqlScalarFromString(x string) *mysqlxpb_datatypes.Scalar {
	t := mysqlxpb_datatypes.Scalar_V_STRING
	s := &mysqlxpb_datatypes.Scalar{
		Type: &t,
		VString: &mysqlxpb_datatypes.Scalar_String{
			Value: []byte(x),
		},
	}
	return s
}

func MysqlScalarFromBasic(x interface{}) *mysqlxpb_datatypes.Scalar {
	switch y := x.(type) {
	case int64:
		return MysqlScalarFromInt64(y)
	case uint64:
		return MysqlScalarFromUint64(y)
	case []byte:
		return MysqlScalarFromBytes(y)
	case float64:
		return MysqlScalarFromFloat64(y)
	case float32:
		return MysqlScalarFromFloat32(y)
	case bool:
		return MysqlScalarFromBool(y)
	case string:
		return MysqlScalarFromString(y)
	default:
		if nil == x {
			return mysqlScalarNull()
		}
		panic("invalid input type")
	}
}

func MysqlAnyFromBasic(x interface{}) *mysqlxpb_datatypes.Any {
	switch y := x.(type) {
	case []interface{}:
		t := mysqlxpb_datatypes.Any_ARRAY
		v := make([]*mysqlxpb_datatypes.Any, len(y))
		for i, e := range y {
			v[i] = MysqlAnyFromBasic(e)
		}
		a := &mysqlxpb_datatypes.Any{
			Type: &t,
			Array: &mysqlxpb_datatypes.Array{
				Value: v,
			},
		}
		return a
	case map[string]interface{}:
		t := mysqlxpb_datatypes.Any_OBJECT
		v := make([]*mysqlxpb_datatypes.Object_ObjectField, 0, len(y))
		for k, e := range y {
			v = append(v, &mysqlxpb_datatypes.Object_ObjectField{
				Key:   &k,
				Value: MysqlAnyFromBasic(e),
			})
		}
		a := &mysqlxpb_datatypes.Any{
			Type: &t,
			Obj: &mysqlxpb_datatypes.Object{
				Fld: v,
			},
		}
		return a
	default:
		t := mysqlxpb_datatypes.Any_SCALAR
		a := &mysqlxpb_datatypes.Any{
			Type:   &t,
			Scalar: MysqlScalarFromBasic(x),
		}
		return a
	}
}

type MysqlDataType uint8

const (
	_                      MysqlDataType = iota
	MYSQL_DATA_TYPE_SINT                 // int64
	MYSQL_DATA_TYPE_UINT                 // uint64
	MYSQL_DATA_TYPE_NULL                 // (nil)
	MYSQL_DATA_TYPE_OCTETS               // []byte
	MYSQL_DATA_TYPE_DOUBLE               // float64
	MYSQL_DATA_TYPE_FLOAT                // float32
	MYSQL_DATA_TYPE_BOOL                 // bool
	MYSQL_DATA_TYPE_STRING               // string
	MYSQL_DATA_TYPE_ARRAY                // []*MysqlData
	MYSQL_DATA_TYPE_OBJECT               // map[string]*MysqlData
)

type MysqlData struct {
	Value   interface{}
	Type    MysqlDataType
	SubType uint32
}

var errMysqlAnyInvalidDatatype = errors.New("invalid datatype of mysql any")

func MysqlDataFromScalar(s *mysqlxpb_datatypes.Scalar) (*MysqlData, error) {
	if nil != s {
		switch s.GetType() {
		case mysqlxpb_datatypes.Scalar_V_SINT:
			return &MysqlData{s.GetVSignedInt(), MYSQL_DATA_TYPE_SINT, 0}, nil
		case mysqlxpb_datatypes.Scalar_V_UINT:
			return &MysqlData{s.GetVUnsignedInt(), MYSQL_DATA_TYPE_UINT, 0}, nil
		case mysqlxpb_datatypes.Scalar_V_NULL:
			return &MysqlData{nil, MYSQL_DATA_TYPE_NULL, 0}, nil
		case mysqlxpb_datatypes.Scalar_V_OCTETS:
			data := &MysqlData{Type: MYSQL_DATA_TYPE_OCTETS}
			var v []byte
			if nil != s.VOctets {
				v, data.SubType = s.VOctets.Value, s.VOctets.GetContentType()
			}
			if 0 == len(v) {
				v = []byte{}
			}
			data.Value = v
			return data, nil
		case mysqlxpb_datatypes.Scalar_V_DOUBLE:
			return &MysqlData{s.GetVDouble(), MYSQL_DATA_TYPE_DOUBLE, 0}, nil
		case mysqlxpb_datatypes.Scalar_V_FLOAT:
			return &MysqlData{s.GetVFloat(), MYSQL_DATA_TYPE_FLOAT, 0}, nil
		case mysqlxpb_datatypes.Scalar_V_BOOL:
			return &MysqlData{s.GetVBool(), MYSQL_DATA_TYPE_BOOL, 0}, nil
		case mysqlxpb_datatypes.Scalar_V_STRING:
			data := &MysqlData{Type: MYSQL_DATA_TYPE_STRING}
			var v string
			if nil != s.VString {
				v, data.SubType = string(s.VString.Value), uint32(s.VString.GetCollation())
			}
			data.Value = v
			return data, nil
		default:
			return nil, errMysqlAnyInvalidDatatype
		}
	}
	return nil, errMysqlAnyInvalidDatatype
}

func MysqlDataFromAny(a *mysqlxpb_datatypes.Any) (*MysqlData, error) {
	if nil != a {
		switch a.GetType() {
		case mysqlxpb_datatypes.Any_SCALAR:
			return MysqlDataFromScalar(a.Scalar)
		case mysqlxpb_datatypes.Any_ARRAY:
			if ar := a.Array; nil != ar {
				v := make([]*MysqlData, len(ar.Value))
				for i, e := range ar.Value {
					d, err := MysqlDataFromAny(e)
					if nil != err {
						return nil, err
					}
					v[i] = d
				}
				return &MysqlData{v, MYSQL_DATA_TYPE_ARRAY, 0}, nil
			}
			return &MysqlData{[]MysqlData{}, MYSQL_DATA_TYPE_ARRAY, 0}, nil
		case mysqlxpb_datatypes.Any_OBJECT:
			if o := a.Obj; nil != o {
				v := make(map[string]*MysqlData, len(o.Fld))
				for _, e := range o.Fld {
					if nil != e && nil != e.Key {
						d, err := MysqlDataFromAny(e.Value)
						if nil != err {
							return nil, err
						}
						v[*e.Key] = d
					}
				}
				return &MysqlData{v, MYSQL_DATA_TYPE_OBJECT, 0}, nil
			}
			return &MysqlData{map[string]MysqlData{}, MYSQL_DATA_TYPE_OBJECT, 0}, nil
		default:
			return nil, errMysqlAnyInvalidDatatype
		}
	}
	return nil, errMysqlAnyInvalidDatatype
}

func (d *MysqlData) ToStringArray() []string {
	if d.Type == MYSQL_DATA_TYPE_ARRAY {
		v := d.Value.([]*MysqlData)
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
