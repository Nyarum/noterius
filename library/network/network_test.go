package network

import (
	"testing"
)

func TestReadToStruct(t *testing.T) {
	var TestStruct struct {
		Id    int32
		Level uint32
		HP    uint32
	}

	netes := NewParser([]byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04})
	if err := netes.ReadInt32(&TestStruct.Id).Error(); err != nil {
		t.Errorf("%v - %v", "Error read from bytes to Id field", err)
	}

	if err := netes.ReadUint32(&TestStruct.Level).Error(); err != nil {
		t.Errorf("%v - %v", "Error read from bytes to Level field", err)
	}

	if err := netes.ReadUint32(&TestStruct.HP).Error(); err != nil {
		t.Errorf("%v - %v", "Error read from byte to HP field", err)
	}
}

func TestWriteFromStruct(t *testing.T) {
	var TestStruct = struct {
		Id    int32
		Level uint32
		HP    uint32
	}{
		2, 3, 4,
	}

	netes := NewParser([]byte{})
	if err := netes.WriteInt32(TestStruct.Id).Error(); err != nil {
		t.Errorf("%v - %v", "Error write to buffer from Id field", err)
	}

	if err := netes.WriteUint32(TestStruct.Level).Error(); err != nil {
		t.Errorf("%v - %v", "Error write to buffer from Level field", err)
	}

	if err := netes.WriteUint32(TestStruct.HP).Error(); err != nil {
		t.Errorf("%v - %v", "Error write to buffer from HP field", err)
	}
}

func BenchmarkReadToStruct(b *testing.B) {
	var TestStruct struct {
		Id    int32
		Level uint32
		HP    uint32
		Name  string
	}

	netes := NewParser([]byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x07, 0x4e, 0x79, 0x61, 0x72, 0x75, 0x6d, 0x00})

	for n := 0; n < b.N; n++ {
		netes.ReadInt32(&TestStruct.Id).ReadUint32(&TestStruct.Level).ReadUint32(&TestStruct.HP).ReadString(&TestStruct.Name)
	}
}

func BenchmarkWriteFromStruct(b *testing.B) {
	var TestStruct = struct {
		Id    int32
		Level int32
		HP    int32
		Name  string
	}{
		2, 3, 4, "Nyarum",
	}

	netes := NewParser([]byte{})

	for n := 0; n < b.N; n++ {
		netes.WriteInt32(TestStruct.Id).WriteInt32(TestStruct.Level).WriteInt32(TestStruct.HP).WriteString(TestStruct.Name)
	}
}
