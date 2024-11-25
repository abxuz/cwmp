package acs

import (
	"net"
)

type Handler interface {
	ServeAcs(ctx *AcsContext)
}

type Server struct {
	Addr    string
	Handler Handler
}

func (s *Server) Serve(l net.Listener) error {
	defer l.Close()
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
