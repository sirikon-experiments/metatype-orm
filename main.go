package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirikon-experiments/metatype-orm/typesystem"
	"log"
	"reflect"
)

func main() {
	var i int
	var s string
	var b bool

	result := query("SELECT 2, 'hello', true",
		typesystem.CreateFieldDefinition(reflect.TypeOf(i)),
		typesystem.CreateFieldDefinition(reflect.TypeOf(s)),
		typesystem.CreateFieldDefinition(reflect.TypeOf(b)))

	for _, field := range result {
		fmt.Println(field.GetValue())
	}
}

func query(s string, fieldDefinitions ...typesystem.FieldDefinition) []*typesystem.Field {
	result := make([]*typesystem.Field, len(fieldDefinitions))

	connStr := "user=postgres password=1234 dbname=metatype sslmode=disable"
	db, err := sql.Open("postgres", connStr); handle(err)
	rows, err := db.Query(s); handle(err)

	defer rows.Close()
	for rows.Next() {
		data := make([]interface{}, len(fieldDefinitions))
		for i, fieldDef := range fieldDefinitions {
			v := reflect.New(fieldDef.Type).Interface()
			data[i] = v
		}

		err = rows.Scan(data...); handle(err)

		for i, fieldDef := range fieldDefinitions {
			v := data[i]
			v = reflect.ValueOf(v).Elem().Interface()
			field := typesystem.CreateField(fieldDef)
			field.SetValue(v)
			result[i] = field
		}
	}

	return result
}

func handle(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}
