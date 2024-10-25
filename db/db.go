package db

import (
	"alura-rest-base/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", config.Envs.DSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Database connected! Let's do this")

	return db
}
