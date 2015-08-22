package network

import (
	"bytes"
)

type Netes interface {
	Buffer() *bytes.Buffer
	Bytes() []byte
	Endian() int
	SetEndian(int) *Parser

	ReadInt8(*int8) *Parser
	ReadInt16(*int16) *Parser
	ReadInt32(*int32) *Parser
	ReadInt64(*int64) *Parser
	ReadUint8(*uint8) *Parser
	ReadUint16(*uint16) *Parser
	ReadUint32(*uint32) *Parser
	ReadUint64(*uint64) *Parser
	ReadFloat32(*float32) *Parser
	ReadFloat64(*float64) *Parser
	ReadString(*string) *Parser
	ReadBytes([]byte, int) *Parser
	ReadBool(*bool) *Parser

	WriteInt8(int8) *Parser
	WriteInt16(int16) *Parser
	WriteInt32(int32) *Parser
	WriteInt64(int64) *Parser
	WriteUint8(uint8) *Parser
	WriteUint16(uint16) *Parser
	WriteUint32(uint32) *Parser
	WriteUint64(uint64) *Parser
	WriteFloat32(float32) *Parser
	WriteFloat64(float64) *Parser
	WriteString(string) *Parser
	WriteBytes([]byte) *Parser
	WriteBool(bool) *Parser
}
