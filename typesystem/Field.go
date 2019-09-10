package typesystem

type Field struct {
	value []byte
}

func (f *Field) GetValueString() string {
	return string(f.value)
}

func (f *Field) SetValueString(value string) {
	f.value = []byte(value)
}
