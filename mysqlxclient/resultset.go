package mysqlxclient

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/ldeng7/go-mysqlx-client/mysqlxclient/mysqldatatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_resultset"
)

func columnMetaFromMsg(msg *mysqlxpb_resultset.ColumnMetaData) *ColumnMeta {
	m := &ColumnMeta{
		Name: string(msg.Name),
		Flag: msg.GetFlags(),
	}
	switch msg.GetType() {
	case mysqlxpb_resultset.ColumnMetaData_BIT:
		m.DataType = COLUMN_DATA_TYPE_BIT
	case mysqlxpb_resultset.ColumnMetaData_SINT:
		m.DataType = COLUMN_DATA_TYPE_SINT
	case mysqlxpb_resultset.ColumnMetaData_UINT:
		m.DataType = COLUMN_DATA_TYPE_UINT
	case mysqlxpb_resultset.ColumnMetaData_FLOAT:
		m.DataType = COLUMN_DATA_TYPE_FLOAT
	case mysqlxpb_resultset.ColumnMetaData_DOUBLE:
		m.DataType = COLUMN_DATA_TYPE_DOUBLE
	case mysqlxpb_resultset.ColumnMetaData_BYTES:
		switch mysqlxpb_resultset.ContentType_BYTES(msg.GetContentType()) {
		case mysqlxpb_resultset.ContentType_BYTES_GEOMETRY:
			m.DataType = COLUMN_DATA_TYPE_GEOMETRY
		default:
			m.DataType = COLUMN_DATA_TYPE_BYTES
		}
	case mysqlxpb_resultset.ColumnMetaData_ENUM:
		m.DataType = COLUMN_DATA_TYPE_ENUM
	case mysqlxpb_resultset.ColumnMetaData_SET:
		m.DataType = COLUMN_DATA_TYPE_SET
	case mysqlxpb_resultset.ColumnMetaData_TIME:
		m.DataType = COLUMN_DATA_TYPE_TIME
	case mysqlxpb_resultset.ColumnMetaData_DATETIME:
		m.DataType = COLUMN_DATA_TYPE_DATETIME
	case mysqlxpb_resultset.ColumnMetaData_DECIMAL:
		m.DataType = COLUMN_DATA_TYPE_DECIMAL
	}
	return m
}

func findResultSetParseBit(msg []byte, cfg *ClientCfg) (interface{}, error) {
	if msg[0] == 0 {
		return uint8(0), nil
	}
	return uint8(1), nil
}

func findResultSetParseSInt(msg []byte, cfg *ClientCfg) (interface{}, error) {
	v, n := binary.Varint(msg)
	if n <= 0 {
		return nil, errMysqlInvalidMessage
	}
	return v, nil
}

func findResultSetParseUInt(msg []byte, cfg *ClientCfg) (interface{}, error) {
	v, n := binary.Uvarint(msg)
	if n <= 0 {
		return nil, errMysqlInvalidMessage
	}
	return v, nil
}

func findResultSetParseFloat(msg []byte, cfg *ClientCfg) (interface{}, error) {
	if len(msg) < 4 {
		return nil, errMysqlInvalidMessage
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(msg)), nil
}

func findResultSetParseDouble(msg []byte, cfg *ClientCfg) (interface{}, error) {
	if len(msg) < 8 {
		return nil, errMysqlInvalidMessage
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(msg)), nil
}

func findResultSetParseBytes(msg []byte, cfg *ClientCfg) (interface{}, error) {
	return msg[:len(msg)-1], nil
}

func findResultSetParseSet(msg []byte, cfg *ClientCfg) (interface{}, error) {
	if len(msg) == 1 && msg[0] != 0 {
		return map[string]struct{}{}, nil
	}
	m := map[string]struct{}{}
	for i, ie, n := 0, len(msg), 0; i < ie; i += n {
		n = int(msg[i])
		i++
		if ie-i < n {
			return nil, errMysqlInvalidMessage
		}
		m[string(msg[i:i+n])] = struct{}{}
	}
	return m, nil
}

func findResultSetParseTime(msg []byte, cfg *ClientCfg) (interface{}, error) {
	t := &mysqldatatypes.Time{}
	if msg[0] == 0 {
		t.Positive = true
	}

	v, i, n := uint64(0), 1, 0
	v, n = binary.Uvarint(msg[i:])
	if n < 0 || v >= 839 {
		return nil, errMysqlInvalidMessage
	} else if n == 0 {
		return t, nil
	}
	t.Hour, i = uint16(v), i+n

	v, n = binary.Uvarint(msg[i:])
	if n < 0 || v >= 60 {
		return nil, errMysqlInvalidMessage
	} else if n == 0 {
		return t, nil
	}
	t.Minute, i = uint8(v), i+n

	v, n = binary.Uvarint(msg[i:])
	if n < 0 || v >= 60 {
		return nil, errMysqlInvalidMessage
	}
	t.Second = uint8(v)

	return t, nil
}

