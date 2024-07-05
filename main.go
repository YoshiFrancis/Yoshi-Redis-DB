package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yoshifrancis/yoshi-redis-db/db"
)

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	ys := db.NewStorage()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go db.StartSession(conn, ys)
	}
}
