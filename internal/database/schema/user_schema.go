package schema

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	raw, err := db.Query(
		`CREATE TABLE IF NOT EXISTS users (
			uuid UUID PRIMARY KEY,
			id VARCHAR(36) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			birthDate DATE NOT NULL,
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

// func CreateGenderTable(db *sql.DB) error {
// 	raw, err := db.Query(
// 		`CREATE TABLE IF NOT EXISTS gender (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			label VARCHAR(8) NOT NULL UNIQUE,
// 			code INT NOT NULL,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 		)
// 	`)
// 	if err != nil {
// 		return err
// 	}
// 	defer raw.Close()
// 	return nil
// }

// func CreateQuestionTable(db *sql.DB) error {
// 	raw, err := db.Query(
// 		`CREATE TABLE IF NOT EXISTS question (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			label VARCHAR(255) NOT NULL UNIQUE,
// 			code INT NOT NULL,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 		)
// 	`)
// 	if err != nil {
// 		return err
// 	}
// 	defer raw.Close()
// 	return nil
// }
