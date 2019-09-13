package core

import "github.com/sirikon-experiments/metatype-orm/cmd/src/core/typesystem"

type DataRepository interface {
	GetAll(itemDef typesystem.ItemDefinition) ([]*typesystem.Item, error)
}
