package mysqlxclient

import (
	"net"
	"time"
)

type noCopy struct{}

func (*noCopy) Lock() {}

type Client struct {
	noCopy noCopy
	pool   *pool
}

type AuthType int

const (
	AUTH_TYPE_MYSQL41 AuthType = iota
	AUTH_TYPE_SHA256_MEMORY
	AUTH_TYPE_PLAIN
)

type ClientCfg struct {
	Username string
	Password string
	DbName   string
	AuthType AuthType

	InitialConns uint
	MaxConns     uint
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewClient(cfg *ClientCfg, newConn NewConn) (*Client, error) {
	pool, err := newPool(cfg, newConn)
	if nil != err {
		return nil, err
	}
	c := &Client{
		pool: pool,
	}
	return c, nil
}

type TcpClientCfg struct {
	ClientCfg
	Addr        string
	DialTimeout time.Duration
	OnDial      func(tc *net.TCPConn) error
	//TODO: tls
}

func NewTcpClient(cfg *TcpClientCfg) (*Client, error) {
	return NewClient(&cfg.ClientCfg, func() (net.Conn, error) {
		c, err := net.DialTimeout("tcp", cfg.Addr, cfg.DialTimeout)
		if nil != err {
			return nil, err
		}

		tc, _ := c.(*net.TCPConn)
		tc.SetKeepAlive(true) //TODO: log the error
		if nil != cfg.OnDial {
			if err = cfg.OnDial(tc); nil != err {
				return nil, err
			}
		}
		return tc, nil
	})
}
