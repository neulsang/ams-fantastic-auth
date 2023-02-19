package database

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	raw, err := db.Query(
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) NOT NULL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			birth_year INT NOT NULL,
			birth_month INT NOT NULL,
			birth_day INT NOT NULL,
			gender ENUM('male', 'female', 'other') NOT NULL,
			password VARCHAR(255) NOT NULL,
			qna_question VARCHAR(255) NOT NULL,
			qna_answer VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	defer raw.Close()
	return nil
}
