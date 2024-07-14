package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "PASSWORD"
	dbName     = "notes"
	sslMode    = "disable"
	driverName = "postgres"
)

func CreateDBConn() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)

	conn, err := sql.Open(driverName, connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseDBConn(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}
