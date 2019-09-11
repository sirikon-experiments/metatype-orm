package typesystem

import (
	"reflect"
)

type FieldDefinition struct {
	Type reflect.Type
}

func CreateFieldDefinition(t reflect.Type) FieldDefinition {
	return FieldDefinition{Type:t}
}
