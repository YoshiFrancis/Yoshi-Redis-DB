package db

import (
	"fmt"
	"strconv"
	"strings"
)

func (p *Parser) ParseBulkString() (req string) {
	byteSize, err := p.r.ReadString('\n')
	if err != nil {
		fmt.Println("error parsing array size: ", err.Error())
	}
	fmt.Println(string(byteSize))
	size, err := strconv.Atoi(strings.TrimSpace(string(byteSize)))
	if err != nil {
		fmt.Println("invalid size in parsing bulk string")
		return
	}
	data := make([]byte, size)
	_, err = p.r.Read(data)
	if err != nil {
		fmt.Println("error reading bulk string")
		return
	}
	p.r.ReadByte()
	p.r.ReadByte()

	req = string(data)
	fmt.Println("parsing bulk string:", req)
	return
}
