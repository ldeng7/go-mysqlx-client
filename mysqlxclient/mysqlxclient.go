package mysqlxclient

import (
	"crypto/tls"
	"errors"
	"net"
	"time"
)

type noCopy struct{}

func (*noCopy) Lock() {}

type Client struct {
	noCopy noCopy
	pool   chan *poolConn
	cfg    ClientCfg
}

type ClientCfg struct {
	Addr     string
	Username string
	Password string
	DbName   string
	Location *time.Location
	AuthType AuthType

	InitialConns uint
	MaxConns     uint
	OnTCPDial    func(tc *net.TCPConn) error
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Tls          *tls.Config
}

func NewClient(cfg *ClientCfg) (*Client, error) {
	if nil == cfg.Location {
		cfg.Location = time.Local
	}
	if 0 == cfg.MaxConns || cfg.InitialConns > cfg.MaxConns {
		return nil, errors.New("invalid config")
	}

	c := &Client{
		cfg:  *cfg,
		pool: make(chan *poolConn, cfg.MaxConns),
	}

	for i := uint(0); i < cfg.InitialConns; i++ {
		pc, err := newPoolConn(&c.cfg)
		if nil != err {
			println("error on newPoolConn:", err.Error()) // TODO: log the error
			continue
		}
		c.pool <- pc
	}

	return c, nil
}

type InsertArgs struct {
	TableName string
	Columns   []string
	Values    [][]interface{}
}

func (c *Client) Insert(ia *InsertArgs) (uint64, uint64, error) {
	pc, err := c.getPoolConn()
	if nil != err {
		return 0, 0, err
	}
	defer c.putPoolConn(pc)
	rows, id, err := pc.insert(ia)
	if nil != err {
		pc.broken = true
	}
	return rows, id, err
}

type FindSelectItem struct {
	Expr *Expr
	As   string
}

func FindSelectItemsFromColumnNames(columnNames []string) []*FindSelectItem {
	ar := make([]*FindSelectItem, len(columnNames))
	for i, n := range columnNames {
		ar[i] = &FindSelectItem{
			Expr: &Expr{
				Type: EXPR_TYPE_COLUMN_NAME,
				Name: n,
			},
		}
	}
	return ar
}

type FindArgs struct {
	TableName string
	Select    []*FindSelectItem
}

type FindResultSet struct {
	Meta []*ColumnMeta
	Rows [][]interface{}
}

func (c *Client) Find(fa *FindArgs) (*FindResultSet, error) {
	pc, err := c.getPoolConn()
	if nil != err {
		return nil, err
	}
	defer c.putPoolConn(pc)
	rs, err := pc.find(fa)
	if nil != err {
		pc.broken = true
	}
	return rs, err
}
