package db

import (
	"database/sql"
	"log"

	dbconfig "../dbconfig"

	_ "github.com/go-sql-driver/mysql"
)

type ReturnType struct {
	id   int
	text string
}

func (returnType *ReturnType) Id() int {
	return returnType.id
}

func (returnType *ReturnType) Text() string {
	return returnType.text
}

func DBQuery(query string) []ReturnType {
	dataSource := dbconfig.Db.User + ":" + dbconfig.Db.Pw + "@tcp(" + dbconfig.Db.Host + ")/" + dbconfig.Db.Database
	conn, err := sql.Open(dbconfig.Db.Engine, dataSource)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var result []ReturnType
	var id int
	var text string
	for rows.Next() {
		err := rows.Scan(&id, &text)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, ReturnType{id: id, text: text})
	}

	return result
}
