package mysqlxclient

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

var authTypePriorities = [...]AuthType{
	AUTH_TYPE_MYSQL41,
	AUTH_TYPE_SHA256_MEMORY,
	AUTH_TYPE_PLAIN,
}
var authTypeNameTable = map[AuthType]string{
	AUTH_TYPE_MYSQL41:       "MYSQL41",
	AUTH_TYPE_SHA256_MEMORY: "SHA256_MEMORY",
	AUTH_TYPE_PLAIN:         "PLAIN",
}

var authTypeMethodTable = map[AuthType]func(*poolConn) error{
	AUTH_TYPE_MYSQL41:       (*poolConn).mysql41Auth,
	AUTH_TYPE_SHA256_MEMORY: (*poolConn).sha256MemoryAuth,
	AUTH_TYPE_PLAIN:         (*poolConn).plainAuth,
}

func (c *poolConn) mysql41Auth() error {
	cha, err := c.m.authenticate(authTypeNameTable[AUTH_TYPE_MYSQL41], nil, true, false)
	if nil != err {
		return err
	}

	u, p, d := []byte(c.p.cfg.Username), []byte(c.p.cfg.Password), []byte(c.p.cfg.DbName)
	req := bytes.NewBuffer(make([]byte, 0, 43+len(u)+len(d)))
	req.Write(d)
	req.WriteByte(0)
	req.Write(u)
	req.WriteByte(0)

	arr := sha1.Sum(p)
	t := make([]byte, len(cha)+20)
	copy(t, cha)
	arr1 := sha1.Sum(arr[:])
	copy(t[len(cha):], arr1[:])
	arr1 = sha1.Sum(t)
	for i := 0; i < 20; i++ {
		arr[i] ^= arr1[i]
	}
	req.WriteString(strings.ToUpper(hex.EncodeToString(arr[:])))
	req.WriteByte(0)

	_, err = c.m.authenticate("", req.Bytes(), false, true)
	return err
}

func (c *poolConn) sha256MemoryAuth() error {
	cha, err := c.m.authenticate(authTypeNameTable[AUTH_TYPE_SHA256_MEMORY], nil, true, false)
	if nil != err {
		return err
	}

	u, p, d := []byte(c.p.cfg.Username), []byte(c.p.cfg.Password), []byte(c.p.cfg.DbName)
	req := bytes.NewBuffer(make([]byte, 0, 67+len(u)+len(d)))
	req.Write(d)
	req.WriteByte(0)
	req.Write(u)
	req.WriteByte(0)

	arr := sha256.Sum256(p)
	t := make([]byte, len(cha)+32)
	copy(t, cha)
	arr1 := sha256.Sum256(arr[:])
	copy(t[len(cha):], arr1[:])
	arr1 = sha256.Sum256(t)
	for i := 0; i < 32; i++ {
		arr[i] ^= arr1[i]
	}
	req.WriteString(strings.ToUpper(hex.EncodeToString(arr[:])))
	req.WriteByte(0)

	_, err = c.m.authenticate("", req.Bytes(), false, true)
	return err
}

func (c *poolConn) plainAuth() error {
	// TODO:
	return errors.New("oops")
}

func (c *poolConn) auth() error {
	caps, err := c.m.getConnectionCapabilities()
	if nil != err {
		return err
	}
	mechs := caps["authentication.mechanisms"].ToStringArray()
	mechmap := map[string]bool{}
	for _, mech := range mechs {
		mechmap[mech] = true
	}

	mech := c.p.cfg.AuthType
	if !mechmap[authTypeNameTable[mech]] {
		for _, t := range authTypePriorities {
			if mechmap[authTypeNameTable[t]] {
				mech = t
				break
			}
		}
		return errors.New("no usable authentication mechanism found")
	}
	return authTypeMethodTable[c.p.cfg.AuthType](c)
}
