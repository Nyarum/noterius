package barrel

import "reflect"

type (
	Unit interface {
		Default()
		Check(*Stats) bool
	}

	Stats struct {
		LenSlice  int
		NameField string
		Endian    int
	}

	Barrel struct {
		Object    Unit
		numField  int
		processor *Processor
		Stats     Stats
	}
)

func NewBarrel() *Barrel {
	return &Barrel{}
}

func (b *Barrel) Load(object Unit, buffer []byte, def bool) reflect.Value {
	if def {
		object.Default()
	}

	b.Object = object
	b.processor = NewProcessor(buffer)

	return reflect.Indirect(reflect.ValueOf(object))
}

func (b *Barrel) Bytes() []byte {
	return b.processor.Bytes()
}
