package db

import (
	"fmt"
	"log"
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
		b, err := p.r.ReadByte()
		if err != nil {
			log.Fatal(err)
			continue
		}

		fmt.Print(string(b))
	}
}
