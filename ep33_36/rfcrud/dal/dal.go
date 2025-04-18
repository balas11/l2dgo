package dal

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func OpenDB(dsn string) error {
	retDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	Db = retDB
	return nil
}

func CloseDB() error {
	log.Println("Closing DB Connection")
	return Db.Close()
}

func PingTest() error {
	return Db.Ping()
}
