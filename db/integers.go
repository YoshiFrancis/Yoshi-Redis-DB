package db

import (
	"fmt"
	"strconv"
	"strings"
)

func (p *Parser) ParseInteger() int {
	numBytes, err := p.r.ReadBytes('\n')
	if err != nil {
		fmt.Println("problem parsing integer: ", err.Error())
		return -1
	}
	negativeFlag := false
	if numBytes[0] == '-' {
		numBytes[0] = '0'
		negativeFlag = true
	} else if numBytes[0] == '+' {
		numBytes[0] = '0'
	}
	parsed_num, err := strconv.Atoi(strings.TrimSpace(string(numBytes)))
	if err != nil {
		fmt.Println("Given an invalid integer")
		return -1
	}
	if negativeFlag {
		parsed_num *= -1
	}
	return parsed_num
}
