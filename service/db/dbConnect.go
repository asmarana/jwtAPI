package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	var err error
	host := "localhost"
	username := "postgres"
	password := "helloworld"
	databaseName := "jwtAuthApi"
	port := "5432"

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Africa/Lagos", username, password, host, port, databaseName)
	Database, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	return Database
}
