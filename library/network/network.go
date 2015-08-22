package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
)

const (
	LittleEndian = iota
	BigEndian
)

type (
	// Parser struct for network library
	Parser struct {
		buffer *bytes.Buffer
		endian int

		Error error
	}
)

// NewParser method for init Parser struct
func NewParser(buf []byte) *Parser {
	parser := new(Parser)
	parser.buffer = bytes.NewBuffer(buf)
	parser.endian = LittleEndian

	return parser
}

// Buffer method for get pointer to bytes.Buffer
func (p *Parser) Buffer() *bytes.Buffer {
	return p.buffer
}

// Bytes method for get []byte from buffer
func (p *Parser) Bytes() []byte {
	return p.buffer.Bytes()
}

// Bytes method for reset buffer
func (p *Parser) Reset() {
	p.buffer.Reset()
}

// Endian method for get current endian number
func (p *Parser) Endian() int {
	return p.endian
}

// SetEndian method for set endian in buffer
func (p *Parser) SetEndian(endian int) *Parser {
	p.endian = endian

	return p
}

func (p *Parser) ReadInt8(value *int8) *Parser {
	bufInt8 := make([]byte, 1)
	if bufInt8 = p.buffer.Next(1); len(bufInt8) < 1 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = int8(bufInt8[0])

	return p
}

func (p *Parser) ReadInt16(value *int16) *Parser {
	bufInt16 := make([]byte, 2)
	if bufInt16 = p.buffer.Next(2); len(bufInt16) < 2 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = int16(binary.LittleEndian.Uint16(bufInt16))
	} else {
		(*value) = int16(binary.BigEndian.Uint16(bufInt16))
	}

	return p
}

func (p *Parser) ReadInt32(value *int32) *Parser {
	bufInt := make([]byte, 4)
	if bufInt = p.buffer.Next(4); len(bufInt) < 4 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = int32(binary.LittleEndian.Uint32(bufInt))
	} else {
		(*value) = int32(binary.BigEndian.Uint32(bufInt))
	}

	return p
}

func (p *Parser) ReadInt64(value *int64) *Parser {
	bufInt64 := make([]byte, 8)
	if bufInt64 = p.buffer.Next(8); len(bufInt64) < 8 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = int64(binary.LittleEndian.Uint64(bufInt64))
	} else {
		(*value) = int64(binary.BigEndian.Uint64(bufInt64))
	}

	return p
}

func (p *Parser) ReadUint8(value *uint8) *Parser {
	bufUint8 := make([]byte, 1)
	if bufUint8 = p.buffer.Next(1); len(bufUint8) < 1 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = uint8(bufUint8[0])

	return p
}

func (p *Parser) ReadUint16(value *uint16) *Parser {
	bufUint16 := make([]byte, 2)
	if bufUint16 = p.buffer.Next(2); len(bufUint16) < 2 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = binary.LittleEndian.Uint16(bufUint16)
	} else {
		(*value) = binary.BigEndian.Uint16(bufUint16)
	}

	return p
}

func (p *Parser) ReadUint32(value *uint32) *Parser {
	bufUint32 := make([]byte, 4)
	if bufUint32 = p.buffer.Next(4); len(bufUint32) < 4 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = binary.LittleEndian.Uint32(bufUint32)
	} else {
		(*value) = binary.BigEndian.Uint32(bufUint32)
	}

	return p
}

func (p *Parser) ReadUint64(value *uint64) *Parser {
	bufUint64 := make([]byte, 8)
	if bufUint64 = p.buffer.Next(8); len(bufUint64) < 8 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = binary.LittleEndian.Uint64(bufUint64)
	} else {
		(*value) = binary.BigEndian.Uint64(bufUint64)
	}

	return p
}

func (p *Parser) ReadFloat32(value *float32) *Parser {
	bufFloat32 := make([]byte, 4)
	if bufFloat32 = p.buffer.Next(4); len(bufFloat32) < 4 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = math.Float32frombits(binary.BigEndian.Uint32(bufFloat32))
	} else {
		(*value) = math.Float32frombits(binary.LittleEndian.Uint32(bufFloat32))
	}

	return p
}

func (p *Parser) ReadFloat64(value *float64) *Parser {
	bufFloat64 := make([]byte, 8)
	if bufFloat64 = p.buffer.Next(8); len(bufFloat64) < 8 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		(*value) = math.Float64frombits(binary.BigEndian.Uint64(bufFloat64))
	} else {
		(*value) = math.Float64frombits(binary.LittleEndian.Uint64(bufFloat64))
	}

	return p
}

