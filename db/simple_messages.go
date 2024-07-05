package db

import (
	"fmt"
	"strings"
)

func (p *Parser) ParseSimpleMessage() string {
	data, err := p.r.ReadString('\n')
	if err != nil {
		fmt.Println("error parsing simple message: ", err.Error())
		return ""
	}
	return strings.TrimSpace((data))
}

func HandleSimpleMessage(m string) string {
	if m == "PING" {
		return "+PONG\r\n"
	} else {
		return "-ERR no such simple message request"
	}
}
