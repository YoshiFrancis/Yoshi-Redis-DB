package db

import (
	"bufio"
	"net"
)

type Parser struct {
	conn net.Conn
	r    *bufio.Reader
	line []byte
}

func NewParser(conn net.Conn) *Parser {
	return &Parser{
		conn: conn,
		r:    bufio.NewReader(conn),
		line: make([]byte, 0),
	}
}
