package db

import (
	"fmt"
	"net"
)

func StartSession(conn net.Conn) {
	fmt.Println("New client sessions started!", conn)
	defer func() {
		fmt.Println("disconnection from client: ", conn)
		conn.Close()
	}()

	p := NewParser(conn)

	for {
		buf := make([]byte, 1024)
		_, err := p.r.Read(buf)
		if err != nil {
			return
		}

		go ServeRequest(p.conn, buf)
	}
}
