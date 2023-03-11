package database_test

import (
	"ams-fantastic-auth/internal/config"
	"ams-fantastic-auth/internal/database"
	"fmt"
	"testing"
	"time"
)

func TestNewDB(t *testing.T) {

	dbConfig := new(config.RDB)
	dbConfig.LoadConfig()
	db, err := database.Open(dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	db.Shutdown()

	fmt.Println("wait 5 sec")
	time.Sleep(time.Second * 5)
}
