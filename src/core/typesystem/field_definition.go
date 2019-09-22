package typesystem

import (
	"reflect"
)

type FieldDefinition struct {
	Name string
	Type reflect.Type
}

func CreateFieldDefinition(name string, t reflect.Type) FieldDefinition {
	return FieldDefinition{
		Name:name,
		Type:t,
	}
}
