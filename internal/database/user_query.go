package database

import (
	"ams-fantastic-auth/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

func InsertUser(db *sql.DB, user *model.User) error {
	stmt, err := db.Prepare("INSERT INTO users (uuid, id, email, name, birthDate, gender, password, qna_question, qna_answer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(uuid.New(), user.ID, user.Email, user.Name, user.BirthDate, user.Gender, user.Password, user.QnA.Question, user.QnA.Answer)
	return execErr
}

func SelectUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT uuid, id, email, name, birthDate, gender, password, qna_question, qna_answer, created_at, updated_at, deleted_at  FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var user model.User
		var getUUID uuid.UUID
		scanErr := rows.Scan(
			&getUUID,
			&user.ID,
			&user.Email,
			&user.Name,
			&user.BirthDate,
			&user.Gender,
			&user.Password,
			&user.QnA.Question,
			&user.QnA.Answer,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt)
		if scanErr != nil {
			return nil, scanErr
		}
		user.UUID = getUUID.String()
		users = append(users, user)
	}
	return users, nil
}

func SelectUser(db *sql.DB, id string) (*model.User, error) {
	var user model.User
	var getUUID uuid.UUID
	err := db.QueryRow("SELECT uuid, id, email, name, birthDate, gender, password, qna_question, qna_answer, created_at, updated_at, deleted_at FROM users WHERE id = ?", id).Scan(&getUUID, &user.ID, &user.Email, &user.Name, &user.BirthDate, &user.Gender, &user.Password, &user.QnA.Question, &user.QnA.Answer, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	user.UUID = getUUID.String()
	return &user, nil
}

func UpdateUser(db *sql.DB, user *model.User) error {
	stmt, err := db.Prepare("UPDATE users SET email = ?, name = ?, birthDate = ?, gender = ?, password = ?, qna_question = ?, qna_answer = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	/*result*/
	_, execErr := stmt.Exec(user.Email, user.Name, user.BirthDate, user.Gender, user.Password, user.QnA.Question, user.QnA.Answer, user.ID)
	if execErr != nil {
		return execErr
	}

	// cnt, affErr := result.RowsAffected()
	// if affErr != nil {
	// 	return affErr
	// }

	return nil
}

func DeleteUser(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(id)
	if execErr != nil {
		return execErr
	}
	return nil
}
