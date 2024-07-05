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
		var value Value
		if b == SIMPLE {
			simple_msg := ParseSimpleMessage(r)
			value = Value{typ: SIMPLE, simple: simple_msg}
		} else if b == BULK {
			bulk_string := ParseBulkString(r)
			value = Value{typ: BULK, bulk: bulk_string}
		} else if b == INTEGER {
			integer := ParseInteger(r)
			value = Value{typ: INTEGER, integer: int32(integer)}
		}

		array[idx] = value
	}

	return array

}
