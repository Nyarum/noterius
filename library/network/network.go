package network

import (
	"bytes"
	"errors"
	"fmt"
)

type Parser struct {
	buffer *bytes.Buffer
}

func NewParser(buf []byte) *Parser {
	parser := new(Parser)
	parser.buffer = bytes.NewBuffer(buf)

	return parser
}

func (p *Parser) Buffer() *bytes.Buffer {
	return p.buffer
}

func (p *Parser) Bytes() []byte {
	return p.buffer.Bytes()
}

func (p *Parser) Read(val interface{}) (err error) {
	switch val.(type) {
	case *int:
		valAun := val.(*int)
		if valAun == nil {
			err = errors.New("Int value is nil")
			return
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

	case int16:
		valAun := val.(int16)

	case int32:
		valAun := val.(int32)

	case int64:
		valAun := val.(int64)

	case uint:
		valAun := val.(uint)

	case uint16:
		valAun := val.(uint16)

	case uint32:
		valAun := val.(uint32)

	case uint64:
		valAun := val.(uint64)

	}
}
