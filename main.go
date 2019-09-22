package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirikon-experiments/metatype-orm/src/core"
	"github.com/sirikon-experiments/metatype-orm/src/core/typesystem"
	"github.com/sirikon-experiments/metatype-orm/src/sql"
	"log"
	"reflect"
)

func main() {
	var integerV int
	var stringV string
	var boolV bool
	integerT := reflect.TypeOf(integerV)
	stringT := reflect.TypeOf(stringV)
	boolT := reflect.TypeOf(boolV)

	itemDef := typesystem.CreateItemDefinition("events",
		typesystem.CreateFieldDefinition("id", integerT),
		typesystem.CreateFieldDefinition("name", stringT),
		typesystem.CreateFieldDefinition("active", boolT))

	repository := getDataRepository()

	data, err := repository.GetAll(itemDef); handle(err)

	jsonResult, err := toJson(data); handle(err)

	fmt.Println(string(jsonResult))
}

func getDataRepository() core.DataRepository {
	return &sql.DataRepository{}
}

func toJson(data []*typesystem.Item) ([]byte, error) {
	result := make([]map[string]interface{}, len(data))

	for i, item := range data {
		obj := make(map[string]interface{})
		for _, col := range item.Fields {
			obj[col.GetName()] = col.GetValue()
		}
		result[i] = obj
	}

	return json.MarshalIndent(result, "", "  ")
}


func handle(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}
