package setdeadline

import (
	"sync"
	"sync/atomic"
)

type MultiClient struct {
	wg    sync.WaitGroup
	conns []*EchoClient
	err   atomic.Value
}

func DialMutiple(addr string, set int, count int) (*MultiClient, error) {
	conns := make([]*EchoClient, count)
	for i := range conns {
		conn, err := Dial(addr, set)
		if err != nil {
			return nil, err
		}
		conns[i] = conn
	}
	ret := &MultiClient{
		conns: conns,
	}

	return ret, nil
}

func (c *MultiClient) Do(n int) {
	var wg sync.WaitGroup
	wg.Add(len(c.conns))
	for i := range c.conns {
		go c.getMultiple(c.conns[i], n, &wg)
	}
	wg.Wait()
}

func (c *MultiClient) Close() {
	for i := range c.conns {
		c.conns[i].Close()
	}
}

func (c *MultiClient) Error() error {
	ret := c.err.Load()
	if ret == nil {
		return nil
	}
	return ret.(error)
}

func (c *MultiClient) getMultiple(conn *EchoClient, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := conn.Do(n); err != nil {
		c.err.Store(err)
	}
}
