package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"reflect"
)

func main() {
	connStr := "user=postgres password=1234 dbname=metatype sslmode=disable"
	db, err := sql.Open("postgres", connStr); handle(err)

	rows, err := db.Query(`SELECT 1`); handle(err)

	defer rows.Close()
	for rows.Next() {
		t := reflect.TypeOf(3)
		data := reflect.New(t).Interface()
		err = rows.Scan(&data); handle(err)
		fmt.Println(data)
	}

	//
	//field := &typesystem.Field{}
	//field.SetValueString("hehehe")
	//fmt.Println(field.GetValueString())
}

func handle(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}
