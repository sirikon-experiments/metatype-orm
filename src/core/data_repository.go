package core

import "github.com/sirikon-experiments/metatype-orm/src/core/typesystem"

type DataRepository interface {
	GetAll(itemDef typesystem.ItemDefinition) ([]*typesystem.Item, error)
}
