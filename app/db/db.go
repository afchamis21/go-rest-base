package db

import (
	"alura-go-base/app/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", config.Envs.DSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
