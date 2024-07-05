package db

import (
	"fmt"
	"strconv"
	"strings"
)

func (p *Parser) ParseDouble() float32 {
	sign, err := p.r.ReadByte()
	if err != nil {
		fmt.Println("Error parsing double", err.Error())
	}
	integral_part_bytes, err := p.r.ReadBytes('.')
	if err != nil {
		fmt.Println("Error parsing double: integral part", err.Error())
	}
	fractional_part_bytes, err := p.r.ReadBytes('\r')
	if err != nil {
		fmt.Println("Error parsing double: fractional part", err.Error())
	}
	integral_part, err := strconv.Atoi(string(integral_part_bytes))
	if err != nil {
		fmt.Println("Parsing double: invalid integral part")
		return 0
	}
	fractional_part, err := strconv.Atoi(strings.TrimSpace(string(fractional_part_bytes)))
	if err != nil {
		fmt.Println("Parsing double: invalid fractional part")
		return 0
	}
	combined := float32(integral_part) + float32(fractional_part)/float32((len(fractional_part_bytes)-1))
	if sign == '-' {
		combined *= -1
	}
	return combined
}
