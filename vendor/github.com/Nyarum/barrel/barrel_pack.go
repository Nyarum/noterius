package barrel

import (
	"fmt"
	"reflect"
)

func (b *Barrel) Pack(value reflect.Value) error {
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			valueType := value.Type().Field(i)

			b.Stats.NameField = valueType.Name
			if !b.Object.Check(&b.Stats) {
				continue
			}

			if b.Stats.Endian == 1 {
				b.processor.SetEndian(BigEndian)
			} else if b.Stats.Endian == 0 {
				b.processor.SetEndian(LittleEndian)
			}

			err := b.Pack(value.Field(i))
			if err != nil {
				return err
			}
		}
	case reflect.Uint8:
		b.processor.WriteUint8(uint8(value.Uint()))
	case reflect.Int8:
		b.processor.WriteInt8(int8(value.Int()))
	case reflect.Uint16:
		b.processor.WriteUint16(uint16(value.Uint()))
	case reflect.Int16:
		b.processor.WriteInt16(int16(value.Int()))
	case reflect.Uint, reflect.Uint32:
		b.processor.WriteUint32(uint32(value.Uint()))
	case reflect.Int, reflect.Int32:
		b.processor.WriteInt32(int32(value.Int()))
	case reflect.Uint64:
		b.processor.WriteUint64(value.Uint())
	case reflect.Int64:
		b.processor.WriteInt64(value.Int())
	case reflect.String:
		b.processor.WriteString(value.String())
	case reflect.Slice:
		b.processor.WriteBytes(value.Bytes())
	case reflect.Bool:
		b.processor.WriteBool(value.Bool())
	case reflect.Interface:
		err := b.Pack(reflect.Indirect(value.Elem()))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Type of object is incorrect.. It is - %v", value.Kind())
	}

	return nil
}