func findResultSetParseDatetime(msg []byte, cfg *ClientCfg) (interface{}, error) {
	year, n := binary.Uvarint(msg)
	if n <= 0 || year >= 10000 {
		return nil, errMysqlInvalidMessage
	}
	i := n

	month, n := binary.Uvarint(msg[i:])
	if n <= 0 || month >= 13 {
		return nil, errMysqlInvalidMessage
	}
	i += n

	day, n := binary.Uvarint(msg[i:])
	if n <= 0 || day >= 32 {
		return nil, errMysqlInvalidMessage
	}
	i += n

	var hour, minute, second uint64
	for {
		hour, n = binary.Uvarint(msg[i:])
		if n < 0 || hour >= 24 {
			return nil, errMysqlInvalidMessage
		} else if n == 0 {
			break
		}
		i += n

		minute, n = binary.Uvarint(msg[i:])
		if n < 0 || minute >= 60 {
			return nil, errMysqlInvalidMessage
		} else if n == 0 {
			break
		}
		i += n

		second, n = binary.Uvarint(msg[i:])
		if n < 0 || second >= 60 {
			return nil, errMysqlInvalidMessage
		}
		break
	}

	t := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, cfg.Location)
	return &t, nil
}

func findResultSetParseDecimal(msg []byte, cfg *ClientCfg) (interface{}, error) {
	l, lf := len(msg), int(msg[0])
	li := l - (lf >> 1) - 2
	if li < 0 {
		return nil, errMysqlInvalidMessage
	}
	li <<= 1

	var sign byte
	if b := msg[l-1]; b&0x0f == 0 {
		sign = b >> 4
	} else {
		sign = b & 0x0f
		li++
	}

	bs := make([]byte, 0, li+lf+2)
	if sign == 0x0d {
		bs = append(bs, '-')
	}
	j := 1
	nextBCD := func() byte {
		j++
		b := msg[j>>1]
		if j&1 == 0 {
			return (b >> 4) + '0'
		} else {
			return (b & 0x0f) + '0'
		}
	}
	for i := 0; i < li; i++ {
		bs = append(bs, nextBCD())
	}
	bs = append(bs, '.')
	for i := 0; i < lf; i++ {
		bs = append(bs, nextBCD())
	}

	return string(bs), nil
}

func findResultSetParseGeometry(msg []byte, cfg *ClientCfg) (interface{}, error) {
	g := &mysqldatatypes.GenericGeometry{}
	if err := g.Decode(msg); nil != err {
		return nil, err
	}
	return g, nil
}

type findResultSetParser = func([]byte, *ClientCfg) (interface{}, error)

var findResultSetParsers = map[ColumnDataType]findResultSetParser{
	COLUMN_DATA_TYPE_BIT:      findResultSetParseBit,
	COLUMN_DATA_TYPE_SINT:     findResultSetParseSInt,
	COLUMN_DATA_TYPE_UINT:     findResultSetParseUInt,
	COLUMN_DATA_TYPE_FLOAT:    findResultSetParseFloat,
	COLUMN_DATA_TYPE_DOUBLE:   findResultSetParseDouble,
	COLUMN_DATA_TYPE_BYTES:    findResultSetParseBytes,
	COLUMN_DATA_TYPE_ENUM:     findResultSetParseBytes,
	COLUMN_DATA_TYPE_SET:      findResultSetParseSet,
	COLUMN_DATA_TYPE_TIME:     findResultSetParseTime,
	COLUMN_DATA_TYPE_DATETIME: findResultSetParseDatetime,
	COLUMN_DATA_TYPE_DECIMAL:  findResultSetParseDecimal,
	COLUMN_DATA_TYPE_GEOMETRY: findResultSetParseGeometry,
}

func (frs *FindResultSet) parse(
	ms []*mysqlxpb_resultset.ColumnMetaData, rs []*mysqlxpb_resultset.Row, cfg *ClientCfg) error {
	nCols := len(ms)
	frs.Meta = make([]*ColumnMeta, nCols)
	frs.Rows = make([][]interface{}, len(rs))
	for i, m := range ms {
		frs.Meta[i] = columnMetaFromMsg(m)
	}

	for i, r := range rs {
		if len(r.Field) != nCols {
			return fmt.Errorf("failed to parse row %d: inconsistent number of columns", i)
		}
		ro := make([]interface{}, nCols)
		for j, bs := range r.Field {
			if 0 == len(bs) {
				continue
			}
			colMeta := frs.Meta[j]
			var err error
			if ro[j], err = findResultSetParsers[colMeta.DataType](bs, cfg); nil != err {
				return fmt.Errorf("failed to parse row %d column %d: %s", i, j, err.Error())
			}
		}
		frs.Rows[i] = ro
	}
	return nil
}
