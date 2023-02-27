package schema

import "database/sql"

func CreateTokenTable(db *sql.DB) error {
	raw, err := db.Query(
		`CREATE TABLE IF NOT EXISTS token (
			id UUID PRIMARY KEY,
			user_id VARCHAR(36) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    		expires_at TIMESTAMP DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	defer raw.Close()
	return nil
}
