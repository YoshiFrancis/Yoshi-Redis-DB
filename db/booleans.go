package db

import "fmt"

func (p *Parser) ParseBoolean() bool {
	b, err := p.r.ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading bytes for parsing boolean")
		return false
	}

	if b[0] == 'f' {
		return false
	} else {
		return true
	}
}
