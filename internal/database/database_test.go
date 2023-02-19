package database_test

import (
	"ams-fantastic-auth/internal/database"
	"testing"
)

func TestNewDB(t *testing.T) {
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
}
