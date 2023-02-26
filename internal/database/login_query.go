package database

import (
	"ams-fantastic-auth/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

func InsertToken(db *sql.DB, tokenInfo *model.Token) error {
	stmt, err := db.Prepare("INSERT INTO token (uuid, user_id, expires_at) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	tokenInfo.UUID = uuid.New().String()
	_, execErr := stmt.Exec(tokenInfo.UUID, tokenInfo.UserID, tokenInfo.ExpiresAt)
	return execErr
}

func SelectToken(db *sql.DB, tokenUUID string) (*model.Token, error) {
	tuuid, parErr := uuid.Parse(tokenUUID)
	if parErr != nil {
		return nil, parErr
	}
	var tokenInfo model.Token
	var getUUID uuid.UUID

	err := db.QueryRow("SELECT uuid, user_id, created_at  FROM token WHERE uuid = ?", tuuid).Scan(&getUUID, &tokenInfo.UserID, &tokenInfo.CreateAt)
	if err != nil {
		return nil, err
	}
	tokenInfo.UUID = getUUID.String()
	return &tokenInfo, nil
}

func DeleteToken(db *sql.DB, tokenUUID string) error {
	tuuid, parErr := uuid.Parse(tokenUUID)
	if parErr != nil {
		return parErr
	}

	stmt, err := db.Prepare("DELETE FROM token WHERE uuid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(tuuid)
	if execErr != nil {
		return execErr
	}
	return nil
}
