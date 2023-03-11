package database

import (
	"ams-fantastic-auth/internal/config"
	"ams-fantastic-auth/internal/database/schema"
	"database/sql"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //초기화를 위해 필요함
)

// Shutdown Functions
type Shutdown func() error

type Database struct {
	connDB   *sql.DB
	Shutdown Shutdown
}

func dsn(cfg *config.RDB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
}

func Open(cfg *config.RDB) (*Database, error) {
	connDsn := dsn(cfg)
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
	return &Database{
		connDB:   db,
		Shutdown: db.Close,
	}, nil
}

func (d *Database) InitTables() error {
	// schema Init
	if err := schema.CreateUsersTable(d.connDB); err != nil {
		return err
	}
	return nil
}

func (d *Database) Connect() *sql.DB {
	return d.connDB
}
