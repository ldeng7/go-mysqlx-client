package mysqlxclient

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/ldeng7/go-mysqlx-client/mysqlxclient/internal/datatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_notice"
)

func (m *messenger) parseNotice(typ mysqlxpb_notice.Frame_Type, payload []byte) (proto.Message, error) {
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

func (m *messenger) printNotice(payload []byte) {
	msg, err := m.parsePayload(mysqlxpb.ServerMessages_NOTICE, payload)
	if nil != err {
		println("error on parse notice:", err.Error())
		return
	}
	n := msg.(*mysqlxpb_notice.Frame)
	nt := mysqlxpb_notice.Frame_Type(n.GetType())
	println("notice type", mysqlxpb_notice.Frame_Type_name[int32(nt)])
	msg, err = m.parseNotice(nt, n.Payload)
	if nil != err {
		println("error on parse notice payload:", err.Error())
		return
	}
	switch nt {
	case mysqlxpb_notice.Frame_WARNING:
		n := msg.(*mysqlxpb_notice.Warning)
		println("code", n.GetCode(), n.GetMsg())
	case mysqlxpb_notice.Frame_SESSION_VARIABLE_CHANGED:
		n := msg.(*mysqlxpb_notice.SessionVariableChanged)
		d, _ := datatypes.MysqlDataFromScalar(n.GetValue())
		println("param", n.GetParam(), fmt.Sprintf("%+v", d.Value))
	case mysqlxpb_notice.Frame_SESSION_STATE_CHANGED:
		n := msg.(*mysqlxpb_notice.SessionStateChanged)
		vs := n.GetValue()
		ss := make([]string, len(vs))
		for i, v := range vs {
			d, _ := datatypes.MysqlDataFromScalar(v)
			ss[i] = fmt.Sprintf("%s", d.Value)
		}
		println("param", mysqlxpb_notice.SessionStateChanged_Parameter_name[int32(n.GetParam())], strings.Join(ss, ", "))
	case mysqlxpb_notice.Frame_GROUP_REPLICATION_STATE_CHANGED:
		n := msg.(*mysqlxpb_notice.GroupReplicationStateChanged)
		println("type", mysqlxpb_notice.GroupReplicationStateChanged_Type_name[int32(n.GetType())])
	}
}
