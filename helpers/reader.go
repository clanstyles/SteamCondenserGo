package helpers

import (
	"bytes"
	"encoding/binary"
)

const (
	intSize = 1
	shortSize = 2
)

type responseReader struct {
	position int
	data 	[]byte
}

func Init(p int, data []byte) responseReader {
	return responseReader {position: p, data: data}
}

func (self *responseReader)ReadByte() (byte) {
	newPos := self.position + intSize
	ret := self.data[self.position:newPos]
	self.position = newPos

	return ret[0]
}

func (self *responseReader)ReadShort() (int64) {
	newPos := self.position + shortSize
	val := self.data[self.position:newPos]
	self.position = newPos

	buf := bytes.NewBuffer(val)
	ret, _ := binary.ReadVarint(buf)

	return ret
}

func (self *responseReader)ReadNullTermString() (string) {
	newPos := self.position
	result := ""
	for index, b := range self.data {
		if index < self.position {
			continue
		}

		newPos++

		if b == '\x00' {
			break
		}

		result += string(b)
	}

	self.position = newPos
	return result
}
