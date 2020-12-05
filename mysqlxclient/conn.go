package mysqlxclient

import (
	"errors"
	"io"
	"net"
	"time"
)

type poolConn struct {
	conn     net.Conn
	p        *pool
	m        *messenger
	broken   bool
	lastUsed time.Time
}

func newPoolConn(conn net.Conn, pool *pool) (*poolConn, error) {
	c := &poolConn{conn: conn, p: pool, lastUsed: time.Now()}
	c.m = &messenger{c}
	if err := c.auth(); nil != err {
		return nil, errors.New("authentication error: " + err.Error())
	}
	return c, nil
}

func (c *poolConn) sendFull(bs []byte) error {
	for {
		if timeout := c.p.cfg.WriteTimeout; timeout > 0 {
			if err := c.conn.SetWriteDeadline(time.Now().Add(timeout)); nil != err {
				return err
			}
		}
		n, err := c.conn.Write(bs)
		if nil != err {
			return err
		}
		if n == len(bs) {
			return nil
		}
		bs = bs[n:]
	}
}

func (c *poolConn) recvFull(bs []byte) error {
	for {
		if timeout := c.p.cfg.ReadTimeout; timeout > 0 {
			if err := c.conn.SetReadDeadline(time.Now().Add(timeout)); nil != err {
				return err
			}
		}
		n, err := c.conn.Read(bs)
		if nil != err && io.EOF != err {
			return err
		}
		if n == len(bs) {
			return nil
		}
		bs = bs[n:]
	}
}

func (c *poolConn) putBack() error {
	if !c.broken {
		c.lastUsed = time.Now()
		return c.p.put(c)
	}
	return c.close()
}

func (c *poolConn) close() error {
	c.p, c.m = nil, nil
	if nil != c.conn {
		return c.conn.Close()
	}
	return nil
}

type NewConn = func() (net.Conn, error)

type pool struct {
	ch      chan *poolConn
	cfg     ClientCfg
	newConn NewConn
}

func newPool(cfg *ClientCfg, newConn NewConn) (*pool, error) {
	if cfg.MaxConns == 0 || cfg.InitialConns > cfg.MaxConns {
		return nil, errors.New("invalid config")
	}

	p := &pool{
		ch:      make(chan *poolConn, cfg.MaxConns),
		cfg:     *cfg,
		newConn: newConn,
	}

	for i := uint(0); i < p.cfg.InitialConns; i++ {
		conn, err := p.newConn()
		if nil != err {
			continue
		}
		c, err := newPoolConn(conn, p)
		if nil != err {
			continue
		}
		p.ch <- c
	}

	return p, nil
}

func (p *pool) get() (*poolConn, error) {
	if nil == p.ch {
		return nil, errors.New("pool closed")
	}

	for {
		select {
		case c := <-p.ch:
			if time.Now().Sub(c.lastUsed) < p.cfg.IdleTimeout {
				return c, nil
			} else {
				c.close()
			}
		default:
			conn, err := p.newConn()
			if nil != err {
				return nil, err
			}
			c, err := newPoolConn(conn, p)
			if nil != err {
				return nil, err
			}
			return c, nil
		}
	}
}

func (p *pool) put(c *poolConn) error {
	if nil == p.ch {
		return c.close()
	}

	select {
	case p.ch <- c:
		return nil
	default:
		return c.close()
	}
}

func (p *pool) Close() {
	if nil == p.ch {
		return
	}
	close(p.ch)
	for c := range p.ch {
		c.close()
	}
	p.ch = nil
}
