package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "1234"
	dbname   = "go_api"
)

func GetConnectionString() string {
	return "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")

	return db, nil
}
