package db

import (
	"fmt"
	"io"
	"strconv"
)

func ParseBulkString(bytesReader io.ByteReader) (req string) {
	byteSize, err := bytesReader.ReadByte()
	if err != nil {
		fmt.Println("error parsing array size: ", err.Error())
	}
	size, err := strconv.Atoi(string(byteSize))
	if err != nil {
		fmt.Println("invalid size in parsing bulk string")
		return
	}

	bytesReader.ReadByte() // \r
	bytesReader.ReadByte() // \n

	bulk_string := make([]byte, size)
	for idx := range size {
		b, _ := bytesReader.ReadByte()
		bulk_string[idx] = b
	}

	bytesReader.ReadByte() // \r
	bytesReader.ReadByte() // \n

	fmt.Println("Parsed bulk string:", string(bulk_string))
	return string(bulk_string)
}
