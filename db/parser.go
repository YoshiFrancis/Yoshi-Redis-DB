package db

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Value struct {
	typ     rune
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

func ServeRequest(conn net.Conn, ys *YoshiStore, req []byte) {
	if req[0] == '*' { // array
		values := ParseArray(conn, req[1:])
		HandleRequest(conn, ys, values)
	} else {
		_, err := conn.Write([]byte("data type must be array (*) in Yoshi-Redis-DB"))
		if err != nil {
			fmt.Println("Error writing error msg response", err.Error())
			return
		}
	}
}

func HandleRequest(conn net.Conn, ys *YoshiStore, values []Value) {
	if values[0].typ != BULK {
		return
	}
	command := strings.ToLower(values[0].bulk)
	if command == "ping" {
		conn.Write([]byte("+PONG\r\n"))
	} else if command == "echo" {
		echoMsg := Echo(values)
		conn.Write([]byte(echoMsg.bulk))
	} else if command == "set" {
		resp := Set(ys, values)
		if resp.typ == SIMPLE {
			conn.Write([]byte(resp.simple))
		} else {
			conn.Write([]byte(resp.err))
		}
	} else if command == "get" {
		resp := Get(ys, values)
		if resp.typ == BULK {
			conn.Write([]byte(resp.bulk))
		} else if resp.typ == ERROR {
			conn.Write([]byte(resp.err))
		}
	}
}
