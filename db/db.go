package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	connection := "user=sam dbname=goshop password=test123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	return db
}
