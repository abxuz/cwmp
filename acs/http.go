package acs

import (
	"bufio"
	"net"
	"net/http"
)

type HttpContext struct {
	conn      net.Conn
	bufReader *bufio.Reader
	bufWriter *bufio.Writer
}

func NewHttpContext(c net.Conn) *HttpContext {
	return &HttpContext{
		conn:      c,
		bufReader: bufio.NewReader(c),
		bufWriter: bufio.NewWriter(c),
	}
}

func (c *HttpContext) ReadRequest() (*http.Request, error) {
	return http.ReadRequest(c.bufReader)
}

func (c *HttpContext) Write(b []byte) (int, error) {
	return c.bufWriter.Write(b)
}

func (c *HttpContext) WriteByte(b byte) error {
	return c.bufWriter.WriteByte(b)
}

func (c *HttpContext) WriteString(s string) (int, error) {
	return c.bufWriter.WriteString(s)
}

func (c *HttpContext) Flush() error {
	return c.bufWriter.Flush()
}

func (c *HttpContext) Close() error {
	return c.conn.Close()
}
