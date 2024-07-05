package db

import (
	"bufio"
	"fmt"
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

func (p *Parser) HandleRead(b byte) {
	fmt.Print(string(b))
	if b == '*' { // array
		arrayReq := p.ParseArray()
		if arrayReq[0] == "PING" {
			p.conn.Write([]byte("+PONG\r\n"))
		}
	} else if b == '+' { // simple message

		simple_msg := p.ParseSimpleMessage()
		simple_response := HandleSimpleMessage(simple_msg)

		_, err := p.conn.Write([]byte(simple_response))
		if err != nil {
			fmt.Println("Error writing simple response", err.Error())
			return
		}
	} else {
		_, err := p.conn.Write([]byte("data type not implemented in Yoshi-Redis-DB"))
		if err != nil {
			fmt.Println("Error writing error msg response", err.Error())
			return
		}
		p.r.ReadBytes('\n')
	}
}
