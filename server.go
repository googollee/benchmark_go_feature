package setdeadline

import (
	"net"
	"sync"
)

type EchoServer struct {
	l  net.Listener
	wg sync.WaitGroup
}

func NewEcho(addr string) (*EchoServer, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	ret := &EchoServer{
		l: l,
	}
	go ret.serve()
	return ret, nil
}

func (s *EchoServer) Addr() net.Addr {
	return s.l.Addr()
}

func (s *EchoServer) Close() error {
	return s.l.Close()
}

func (s *EchoServer) Wait() {
	s.wg.Wait()
}

func (s *EchoServer) handle(c net.Conn) {
	defer func() {
		s.wg.Done()
		c.Close()
	}()

	var b [1024]byte
	for {
		n, err := c.Read(b[:])
		if err != nil {
			return
		}
		_, err = c.Write(b[:n])
		if err != nil {
			return
		}
	}
}

func (s *EchoServer) serve() {
	s.wg.Add(1)
	defer s.wg.Done()
	for {
		c, err := s.l.Accept()
		if err != nil {
			return
		}
		s.wg.Add(1)
		go s.handle(c)
	}
}
