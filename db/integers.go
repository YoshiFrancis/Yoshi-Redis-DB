package db

import (
	"fmt"
	"io"
	"strconv"
)

func ParseInteger(bytesReader io.ByteReader) int {
	sign, err := bytesReader.ReadByte()
	if err != nil {
		fmt.Println("problem parsing integer: ", err.Error())
		return -1
	}
	negativeFlag := false
	if sign == '-' {
		negativeFlag = true
	}
	numBytes, _ := bytesReader.ReadByte()
	parsed_num, err := strconv.Atoi(string(numBytes))
	if err != nil {
		fmt.Println("Given an invalid integer")
		return -1
	}
	if negativeFlag {
		parsed_num *= -1
	}
	bytesReader.ReadByte() // \r
	bytesReader.ReadByte() // \n
	return parsed_num
}
