package setdeadline

import (
	"net"
	"time"
)

const (
	SetDeadLineEach int = iota
	SetDeadLineBoth
	SetDeadLineNon
)

type EchoClient struct {
	c   net.Conn
	set int
}

func Dial(addr string, set int) (*EchoClient, error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	ret := &EchoClient{
		set: set,
		c:   c,
	}
	return ret, nil
}

func (c *EchoClient) Close() error {
	return c.c.Close()
}

func (c *EchoClient) Do(n int) error {
	out := []byte("11111111111111111111111111111")
	in := make([]byte, len(out))
	for i := 0; i < n; i++ {
		switch c.set {
		case SetDeadLineEach:
			if err := c.c.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}
		case SetDeadLineBoth:
			if err := c.c.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}
		}
		if _, err := c.c.Write(out); err != nil {
			return err
		}
		if c.set == SetDeadLineEach {
			if err := c.c.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}
		}
		if _, err := c.c.Read(in); err != nil {
			return err
		}
	}
	return nil
}
