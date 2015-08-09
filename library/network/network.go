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
	Parser struct {
		buffer *bytes.Buffer
		endian int

		Error error
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

func (p *Parser) Read(val interface{}) (pR *Parser) {
	pR = p

	switch val.(type) {
	case *int:
		valAun := val.(*int)
		if valAun == nil {
			p.Error = errors.New("Int value is nil")
			return
		}

		var bufInt []byte
		if bufInt = p.buffer.Next(4); len(bufInt) < 4 {
			p.Error = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = int(binary.LittleEndian.Uint32(bufInt))
		} else {
			(*valAun) = int(binary.BigEndian.Uint32(bufInt))
		}

	case *int16:
		valAun := val.(*int16)
		if valAun == nil {
			p.Error = errors.New("Int16 value is nil")
			return
		}

	case *int32:
		valAun := val.(*int32)
		if valAun == nil {
			p.Error = errors.New("Int32 value is nil")
			return
		}

		var bufInt32 []byte
		if bufInt32 = p.buffer.Next(4); len(bufInt32) < 4 {
			p.Error = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = int32(binary.LittleEndian.Uint32(bufInt32))
		} else {
			(*valAun) = int32(binary.BigEndian.Uint32(bufInt32))
		}

	case *int64:
		valAun := val.(*int64)
		if valAun == nil {
			p.Error = errors.New("Int64 value is nil")
			return
		}

	case *uint:
		valAun := val.(*uint)
		if valAun == nil {
			p.Error = errors.New("Uint value is nil")
			return
		}

	case *uint16:
		valAun := val.(*uint16)
		if valAun == nil {
			p.Error = errors.New("Uint16 value is nil")
			return
		}

	case *uint32:
		valAun := val.(*uint32)
		if valAun == nil {
			p.Error = errors.New("Uint32 value is nil")
			return
		}

		var bufUint32 []byte
		if bufUint32 = p.buffer.Next(4); len(bufUint32) < 4 {
			p.Error = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = binary.LittleEndian.Uint32(bufUint32)
		} else {
			(*valAun) = binary.BigEndian.Uint32(bufUint32)
		}

	case *uint64:
		valAun := val.(*uint64)
		if valAun == nil {
			p.Error = errors.New("Uint64 value is nil")
			return
		}

	case *float32:
		valAun := val.(*float32)
		if valAun == nil {
			p.Error = errors.New("Uint64 value is nil")
			return
		}

		var bufFloat32 []byte
		if bufFloat32 = p.buffer.Next(4); len(bufFloat32) < 4 {
			p.Error = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = math.Float32frombits(binary.BigEndian.Uint32(bufFloat32))
		} else {
			(*valAun) = math.Float32frombits(binary.LittleEndian.Uint32(bufFloat32))
		}

	case *float64:
		valAun := val.(*float64)
		if valAun == nil {
			p.Error = errors.New("Uint64 value is nil")
			return
		}

		var bufFloat64 []byte
		if bufFloat64 = p.buffer.Next(8); len(bufFloat64) < 4 {
			p.Error = errors.New("Not enough bytes in buffer")
			return
		}

		if p.Endian() == BigEndian {
			(*valAun) = math.Float64frombits(binary.BigEndian.Uint64(bufFloat64))
		} else {
			(*valAun) = math.Float64frombits(binary.LittleEndian.Uint64(bufFloat64))
		}

	default:
		p.Error = errors.New("Assigned type is not supported")
	}

	return
}

func (p *Parser) Write(val interface{}) (pR *Parser) {
	pR = p

	switch val.(type) {
	case int:
		valAun := val.(int)

		buf := make([]byte, 4)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
			buf[2] = byte(valAun >> 16)
			buf[3] = byte(valAun >> 24)
		} else {
			buf[0] = byte(valAun >> 24)
			buf[1] = byte(valAun >> 16)
			buf[2] = byte(valAun >> 8)
			buf[3] = byte(valAun)
		}

		p.buffer.Write(buf)

	case int8:
		valAun := val.(int8)

		p.buffer.WriteByte(byte(valAun))

	case int16:
		valAun := val.(int16)

		buf := make([]byte, 2)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
		} else {
			buf[0] = byte(valAun >> 8)
			buf[1] = byte(valAun)
		}

		p.buffer.Write(buf)

	case int32:
		valAun := val.(int32)
		buf := make([]byte, 4)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
			buf[2] = byte(valAun >> 16)
			buf[3] = byte(valAun >> 24)
		} else {
			buf[0] = byte(valAun >> 24)
			buf[1] = byte(valAun >> 16)
			buf[2] = byte(valAun >> 8)
			buf[3] = byte(valAun)
		}

		p.buffer.Write(buf)

	case int64:
		valAun := val.(int64)

		buf := make([]byte, 8)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
			buf[2] = byte(valAun >> 16)
			buf[3] = byte(valAun >> 24)
			buf[4] = byte(valAun >> 32)
			buf[5] = byte(valAun >> 40)
			buf[6] = byte(valAun >> 48)
			buf[7] = byte(valAun >> 56)
		} else {
			buf[0] = byte(valAun >> 56)
			buf[1] = byte(valAun >> 48)
			buf[2] = byte(valAun >> 40)
			buf[3] = byte(valAun >> 32)
			buf[4] = byte(valAun >> 24)
			buf[5] = byte(valAun >> 16)
			buf[6] = byte(valAun >> 8)
			buf[7] = byte(valAun)
		}

		p.buffer.Write(buf)

	case uint:
		valAun := val.(uint)

		buf := make([]byte, 4)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
			buf[2] = byte(valAun >> 16)
			buf[3] = byte(valAun >> 24)
		} else {
			buf[0] = byte(valAun >> 24)
			buf[1] = byte(valAun >> 16)
			buf[2] = byte(valAun >> 8)
			buf[3] = byte(valAun)
		}

		p.buffer.Write(buf)

	case uint8:
		valAun := val.(uint8)

		p.buffer.WriteByte(byte(valAun))

	case uint16:
		valAun := val.(uint16)

		buf := make([]byte, 2)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
		} else {
			buf[0] = byte(valAun >> 8)
			buf[1] = byte(valAun)
		}

		p.buffer.Write(buf)

	case uint32:
		valAun := val.(uint32)

		buf := make([]byte, 4)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
			buf[2] = byte(valAun >> 16)
			buf[3] = byte(valAun >> 24)
		} else {
			buf[0] = byte(valAun >> 24)
			buf[1] = byte(valAun >> 16)
			buf[2] = byte(valAun >> 8)
			buf[3] = byte(valAun)
		}

		p.buffer.Write(buf)

	case uint64:
		valAun := val.(uint64)

		buf := make([]byte, 8)

		if p.Endian() == BigEndian {
			buf[0] = byte(valAun)
			buf[1] = byte(valAun >> 8)
			buf[2] = byte(valAun >> 16)
			buf[3] = byte(valAun >> 24)
			buf[4] = byte(valAun >> 32)
			buf[5] = byte(valAun >> 40)
			buf[6] = byte(valAun >> 48)
			buf[7] = byte(valAun >> 56)
		} else {
			buf[0] = byte(valAun >> 56)
			buf[1] = byte(valAun >> 48)
			buf[2] = byte(valAun >> 40)
			buf[3] = byte(valAun >> 32)
			buf[4] = byte(valAun >> 24)
			buf[5] = byte(valAun >> 16)
			buf[6] = byte(valAun >> 8)
			buf[7] = byte(valAun)
		}

		p.buffer.Write(buf)

	case float32:
		valAun := val.(float32)

		bitsFloat32 := math.Float32bits(valAun)
		bufFloat32 := make([]byte, 4)

		if p.Endian() == BigEndian {
			binary.BigEndian.PutUint32(bufFloat32, bitsFloat32)
		} else {
			binary.LittleEndian.PutUint32(bufFloat32, bitsFloat32)
		}

		p.buffer.Write(bufFloat32)

	case float64:
		valAun := val.(float64)

		bitsFloat64 := math.Float64bits(valAun)
		bufFloat64 := make([]byte, 8)

		if p.Endian() == BigEndian {
			binary.BigEndian.PutUint64(bufFloat64, bitsFloat64)
		} else {
			binary.LittleEndian.PutUint64(bufFloat64, bitsFloat64)
		}

		p.buffer.Write(bufFloat64)

	default:
		p.Error = errors.New("Assigned type is not supported")
	}

	return
}
