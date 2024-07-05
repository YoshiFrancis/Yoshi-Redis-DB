package db

import (
	"fmt"
	"strconv"
	"strings"
)

func (p *Parser) ParseArray() (arr []string) {

	byteSize, err := p.r.ReadString('\n')
	if err != nil {
		fmt.Println("error parsing array size:", err.Error())
	}
	fmt.Println("\nbytesize in string:", string(byteSize))
	size, err := strconv.Atoi(strings.TrimSpace(string(byteSize)))
	if err != nil {
		fmt.Println("invalid size parsing array")
		return
	}

	array := make([]string, size)

	for idx := range size { // assuiming it is bulk string
		b, _ := p.r.ReadByte() // the initial $

		if b == ':' { // integer
			integer := p.ParseInteger()
			fmt.Println(integer)
		} else if b == '#' { // boolean
			_, err := p.conn.Write([]byte("data type not implemented in Yoshi-Redis-DB"))
			if err != nil {
				fmt.Println("Error writing error msg response", err.Error())
				return
			}
			p.r.ReadBytes('\n')
		} else if b == ',' { // double
			_, err := p.conn.Write([]byte("data type not implemented in Yoshi-Redis-DB"))
			if err != nil {
				fmt.Println("Error writing error msg response", err.Error())
				return
			}
			p.r.ReadBytes('\n')
		} else if b == '(' { // big number
			_, err := p.conn.Write([]byte("data type not implemented in Yoshi-Redis-DB"))
			if err != nil {
				fmt.Println("Error writing error msg response", err.Error())
				return
			}
			p.r.ReadBytes('\n')
		} else if b == '&' { // maps
			_, err := p.conn.Write([]byte("data type not implemented in Yoshi-Redis-DB"))
			if err != nil {
				fmt.Println("Error writing error msg response", err.Error())
				return
			}
			p.r.ReadBytes('\n')
		} else if b == '~' { // SETS
			_, err := p.conn.Write([]byte("data type not implemented in Yoshi-Redis-DB"))
			if err != nil {
				fmt.Println("Error writing error msg response", err.Error())
				return
			}
			p.r.ReadBytes('\n')
		}

		array[idx] = p.ParseBulkString()
	}

	fmt.Println("Parsed array: ", array)
	return array

}
