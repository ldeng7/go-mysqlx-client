package mysqlxclient

import (
	"github.com/golang/protobuf/proto"
	"github.com/ldeng7/go-mysqlx-client/mysqlxclient/internal/datatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_notice"
)

func (m *messenger) parseNoticePayload(typ mysqlxpb_notice.Frame_Type, payload []byte) (proto.Message, error) {
	var msg proto.Message
	switch typ {
	case mysqlxpb_notice.Frame_WARNING:
		msg = &mysqlxpb_notice.Warning{}
	case mysqlxpb_notice.Frame_SESSION_VARIABLE_CHANGED:
		msg = &mysqlxpb_notice.SessionVariableChanged{}
	case mysqlxpb_notice.Frame_SESSION_STATE_CHANGED:
		msg = &mysqlxpb_notice.SessionStateChanged{}
	case mysqlxpb_notice.Frame_GROUP_REPLICATION_STATE_CHANGED:
		msg = &mysqlxpb_notice.GroupReplicationStateChanged{}
	case mysqlxpb_notice.Frame_SERVER_HELLO:
		msg = &mysqlxpb_notice.ServerHello{}
	default:
		return nil, errMysqlInvalidMessage
	}
	if err := proto.Unmarshal(payload, msg); nil != err {
		return nil, err
	}
	return msg, nil
}

func (m *messenger) parseSessionStateChangedNotice(msg proto.Message) (*mysqlxpb_notice.SessionStateChanged, error) {
	n, nt := msg.(*mysqlxpb_notice.Frame), mysqlxpb_notice.Frame_SESSION_STATE_CHANGED
	if mysqlxpb_notice.Frame_Type(n.GetType()) != nt {
		return nil, nil
	}
	msg, err := m.parseNoticePayload(nt, n.Payload)
	if nil != err {
		return nil, err
	}
	return msg.(*mysqlxpb_notice.SessionStateChanged), nil
}

func (m *messenger) uint64ValueFromSessionStateChangedNotice(n *mysqlxpb_notice.SessionStateChanged) (uint64, error) {
	if len(n.Value) < 1 {
		return 0, errMysqlInvalidMessage
	} else if d, err := datatypes.MysqlDataFromScalar(n.Value[0]); nil != err {
		return 0, err
	} else if d.Type != datatypes.MYSQL_DATA_TYPE_UINT {
		return 0, errMysqlInvalidMessage
	} else {
		return d.Value.(uint64), nil
	}
}
