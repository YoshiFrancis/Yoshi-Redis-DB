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
		p.r.ReadByte() // the initial $
		array[idx] = p.ParseBulkString()
	}

	fmt.Println("Parsed array: ", array)
	return array

}
