package db

import (
	"io"
	"strings"
)

func ParseSimpleMessage(byteReader io.ByteReader) string {
	simple_msg := make([]byte, 0)
	for {
		data, err := byteReader.ReadByte()
		if err != nil || data == '\n' {
			break
		}
		simple_msg = append(simple_msg, data)
	}

	return strings.TrimSpace(string(simple_msg))
}

func HandleSimpleMessage(m string) string {
	if m == "PING" {
		return "+PONG\r\n"
	} else {
		return "-ERR no such simple message request"
	}
}
