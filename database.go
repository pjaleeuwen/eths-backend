package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	databaseHost = os.Getenv("ETHS_DATABASE_HOST")
	databasePort = os.Getenv("ETHS_DATABASE_PORT")
	databaseUser = os.Getenv("ETHS_DATABASE_USER")
	databasePass = os.Getenv("ETHS_DATABASE_PASS")
	databaseName = os.Getenv("ETHS_DATABASE_NAME")
	databaseSSL = "disable"

)

type Database interface{}

type pqDatabase struct {
	db *sql.DB
}

func database() Database {
	connStr := fmt.Sprintf(" host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		databaseHost,
		databasePort,
		databaseUser,
		databasePass,
		databaseName,
		databaseSSL)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &pqDatabase{db}
}
