package database

import (
	"database/sql"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //초기화를 위해 필요함
)

const (
	username = "tester"
	password = "test001"
	hostname = "localhost:3306"
	dbname   = "testdb"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func New() (*sql.DB, error) {
	connDsn := dsn()
	log.Println("dns: ", connDsn)

	db, err := sql.Open("mysql", connDsn)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Printf("Error %s when ping DB\n", err)
		return nil, pingErr
	}
	return db, nil
}
