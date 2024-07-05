package db

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
)

func ParseArray(conn net.Conn, msg []byte) (arr []Value) {

	r := bytes.NewReader(msg)
	byteSize, err := r.ReadByte()
	if err != nil {
		fmt.Println("error parsing array size:", err.Error())
	}
	fmt.Println("\nbytesize in string:", string(byteSize))
	size, err := strconv.Atoi(string(byteSize))
	if err != nil {
		fmt.Println("invalid size parsing array")
		return
	}

	r.ReadByte() // \r
	r.ReadByte() // \n

	array := make([]Value, size)

	for idx := range size { // assuiming it is bulk string
		b, _ := r.ReadByte() // the initial $

		if b == SIMPLE {
			simple_msg := ParseSimpleMessage(r)
			value := Value{typ: "simple", simple: simple_msg}
			array[idx] = value
		} else if b == BULK {
			bulk_string := ParseBulkString(r)
			value := Value{typ: "bulk", bulk: bulk_string}
			array[idx] = value
		}

	}

	fmt.Println("Parsed array: ", array)
	return array

}
