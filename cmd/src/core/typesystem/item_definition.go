package typesystem

type ItemDefinition struct {
	Name string
	Fields []FieldDefinition
}

func CreateItemDefinition(name string, fieldDefs ...FieldDefinition) ItemDefinition {
	return ItemDefinition{
		Name:   name,
		Fields: fieldDefs,
	}
}
