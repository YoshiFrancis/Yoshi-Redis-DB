package db

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Value struct {
	typ     string
	bulk    string
	simple  string
	integer int32
	double  float32
	err     string
	array   []Value
}

const (
	ERROR   = '-'
	SIMPLE  = '+'
	ARRAY   = '*'
	INTEGER = ':'
	BULK    = '$'
	DOUBLE  = ','
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

func ServeRequest(conn net.Conn, req []byte) {
	if req[0] == '*' { // array
		values := ParseArray(conn, req[1:])
		HandleRequest(conn, values)
	} else {
		_, err := conn.Write([]byte("data type must be array (*) in Yoshi-Redis-DB"))
		if err != nil {
			fmt.Println("Error writing error msg response", err.Error())
			return
		}
	}
}

func HandleRequest(conn net.Conn, values []Value) {
	if values[0].typ != "bulk" {
		return
	}
	command := strings.ToLower(values[0].bulk)
	if command == "ping" {
		conn.Write([]byte("+PONG\r\n"))
	} else if command == "echo" {
		echoMsg := fmt.Sprintf("+%s\r\n", values[1].bulk)
		conn.Write([]byte(echoMsg))
	}
}
