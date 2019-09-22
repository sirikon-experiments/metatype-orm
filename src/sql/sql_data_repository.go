package sql

import (
	"database/sql"
	"github.com/sirikon-experiments/metatype-orm/src/core/typesystem"
	"reflect"
	"strings"
)

type DataRepository struct {
}

func (dr *DataRepository) GetAll(itemDef typesystem.ItemDefinition) ([]*typesystem.Item, error) {
	sqlQuery := "SELECT "
	cols := make([]string, len(itemDef.Fields))
	for i, field := range itemDef.Fields {
		cols[i] = field.Name
	}
	sqlQuery += strings.Join(cols, ", ")
	sqlQuery += " FROM "+itemDef.Name

	data, err := query(sqlQuery, itemDef.Fields...)
	if err != nil {
		return nil, err
	}

	result := make([]*typesystem.Item, len(data))

	for i, row := range data {
		result[i] = &typesystem.Item{Fields:row}
	}

	return result, nil
}

func query(s string, fieldDefinitions ...typesystem.FieldDefinition) ([][]*typesystem.Field, error) {
	result := make([][]*typesystem.Field, 0)

	connStr := "user=postgres password=1234 dbname=metatype sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		row := make([]*typesystem.Field, len(fieldDefinitions))

		data := make([]interface{}, len(fieldDefinitions))
		for i, fieldDef := range fieldDefinitions {
			v := reflect.New(fieldDef.Type).Interface()
			data[i] = v
		}

		err = rows.Scan(data...)
		if err != nil {
			return nil, err
		}

		for i, fieldDef := range fieldDefinitions {
			v := data[i]
			v = reflect.ValueOf(v).Elem().Interface()
			field := typesystem.CreateField(fieldDef)
			field.SetValue(v)
			row[i] = field
		}

		result = append(result, row)
	}

	return result, nil
}
