package database

import (
	"ams-fantastic-auth/internal/configs"
	"database/sql"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //초기화를 위해 필요함
)

func dsn(config configs.DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.UserName, config.Password, config.HostName, config.DBName)
}

func New(config configs.DBConfig) (*sql.DB, error) {
	connDsn := dsn(config)
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
