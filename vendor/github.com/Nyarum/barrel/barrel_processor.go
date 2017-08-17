package barrel

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"

	"github.com/yuin/charsetutil"
)

const (
	BigEndian = iota
	LittleEndian
)

type (
	// Processor struct for network library
	Processor struct {
		buffer *bytes.Buffer
		endian int

		err error
	}
)

// NewProcessor method for init Processor struct
func NewProcessor(buf []byte) *Processor {
	processor := new(Processor)
	processor.buffer = bytes.NewBuffer(buf)
	processor.endian = BigEndian

	return processor
}

// Error method for get error from netes
func (p *Processor) Error() error {
	return p.err
}

// Buffer method for get pointer to bytes.Buffer
func (p *Processor) Buffer() *bytes.Buffer {
	return p.buffer
}

// Bytes method for get []byte from buffer
func (p *Processor) Bytes() []byte {
	return p.buffer.Bytes()
}

// Bytes method to copy []byte from buffer
func (p *Processor) Clone() []byte {
	return append([]byte(nil), p.Bytes()...)
}

// Reset method for reset buffer
func (p *Processor) Reset() {
	p.buffer.Next(p.buffer.Len())
}

// Endian method for get current endian number
func (p *Processor) Endian() int {
	return p.endian
}

// SetEndian method for set endian in buffer
func (p *Processor) SetEndian(endian int) *Processor {
	p.endian = endian

	return p
}

func (p *Processor) ReadInt8(value *int8) *Processor {
	bufInt8 := make([]byte, 1)
	if bufInt8 = p.buffer.Next(1); len(bufInt8) < 1 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = int8(bufInt8[0])

	return p
}

func (p *Processor) ReadInt16(value *int16) *Processor {
	bufInt16 := make([]byte, 2)
	if bufInt16 = p.buffer.Next(2); len(bufInt16) < 2 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = int16(binary.LittleEndian.Uint16(bufInt16))
	} else {
		(*value) = int16(binary.BigEndian.Uint16(bufInt16))
	}

	return p
}

func (p *Processor) ReadInt32(value *int32) *Processor {
	bufInt := make([]byte, 4)
	if bufInt = p.buffer.Next(4); len(bufInt) < 4 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = int32(binary.LittleEndian.Uint32(bufInt))
	} else {
		(*value) = int32(binary.BigEndian.Uint32(bufInt))
	}

	return p
}

func (p *Processor) ReadInt64(value *int64) *Processor {
	bufInt64 := make([]byte, 8)
	if bufInt64 = p.buffer.Next(8); len(bufInt64) < 8 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = int64(binary.LittleEndian.Uint64(bufInt64))
	} else {
		(*value) = int64(binary.BigEndian.Uint64(bufInt64))
	}

	return p
}

func (p *Processor) ReadUint8(value *uint8) *Processor {
	bufUint8 := make([]byte, 1)
	if bufUint8 = p.buffer.Next(1); len(bufUint8) < 1 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = uint8(bufUint8[0])

	return p
}

func (p *Processor) ReadUint16(value *uint16) *Processor {
	bufUint16 := make([]byte, 2)
	if bufUint16 = p.buffer.Next(2); len(bufUint16) < 2 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = binary.LittleEndian.Uint16(bufUint16)
	} else {
		(*value) = binary.BigEndian.Uint16(bufUint16)
	}

	return p
}

func (p *Processor) ReadUint32(value *uint32) *Processor {
	bufUint32 := make([]byte, 4)
	if bufUint32 = p.buffer.Next(4); len(bufUint32) < 4 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = binary.LittleEndian.Uint32(bufUint32)
	} else {
		(*value) = binary.BigEndian.Uint32(bufUint32)
	}

	return p
}

func (p *Processor) ReadUint64(value *uint64) *Processor {
	bufUint64 := make([]byte, 8)
	if bufUint64 = p.buffer.Next(8); len(bufUint64) < 8 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = binary.LittleEndian.Uint64(bufUint64)
	} else {
		(*value) = binary.BigEndian.Uint64(bufUint64)
	}

	return p
}

func (p *Processor) ReadFloat32(value *float32) *Processor {
	bufFloat32 := make([]byte, 4)
	if bufFloat32 = p.buffer.Next(4); len(bufFloat32) < 4 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = math.Float32frombits(binary.LittleEndian.Uint32(bufFloat32))
	} else {
		(*value) = math.Float32frombits(binary.BigEndian.Uint32(bufFloat32))
	}

	return p
}

func (p *Processor) ReadFloat64(value *float64) *Processor {
	bufFloat64 := make([]byte, 8)
	if bufFloat64 = p.buffer.Next(8); len(bufFloat64) < 8 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		(*value) = math.Float64frombits(binary.LittleEndian.Uint64(bufFloat64))
	} else {
		(*value) = math.Float64frombits(binary.BigEndian.Uint64(bufFloat64))
	}

	return p
}

func (p *Processor) ReadString(value *string) *Processor {
	var lnString uint16
	bufLenString := make([]byte, 2)
	if bufLenString = p.buffer.Next(2); len(bufLenString) < 2 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		lnString = binary.LittleEndian.Uint16(bufLenString)
	} else {
		lnString = binary.BigEndian.Uint16(bufLenString)
	}

	bufString := make([]byte, lnString)
	if bufString = p.buffer.Next(int(lnString)); len(bufString) < int(lnString) {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	bufString = bytes.TrimSuffix(bufString, []byte{0x00})

	(*value) = string(bufString)

	return p
}

