package typesystem

import "reflect"

type Field struct {
	definition FieldDefinition
	value interface{}
}

func CreateField(def FieldDefinition) *Field {
	return &Field{
		definition: def,
		value: nil,
	}
}

func (f *Field) GetValue() interface{} {
	return f.value
}

func (f *Field) SetValue(value interface{}) {
	if reflect.TypeOf(value) == f.GetType() {
		f.value = value
	}
}

func (f *Field) GetType() reflect.Type {
	return f.definition.Type
}
