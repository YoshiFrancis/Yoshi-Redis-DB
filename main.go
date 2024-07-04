package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yoshifrancis/yoshi-redis-db/db"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go db.StartSession(conn)
	}
}