func (p *Parser) ReadString(value *string) *Parser {
	var lnString uint16
	bufLenString := make([]byte, 2)
	if bufLenString = p.buffer.Next(2); len(bufLenString) < 2 {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == BigEndian {
		lnString = binary.LittleEndian.Uint16(bufLenString)
	} else {
		lnString = binary.BigEndian.Uint16(bufLenString)
	}

	bufString := make([]byte, lnString)
	if bufString = p.buffer.Next(int(lnString)); len(bufString) < int(lnString) {
		p.Error = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = string(bufString)

	return p
}

func (p *Parser) ReadBytes(value []byte, ln int) *Parser {
	return p
}

func (p *Parser) ReadBool(value *bool) *Parser {
	return p
}

func (p *Parser) WriteInt8(value int8) *Parser {
	p.buffer.WriteByte(byte(value))

	return p
}

func (p *Parser) WriteInt16(value int16) *Parser {
	buf := make([]byte, 2)

	if p.Endian() == BigEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
	} else {
		buf[0] = byte(value >> 8)
		buf[1] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Parser) WriteInt32(value int32) *Parser {
	buf := make([]byte, 4)

	if p.Endian() == BigEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
		buf[2] = byte(value >> 16)
		buf[3] = byte(value >> 24)
	} else {
		buf[0] = byte(value >> 24)
		buf[1] = byte(value >> 16)
		buf[2] = byte(value >> 8)
		buf[3] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Parser) WriteInt64(value int64) *Parser {
	buf := make([]byte, 8)

	if p.Endian() == BigEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
		buf[2] = byte(value >> 16)
		buf[3] = byte(value >> 24)
		buf[4] = byte(value >> 32)
		buf[5] = byte(value >> 40)
		buf[6] = byte(value >> 48)
		buf[7] = byte(value >> 56)
	} else {
		buf[0] = byte(value >> 56)
		buf[1] = byte(value >> 48)
		buf[2] = byte(value >> 40)
		buf[3] = byte(value >> 32)
		buf[4] = byte(value >> 24)
		buf[5] = byte(value >> 16)
		buf[6] = byte(value >> 8)
		buf[7] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Parser) WriteUint8(value uint8) *Parser {
	p.buffer.WriteByte(byte(value))

	return p
}

func (p *Parser) WriteUint16(value uint16) *Parser {
	buf := make([]byte, 2)

	if p.Endian() == BigEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
	} else {
		buf[0] = byte(value >> 8)
		buf[1] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Parser) WriteUint32(value uint32) *Parser {
	buf := make([]byte, 4)

	if p.Endian() == BigEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
		buf[2] = byte(value >> 16)
		buf[3] = byte(value >> 24)
	} else {
		buf[0] = byte(value >> 24)
		buf[1] = byte(value >> 16)
		buf[2] = byte(value >> 8)
		buf[3] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Parser) WriteUint64(value uint64) *Parser {
	buf := make([]byte, 8)

	if p.Endian() == BigEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
		buf[2] = byte(value >> 16)
		buf[3] = byte(value >> 24)
		buf[4] = byte(value >> 32)
		buf[5] = byte(value >> 40)
		buf[6] = byte(value >> 48)
		buf[7] = byte(value >> 56)
	} else {
		buf[0] = byte(value >> 56)
		buf[1] = byte(value >> 48)
		buf[2] = byte(value >> 40)
		buf[3] = byte(value >> 32)
		buf[4] = byte(value >> 24)
		buf[5] = byte(value >> 16)
		buf[6] = byte(value >> 8)
		buf[7] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Parser) WriteFloat32(value float32) *Parser {
	bitsFloat32 := math.Float32bits(value)
	bufFloat32 := make([]byte, 4)

	if p.Endian() == BigEndian {
		binary.BigEndian.PutUint32(bufFloat32, bitsFloat32)
	} else {
		binary.LittleEndian.PutUint32(bufFloat32, bitsFloat32)
	}

	p.buffer.Write(bufFloat32)

	return p
}

func (p *Parser) WriteFloat64(value float64) *Parser {
	bitsFloat64 := math.Float64bits(value)
	bufFloat64 := make([]byte, 8)

	if p.Endian() == BigEndian {
		binary.BigEndian.PutUint64(bufFloat64, bitsFloat64)
	} else {
		binary.LittleEndian.PutUint64(bufFloat64, bitsFloat64)
	}

	p.buffer.Write(bufFloat64)

	return p
}

func (p *Parser) WriteString(value string) *Parser {
	value += string([]byte{0x00})

	// Write len header for string, 2 bytes
	ln := len(value)

	buf := make([]byte, 2)

	if p.Endian() == BigEndian {
		buf[0] = byte(ln)
		buf[1] = byte(ln >> 8)
	} else {
		buf[0] = byte(ln >> 8)
		buf[1] = byte(ln)
	}

	p.buffer.Write(buf)

	// Write string
	p.buffer.WriteString(value)

	return p
}

func (p *Parser) WriteBytes(value []byte) *Parser {
	p.buffer.Write(value)

	return p
}

func (p *Parser) WriteBool(value bool) *Parser {
	return p
}
