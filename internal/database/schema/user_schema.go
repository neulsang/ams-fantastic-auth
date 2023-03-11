package schema

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	raw, err := db.Query(
		`CREATE TABLE IF NOT EXISTS user (
			id UUID PRIMARY KEY,			
			email VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			nick_name VARCHAR(36),
			phone_number VARCHAR(20) NOT NULL UNIQUE,
			birth_date DATE NOT NULL,			
			password VARCHAR(255) NOT NULL,
			gender ENUM('male', 'female', 'other') NOT NULL,
			qna_question VARCHAR(255) NOT NULL,
			qna_answer VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP DEFAULT 0
		)
	`)
	if err != nil {
		return err
	}
	defer raw.Close()
	return nil
}
