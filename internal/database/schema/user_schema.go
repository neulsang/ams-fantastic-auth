package schema

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	raw, err := db.Query(
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			nick_name VARCHAR(36),
			email VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			birth_date DATE NOT NULL,
			gender ENUM('male', 'female', 'other') NOT NULL,
			password VARCHAR(255) NOT NULL,
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
