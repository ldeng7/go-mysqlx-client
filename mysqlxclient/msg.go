package mysqlxclient

import (
	"encoding/binary"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_connection"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_notice"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_resultset"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_session"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_sql"
)

type messenger struct {
	c *poolConn
}

func (m *messenger) sendMsg(typ mysqlxpb.ClientMessages_Type, msg proto.Message) error {
	payload, err := proto.Marshal(msg)
	if nil != err {
		return err
	}

	bs := make([]byte, 5+len(payload))
	binary.LittleEndian.PutUint32(bs, uint32(len(payload))+1)
	bs[4] = byte(typ)
	copy(bs[5:], payload)

	// XXX: copy buffer or send twice?
	return m.c.sendFull(bs)
}

func newServerMsgByType(typ mysqlxpb.ServerMessages_Type) (proto.Message, error) {
	switch typ {
	case mysqlxpb.ServerMessages_OK:
		return &mysqlxpb.Ok{}, nil
	case mysqlxpb.ServerMessages_ERROR:
		return &mysqlxpb.Error{}, nil
	case mysqlxpb.ServerMessages_CONN_CAPABILITIES:
		return &mysqlxpb_connection.Capabilities{}, nil
	case mysqlxpb.ServerMessages_SESS_AUTHENTICATE_CONTINUE:
		return &mysqlxpb_session.AuthenticateContinue{}, nil
	case mysqlxpb.ServerMessages_SESS_AUTHENTICATE_OK:
		return &mysqlxpb_session.AuthenticateOk{}, nil
	case mysqlxpb.ServerMessages_NOTICE:
		return &mysqlxpb_notice.Frame{}, nil
	case mysqlxpb.ServerMessages_RESULTSET_COLUMN_META_DATA:
		return &mysqlxpb_resultset.ColumnMetaData{}, nil
	case mysqlxpb.ServerMessages_RESULTSET_ROW:
		return &mysqlxpb_resultset.Row{}, nil
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE:
		return &mysqlxpb_resultset.FetchDone{}, nil
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_SUSPENDED:
		return &mysqlxpb_resultset.FetchSuspended{}, nil
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
		return &mysqlxpb_resultset.FetchDoneMoreResultsets{}, nil
	case mysqlxpb.ServerMessages_SQL_STMT_EXECUTE_OK:
		return &mysqlxpb_sql.StmtExecuteOk{}, nil
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS:
		return &mysqlxpb_resultset.FetchDoneMoreOutParams{}, nil
	case mysqlxpb.ServerMessages_COMPRESSION:
		return &mysqlxpb_connection.Compression{}, nil
	default:
		return nil, errors.New("invalid server message type")
	}
}

func (m *messenger) recvPayload() (mysqlxpb.ServerMessages_Type, []byte, error) {
	bs := make([]byte, 5)
	if err := m.c.recvFull(bs); nil != err {
		return 0, nil, err
	}

	typ := mysqlxpb.ServerMessages_Type(bs[4])
	payload := make([]byte, binary.LittleEndian.Uint32(bs)-1)
	if err := m.c.recvFull(payload); nil != err {
		return 0, nil, err
	}

	return typ, payload, nil
}

func (m *messenger) parsePayload(typ mysqlxpb.ServerMessages_Type, payload []byte) (interface{}, error) {
	obj, err := newServerMsgByType(typ)
	if nil != err {
		return nil, err
	}
	if err = proto.Unmarshal(payload, obj); nil != err {
		return nil, err
	}
	return obj, nil
}

func (m *messenger) parseMysqlError(payload []byte) error {
	obj, err := m.parsePayload(mysqlxpb.ServerMessages_ERROR, payload)
	if nil != err {
		return err
	}
	e := obj.(*mysqlxpb.Error)
	return &MysqlError{e.GetCode(), e.GetMsg()}
}

func (m *messenger) recvMsgUntil(typ mysqlxpb.ServerMessages_Type, ignoreErrMsg bool) (interface{}, error) {
	for {
		t, payload, err := m.recvPayload()
		if nil != err {
			return nil, err
		}
		switch t {
		case typ:
			return m.parsePayload(typ, payload)
		case mysqlxpb.ServerMessages_ERROR:
			if !ignoreErrMsg {
				return nil, m.parseMysqlError(payload)
			}
		}
	}
}

// functions

func (m *messenger) getConnectionCapabilities() (map[string]*MysqlData, error) {
	err := m.sendMsg(mysqlxpb.ClientMessages_CON_CAPABILITIES_GET, &mysqlxpb_connection.CapabilitiesGet{})
	if nil != err {
		return nil, err
	}
	obj, err := m.recvMsgUntil(mysqlxpb.ServerMessages_CONN_CAPABILITIES, false)
	if nil != err {
		return nil, err
	}

	caps := obj.(*mysqlxpb_connection.Capabilities).Capabilities
	dm := make(map[string]*MysqlData, len(caps))
	for _, cap := range caps {
		if k := cap.GetName(); 0 != len(k) {
			d := mysqlAnyToData(cap.GetValue())
			dm[k] = &d
		}
	}
	return dm, nil
}

func (m *messenger) authenticate(mech string, data []byte, sendStart, recvOk bool) ([]byte, error) {
	var typ mysqlxpb.ClientMessages_Type
	var msg proto.Message
	if sendStart {
		typ = mysqlxpb.ClientMessages_SESS_AUTHENTICATE_START
		msg = &mysqlxpb_session.AuthenticateStart{MechName: &mech, AuthData: data}
	} else {
		typ = mysqlxpb.ClientMessages_SESS_AUTHENTICATE_CONTINUE
		msg = &mysqlxpb_session.AuthenticateContinue{AuthData: data}
	}
	if err := m.sendMsg(typ, msg); nil != err {
		return nil, err
	}

	var typ1 mysqlxpb.ServerMessages_Type
	if recvOk {
		typ1 = mysqlxpb.ServerMessages_SESS_AUTHENTICATE_OK
	} else {
		typ1 = mysqlxpb.ServerMessages_SESS_AUTHENTICATE_CONTINUE
	}
	obj, err := m.recvMsgUntil(typ1, false)
	if nil != err {
		return nil, err
	}

	if !recvOk {
		return obj.(*mysqlxpb_session.AuthenticateContinue).AuthData, nil
	}
	return nil, nil
}
