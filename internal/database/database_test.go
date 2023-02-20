package database_test

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"testing"
)

func TestNewDB(t *testing.T) {
	dbConfig := configs.Database()
	db, err := database.New(dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
}
