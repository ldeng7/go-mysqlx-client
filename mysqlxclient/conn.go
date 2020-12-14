package mysqlxclient

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_connection"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_notice"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_resultset"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_session"
	"github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_sql"
)

var errMysqlInvalidMessage = errors.New("invalid data from mysql")

type messenger struct {
	pc *poolConn
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

	return m.pc.sendFull(bs)
}

func (m *messenger) recvPayload() (mysqlxpb.ServerMessages_Type, []byte, error) {
	bs := make([]byte, 5)
	if err := m.pc.recvFull(bs); nil != err {
		return 0, nil, err
	}

	typ := mysqlxpb.ServerMessages_Type(bs[4])
	payload := make([]byte, binary.LittleEndian.Uint32(bs)-1)
	if err := m.pc.recvFull(payload); nil != err {
		return 0, nil, err
	}

	return typ, payload, nil
}

func (m *messenger) parsePayload(typ mysqlxpb.ServerMessages_Type, payload []byte) (proto.Message, error) {
	var msg proto.Message
	switch typ {
	case mysqlxpb.ServerMessages_OK:
		msg = &mysqlxpb.Ok{}
	case mysqlxpb.ServerMessages_ERROR:
		msg = &mysqlxpb.Error{}
	case mysqlxpb.ServerMessages_CONN_CAPABILITIES:
		msg = &mysqlxpb_connection.Capabilities{}
	case mysqlxpb.ServerMessages_SESS_AUTHENTICATE_CONTINUE:
		msg = &mysqlxpb_session.AuthenticateContinue{}
	case mysqlxpb.ServerMessages_SESS_AUTHENTICATE_OK:
		msg = &mysqlxpb_session.AuthenticateOk{}
	case mysqlxpb.ServerMessages_NOTICE:
		msg = &mysqlxpb_notice.Frame{}
	case mysqlxpb.ServerMessages_RESULTSET_COLUMN_META_DATA:
		msg = &mysqlxpb_resultset.ColumnMetaData{}
	case mysqlxpb.ServerMessages_RESULTSET_ROW:
		msg = &mysqlxpb_resultset.Row{}
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE:
		msg = &mysqlxpb_resultset.FetchDone{}
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_SUSPENDED:
		msg = &mysqlxpb_resultset.FetchSuspended{}
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
		msg = &mysqlxpb_resultset.FetchDoneMoreResultsets{}
	case mysqlxpb.ServerMessages_SQL_STMT_EXECUTE_OK:
		msg = &mysqlxpb_sql.StmtExecuteOk{}
	case mysqlxpb.ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS:
		msg = &mysqlxpb_resultset.FetchDoneMoreOutParams{}
	case mysqlxpb.ServerMessages_COMPRESSION:
		msg = &mysqlxpb_connection.Compression{}
	default:
		return nil, errMysqlInvalidMessage
	}
	if err := proto.Unmarshal(payload, msg); nil != err {
		return nil, err
	}
	return msg, nil
}

func (m *messenger) recvMsgUntilTypes(
	types ...mysqlxpb.ServerMessages_Type) (proto.Message, mysqlxpb.ServerMessages_Type, error) {
	var typeSet uint32
	for _, t := range types {
		typeSet |= (1 << t)
	}

	for {
		t, payload, err := m.recvPayload()
		if nil != err {
			return nil, 0, err
		}
		//println("recv type", mysqlxpb.ServerMessages_Type_name[int32(t)]) // TODO: log it

		if ((1 << t) & typeSet) != 0 {
			msg, err := m.parsePayload(t, payload)
			if nil != err {
				return nil, 0, err
			}
			return msg, t, nil
		}
		switch t {
		case mysqlxpb.ServerMessages_NOTICE:
			if false { // FIXME: for dev only
				m.printNotice(payload)
			}
		case mysqlxpb.ServerMessages_ERROR:
			msg, err := m.parsePayload(mysqlxpb.ServerMessages_ERROR, payload)
			if nil != err {
				return nil, 0, err
			}
			e := msg.(*mysqlxpb.Error)
			return nil, 0, &MysqlError{e.GetCode(), e.GetMsg()}
		}
	}
}

type poolConn struct {
	conn     net.Conn
	m        *messenger
	cfg      *ClientCfg
	broken   bool
	lastUsed time.Time
}

func newPoolConn(cfg *ClientCfg) (*poolConn, error) {
	conn, err := net.DialTimeout("tcp", cfg.Addr, cfg.DialTimeout)
	if nil != err {
		return nil, err
	}
	defer func() {
		if nil != err {
			conn.Close()
		}
	}()

	tc, _ := conn.(*net.TCPConn)
	tc.SetKeepAlive(true) // TODO: log the error
	if nil != cfg.OnTCPDial {
		if err = cfg.OnTCPDial(tc); nil != err {
			return nil, err
		}
	}

	pc := &poolConn{conn: conn, cfg: cfg, lastUsed: time.Now()}
	pc.m = &messenger{pc}
	if err = pc.negotiate(); nil != err {
		return nil, err
	}
	return pc, nil
}

func (pc *poolConn) sendFull(bs []byte) error {
	for {
		if timeout := pc.cfg.WriteTimeout; timeout > 0 {
			if err := pc.conn.SetWriteDeadline(time.Now().Add(timeout)); nil != err {
				return err
			}
		}
		n, err := pc.conn.Write(bs)
		if nil != err {
			return err
		}
		if n == len(bs) {
			return nil
		}
		bs = bs[n:]
	}
}

func (pc *poolConn) recvFull(bs []byte) error {
	for {
		if timeout := pc.cfg.ReadTimeout; timeout > 0 {
			if err := pc.conn.SetReadDeadline(time.Now().Add(timeout)); nil != err {
				return err
			}
		}
		n, err := pc.conn.Read(bs)
		if nil != err && io.EOF != err {
			return err
		}
		if n == len(bs) {
			return nil
		}
		bs = bs[n:]
	}
}

func (pc *poolConn) close() error {
	pc.m = nil
	if nil != pc.conn {
		// TODO: close mysql connection/session

		conn := pc.conn
		pc.conn = nil
		err := conn.Close()
		// TODO: log the error
		return err
	}
	return nil
}

func (c *Client) getPoolConn() (*poolConn, error) {
	for {
		select {
		case pc := <-c.pool:
			if time.Now().Sub(pc.lastUsed) < c.cfg.IdleTimeout {
				return pc, nil
			} else {
				pc.close()
			}
		default:
			pc, err := newPoolConn(&c.cfg)
			if nil != err {
				return nil, err
			}
			return pc, nil
		}
	}
}

func (c *Client) putPoolConn(pc *poolConn) error {
	if pc.broken {
		return pc.close()
	}
	pc.lastUsed = time.Now()
	select {
	case c.pool <- pc:
		return nil
	default:
		return pc.close()
	}
}

func (c *Client) close() {
	close(c.pool)
	for pc := range c.pool {
		pc.close()
	}
	c.pool = nil
}
