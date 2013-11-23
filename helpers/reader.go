package helpers

import (
	"bytes"
	"encoding/binary"
)

func ReadByte(data []byte, position int) (byte, int) {
	size := 1
	newPos := position + size
	ret := data[position:newPos]
	return ret[0], newPos
}

func ReadShort(data []byte, position int) (int64, int) {
	size := 2
	newPos := position + size
	val := data[position:newPos]
	buf := bytes.NewBuffer(val)

	ret, _ := binary.ReadVarint(buf)
	return ret, newPos
}

func ReadNullTermString(data []byte, position int) (string, int) {
	newPos := position
	result := ""
	for index, b := range data {
		if index < position {
			continue
		}

		newPos++

		if b == '\x00' {
			break
		}

		result += string(b)
	}

	return result, newPos
}
