package acs

import (
	"net"
)

type Handler interface {
	ServeAcs(ctx *AcsContext)
}

type HandleFunc func(ctx *AcsContext)

func (f HandleFunc) ServeAcs(ctx *AcsContext) {
	f(ctx)
}

type Server struct {
	Addr    string
	Handler Handler

	l net.Listener
}

func (s *Server) Serve(l net.Listener) error {
	s.l = l
	defer s.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go func() {
			defer conn.Close()
			s.Handler.ServeAcs(NewAcsContext(NewHttpContext(conn)))
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