func (p *Processor) ReadString1251(value *string) *Processor {
	var lnString uint16
	bufLenString := make([]byte, 2)
	if bufLenString = p.buffer.Next(2); len(bufLenString) < 2 {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	if p.Endian() == LittleEndian {
		lnString = binary.LittleEndian.Uint16(bufLenString)
	} else {
		lnString = binary.BigEndian.Uint16(bufLenString)
	}

	bufString := make([]byte, lnString)
	if bufString = p.buffer.Next(int(lnString)); len(bufString) < int(lnString) {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	bufString = bytes.TrimSuffix(bufString, []byte{0x00})

	covertChars, err := charsetutil.DecodeBytes(bufString, "cp1251")
	if err != nil {
		p.err = err
		return p
	}

	(*value) = covertChars

	return p
}

func (p *Processor) ReadStringEOF(value *string) *Processor {
	str, err := p.buffer.ReadString(0x00)
	if err != nil {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = str

	return p
}

func (p *Processor) ReadStringWithLen(ln int32, value *string) *Processor {
	bufString := make([]byte, ln)
	if bufString = p.buffer.Next(int(ln)); len(bufString) < int(ln) {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = string(bufString)

	return p
}

func (p *Processor) ReadBytes(value *[]byte, ln int) *Processor {
	bufBytes := make([]byte, ln)
	if bufBytes = p.buffer.Next(int(ln)); len(bufBytes) < int(ln) {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = bufBytes

	return p
}

func (p *Processor) ReadArray(value *[]byte, ln int) *Processor {
	bufBytes := make([]byte, ln)
	if bufBytes = p.buffer.Next(int(ln)); len(bufBytes) < int(ln) {
		p.err = errors.New("Not enough bytes in buffer")
		return p
	}

	(*value) = bufBytes

	return p
}

func (p *Processor) ReadBool(value *bool) *Processor {
	return p
}

func (p *Processor) WriteInt8(value int8) *Processor {
	p.buffer.WriteByte(byte(value))

	return p
}

func (p *Processor) WriteInt16(value int16) *Processor {
	buf := make([]byte, 2)

	if p.Endian() == LittleEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
	} else {
		buf[0] = byte(value >> 8)
		buf[1] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Processor) WriteInt32(value int32) *Processor {
	buf := make([]byte, 4)

	if p.Endian() == LittleEndian {
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

func (p *Processor) WriteInt64(value int64) *Processor {
	buf := make([]byte, 8)

	if p.Endian() == LittleEndian {
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

func (p *Processor) WriteUint8(value uint8) *Processor {
	p.buffer.WriteByte(byte(value))

	return p
}

func (p *Processor) WriteUint16(value uint16) *Processor {
	buf := make([]byte, 2)

	if p.Endian() == LittleEndian {
		buf[0] = byte(value)
		buf[1] = byte(value >> 8)
	} else {
		buf[0] = byte(value >> 8)
		buf[1] = byte(value)
	}

	p.buffer.Write(buf)

	return p
}

func (p *Processor) WriteUint32(value uint32) *Processor {
	buf := make([]byte, 4)

	if p.Endian() == LittleEndian {
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

func (p *Processor) WriteUint64(value uint64) *Processor {
	buf := make([]byte, 8)

	if p.Endian() == LittleEndian {
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

func (p *Processor) WriteFloat32(value float32) *Processor {
	bitsFloat32 := math.Float32bits(value)
	bufFloat32 := make([]byte, 4)

	if p.Endian() == LittleEndian {
		binary.BigEndian.PutUint32(bufFloat32, bitsFloat32)
	} else {
		binary.LittleEndian.PutUint32(bufFloat32, bitsFloat32)
	}

	p.buffer.Write(bufFloat32)

	return p
}

func (p *Processor) WriteFloat64(value float64) *Processor {
	bitsFloat64 := math.Float64bits(value)
	bufFloat64 := make([]byte, 8)

	if p.Endian() == LittleEndian {
		binary.BigEndian.PutUint64(bufFloat64, bitsFloat64)
	} else {
		binary.LittleEndian.PutUint64(bufFloat64, bitsFloat64)
	}

	p.buffer.Write(bufFloat64)

	return p
}

func (p *Processor) WriteString(value string) *Processor {
	value += string([]byte{0x00})

	// Write len header for string, 2 bytes
	ln := len(value)

	buf := make([]byte, 2)

	if p.Endian() == LittleEndian {
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

func (p *Processor) WriteString1251(value string) *Processor {
	convertChars, err := charsetutil.EncodeString(value, "cp1251")
	if err != nil {
		p.err = err
		return p
	}

	value = string(convertChars)

	value += string([]byte{0x00})

	// Write len header for string, 2 bytes
	ln := len(value)

	buf := make([]byte, 2)

	if p.Endian() == LittleEndian {
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

func (p *Processor) WriteByte(value byte) *Processor {
	p.buffer.WriteByte(value)

	return p
}

func (p *Processor) WriteBytes(value []byte) *Processor {
	p.buffer.Write(value)

	return p
}

func (p *Processor) WriteBool(value bool) *Processor {
	var boolValue uint8 = 0

	if value {
		boolValue = 1
	}

	p.WriteUint8(boolValue)

	return p
}
