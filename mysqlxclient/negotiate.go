package mysqlxclient

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/ldeng7/go-mysqlx-client/mysqlxclient/internal/datatypes"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_connection"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_session"
)

func (m *messenger) getConnectionCapabilities() (map[string]*datatypes.MysqlData, error) {
	err := m.sendMsg(mysqlxpb.ClientMessages_CON_CAPABILITIES_GET, &mysqlxpb_connection.CapabilitiesGet{})
	if nil != err {
		return nil, err
	}
	msg, _, err := m.recvMsgUntilTypes(mysqlxpb.ServerMessages_CONN_CAPABILITIES)
	if nil != err {
		return nil, err
	}

	caps := msg.(*mysqlxpb_connection.Capabilities).Capabilities
	dm := make(map[string]*datatypes.MysqlData, len(caps))
	for _, cap := range caps {
		if k := cap.GetName(); 0 != len(k) {
			d, err := datatypes.MysqlDataFromAny(cap.GetValue())
			if nil != err {
				return nil, err
			}
			dm[k] = d
		}
	}
	return dm, nil
}

func (m *messenger) setConnectionCapabilities(capMap map[string]interface{}) error {
	caps := make([]*mysqlxpb_connection.Capability, 0, len(capMap))
	for k, v := range capMap {
		k, v := k, v
		caps = append(caps, &mysqlxpb_connection.Capability{Name: &k, Value: datatypes.MysqlAnyFromBasic(v)})
	}
	msg := &mysqlxpb_connection.CapabilitiesSet{
		Capabilities: &mysqlxpb_connection.Capabilities{
			Capabilities: caps,
		},
	}

	if err := m.sendMsg(mysqlxpb.ClientMessages_CON_CAPABILITIES_SET, msg); nil != err {
		return err
	}
	_, _, err := m.recvMsgUntilTypes(mysqlxpb.ServerMessages_OK)
	return err
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
	obj, _, err := m.recvMsgUntilTypes(typ1)
	if nil != err {
		return nil, err
	}

	if !recvOk {
		return obj.(*mysqlxpb_session.AuthenticateContinue).AuthData, nil
	}
	return nil, nil
}

func (m *messenger) close() error {
	if err := m.sendMsg(mysqlxpb.ClientMessages_CON_CLOSE, &mysqlxpb_connection.Close{}); nil != err {
		return err
	}
	_, _, err := m.recvMsgUntilTypes(mysqlxpb.ServerMessages_OK)
	return err
}

type AuthType int

const (
	AUTH_TYPE_MYSQL41 AuthType = iota
	AUTH_TYPE_PLAIN
)

var authTypePriorities = [...]AuthType{
	AUTH_TYPE_MYSQL41,
	AUTH_TYPE_PLAIN,
}
var authTypeNameTable = map[AuthType]string{
	AUTH_TYPE_MYSQL41: "MYSQL41",
	AUTH_TYPE_PLAIN:   "PLAIN",
}

var authTypeMethodTable = map[AuthType]func(*poolConn) error{
	AUTH_TYPE_MYSQL41: (*poolConn).mysql41Auth,
	AUTH_TYPE_PLAIN:   (*poolConn).plainAuth,
}

func (pc *poolConn) mysql41Auth() error {
	cha, err := pc.m.authenticate(authTypeNameTable[AUTH_TYPE_MYSQL41], nil, true, false)
	if nil != err {
		return err
	}

	u, p, d := []byte(pc.cfg.Username), []byte(pc.cfg.Password), []byte(pc.cfg.DbName)
	req := bytes.NewBuffer(make([]byte, 0, len(d)+len(u)+3+sha1.Size*2))
	req.Write(d)
	req.WriteByte(0)
	req.Write(u)
	req.WriteByte(0)
	req.WriteByte('*')

	arr := sha1.Sum(p)
	t := make([]byte, len(cha)+sha1.Size)
	copy(t, cha)
	arr1 := sha1.Sum(arr[:])
	copy(t[len(cha):], arr1[:])
	arr1 = sha1.Sum(t)
	for i := 0; i < sha1.Size; i++ {
		arr[i] ^= arr1[i]
	}
	req.WriteString(hex.EncodeToString(arr[:]))

	_, err = pc.m.authenticate("", req.Bytes(), false, true)
	return err
}

func (pc *poolConn) plainAuth() error {
	u, p, d := []byte(pc.cfg.Username), []byte(pc.cfg.Password), []byte(pc.cfg.DbName)
	req := bytes.NewBuffer(make([]byte, 0, len(d)+len(u)+len(p)+2))
	req.Write(d)
	req.WriteByte(0)
	req.Write(u)
	req.WriteByte(0)
	req.Write(p)

	_, err := pc.m.authenticate(authTypeNameTable[AUTH_TYPE_PLAIN], req.Bytes(), true, true)
	return err
}

func (pc *poolConn) auth(mechs []string) error {
	mechMap := map[string]bool{}
	for _, mech := range mechs {
		mechMap[mech] = true
	}

	mech := pc.cfg.AuthType
	if !mechMap[authTypeNameTable[mech]] {
		for _, t := range authTypePriorities {
			if mechMap[authTypeNameTable[t]] {
				mech = t
				break
			}
		}
		return errors.New("no usable authentication mechanism found")
	}
	return authTypeMethodTable[mech](pc)
}

func (pc *poolConn) negotiate() error {
	caps, err := pc.m.getConnectionCapabilities()
	if nil != err {
		return err
	}

	if tlsCfg := pc.cfg.Tls; nil != tlsCfg {
		if tlsCap, ok := caps["tls"]; !ok || tlsCap.Type != datatypes.MYSQL_DATA_TYPE_BOOL {
			return errors.New("invalid tls capability from server")
		} else if !tlsCap.Value.(bool) {
			if err = pc.m.setConnectionCapabilities(map[string]interface{}{"tls": true}); nil != err {
				return err
			}
		}

		tlsConn := tls.Client(pc.conn, tlsCfg)
		defer func() {
			if nil != err {
				tlsConn.Close()
			}
		}()
		if err = tlsConn.Handshake(); nil != err {
			return err
		}
		pc.conn = tlsConn

		caps, err = pc.m.getConnectionCapabilities()
		if nil != err {
			return err
		}
	}

	if err = pc.auth(caps["authentication.mechanisms"].ToStringArray()); nil != err {
		return errors.New("authentication error: " + err.Error())
	}

	return nil
}
