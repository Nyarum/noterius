package packet

import (
	"fmt"
	"reflect"
	"testing"
)

func moveRangeStruct(sub string, reflectStruct reflect.Value) {
	if reflectStruct.Kind() == reflect.Struct {
		for i := 0; i < reflectStruct.NumField(); i++ {
			fieldType := reflectStruct.Type().Field(i)
			fieldValue := reflectStruct.Field(i)

			switch fieldValue.Kind() {
			case reflect.Uint8:
				fmt.Println("netes.WriteUint8(" + sub + fieldType.Name + ")")
			case reflect.Uint16:
				fmt.Println("netes.WriteUint16(" + sub + fieldType.Name + ")")
			case reflect.Uint32:
				fmt.Println("netes.WriteUint32(" + sub + fieldType.Name + ")")
			case reflect.Bool:
				fmt.Println("netes.WriteBool(" + sub + fieldType.Name + ")")
			case reflect.String:
				fmt.Println("netes.WriteString(" + sub + fieldType.Name + ")")
			case reflect.Slice:
				fmt.Println("for _, v := range " + sub + fieldType.Name + " {")
				//moveRangeStruct(sub+fieldType.Name+".", reflect.Indirect(reflect.ValueOf(reflect.New(fieldValue.Type().Elem()))))
				fmt.Println("}")
			case reflect.Array:
				fmt.Println("for _, v := range " + sub + fieldType.Name + " {")
				//moveRangeStruct(field.MustStruct())
				fmt.Println("}")
			default:
				moveRangeStruct(sub+fieldType.Name+".", fieldValue)
			}
		}
	}
}

func TestGeneratorFuncs(t *testing.T) {
	outEnterMap := OutcomingEnterMap{}
	reflectPointer := reflect.Indirect(reflect.ValueOf(&outEnterMap))

	moveRangeStruct("i.", reflectPointer)
}
