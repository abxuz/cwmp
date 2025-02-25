package acs

import (
	"bytes"
	"net"
	"sync"
)

type Handler interface {
	ServeAcs(ctx *AcsContext)
}

type HandleFunc func(ctx *AcsContext)

func (f HandleFunc) ServeAcs(ctx *AcsContext) {
	f(ctx)
}

type Server struct {
	Addr string

	Handler Handler
	pool    *sync.Pool
	l       net.Listener
}

func (s *Server) Serve(l net.Listener) error {
	s.pool = &sync.Pool{
		New: func() any { return bytes.NewBuffer(nil) },
	}
	s.l = l
	defer s.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go func() {
			ctx := &AcsContext{
				httpCtx: NewHttpContext(conn),
				buffer:  s.pool.Get().(*bytes.Buffer),
			}
			ctx.buffer.Reset()
			s.Handler.ServeAcs(ctx)
			s.pool.Put(ctx.buffer)
			conn.Close()
		}()
	}
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	return s.Serve(l)
}

func (s *Server) Close() error {
	if s.l == nil {
		return nil
	}
	return s.l.Close()
}
