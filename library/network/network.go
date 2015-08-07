package network

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	LittleEndian = iota
	BigEndian
)

type (
	Parser struct {
		buffer *bytes.Buffer
		endian int
	}
)

func NewParser(buf []byte) *Parser {
	parser := new(Parser)
	parser.buffer = bytes.NewBuffer(buf)
	parser.endian = LittleEndian

	return parser
}

func (p *Parser) Buffer() *bytes.Buffer {
	return p.buffer
}

func (p *Parser) Bytes() []byte {
	return p.buffer.Bytes()
}

func (p *Parser) Endian() int {
	return p.endian
}

func (p *Parser) SetEndian(endian int) *Parser {
	p.endian = endian

	return p
}

func (p *Parser) Read(val interface{}) (err error) {
	switch val.(type) {
	case *int:
		valAun := val.(*int)
		if valAun == nil {
			err = errors.New("Int value is nil")
			return
		}

		var bufInt []byte
		if bufInt = p.buffer.Next(4); len(bufInt) < 4 {
			err = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = int(binary.BigEndian.Uint32(bufInt))
		} else {
			(*valAun) = int(binary.LittleEndian.Uint32(bufInt))
		}

	case *int16:
		valAun := val.(*int16)
		if valAun == nil {
			err = errors.New("Int16 value is nil")
			return
		}

	case *int32:
		valAun := val.(*int32)
		if valAun == nil {
			err = errors.New("Int32 value is nil")
			return
		}

		var bufInt32 []byte
		if bufInt32 = p.buffer.Next(4); len(bufInt32) < 4 {
			err = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = int32(binary.BigEndian.Uint32(bufInt32))
		} else {
			(*valAun) = int32(binary.LittleEndian.Uint32(bufInt32))
		}

	case *int64:
		valAun := val.(*int64)
		if valAun == nil {
			err = errors.New("Int64 value is nil")
			return
		}

	case *uint:
		valAun := val.(*uint)
		if valAun == nil {
			err = errors.New("Uint value is nil")
			return
		}

	case *uint16:
		valAun := val.(*uint16)
		if valAun == nil {
			err = errors.New("Uint16 value is nil")
			return
		}

	case *uint32:
		valAun := val.(*uint32)
		if valAun == nil {
			err = errors.New("Uint32 value is nil")
			return
		}

	case *uint64:
		valAun := val.(*uint64)
		if valAun == nil {
			err = errors.New("Uint64 value is nil")
			return
		}

	}

	return
}

func (p *Parser) Write(val interface{}) {
	switch val.(type) {
	case int:
		valAun := val.(int)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 24))
		} else {
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case int8:
		valAun := val.(int8)

		p.buffer.WriteByte(byte(valAun))

	case int16:
		valAun := val.(int16)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
		} else {
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case int32:
		valAun := val.(int32)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 24))
		} else {
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case int64:
		valAun := val.(int64)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 32))
			p.buffer.WriteByte(byte(valAun >> 40))
			p.buffer.WriteByte(byte(valAun >> 48))
			p.buffer.WriteByte(byte(valAun >> 56))
		} else {
			p.buffer.WriteByte(byte(valAun >> 56))
			p.buffer.WriteByte(byte(valAun >> 48))
			p.buffer.WriteByte(byte(valAun >> 40))
			p.buffer.WriteByte(byte(valAun >> 32))
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case uint:
		valAun := val.(uint)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 24))
		} else {
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case uint8:
		valAun := val.(uint8)

		p.buffer.WriteByte(byte(valAun))

	case uint16:
		valAun := val.(uint16)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
		} else {
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case uint32:
		valAun := val.(uint32)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 24))
		} else {
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	case uint64:
		valAun := val.(uint64)

		if p.Endian() == BigEndian {
			p.buffer.WriteByte(byte(valAun))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 32))
			p.buffer.WriteByte(byte(valAun >> 40))
			p.buffer.WriteByte(byte(valAun >> 48))
			p.buffer.WriteByte(byte(valAun >> 56))
		} else {
			p.buffer.WriteByte(byte(valAun >> 56))
			p.buffer.WriteByte(byte(valAun >> 48))
			p.buffer.WriteByte(byte(valAun >> 40))
			p.buffer.WriteByte(byte(valAun >> 32))
			p.buffer.WriteByte(byte(valAun >> 24))
			p.buffer.WriteByte(byte(valAun >> 16))
			p.buffer.WriteByte(byte(valAun >> 8))
			p.buffer.WriteByte(byte(valAun))
		}

	}
}
