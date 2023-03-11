package database

import (
	"ams-fantastic-auth/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

func InsertUser(db *sql.DB, user *model.RegisterRequest) error {
	stmt, err := db.Prepare("INSERT INTO user (id, nick_name, email, phone_number, name, birth_date, gender, password, qna_question, qna_answer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(uuid.New(), user.NickName, user.Email, user.PhoneNumber, user.Name, user.BirthDate, user.Gender, user.Password, user.Qna.Question, user.Qna.Answer)
	return execErr
}

func SelectUsers(db *sql.DB) ([]model.UserResponse, error) {
	rows, err := db.Query("SELECT id, nick_name, email, phone_number, name, birth_date, gender, qna_question, qna_answer, created_at, updated_at, deleted_at  FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.UserResponse{}

	for rows.Next() {
		var user model.UserResponse
		var getUUID uuid.UUID
		scanErr := rows.Scan(
			&getUUID,
			&user.NickName,
			&user.Email,
			&user.PhoneNumber,
			&user.Name,
			&user.BirthDate,
			&user.Gender,
			&user.Qna.Question,
			&user.Qna.Answer,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt)
		if scanErr != nil {
			return nil, scanErr
		}
		user.ID = getUUID.String()
		users = append(users, user)
	}
	return users, nil
}

func SelectUserPassword(db *sql.DB, email string) (string, error) {
	var password string
	err := db.QueryRow("SELECT password FROM user WHERE email = ?", email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func SelectUserById(db *sql.DB, id string) (*model.UserResponse, error) {
	var user model.UserResponse
	var getUUID uuid.UUID
	err := db.QueryRow("SELECT id, nick_name, email, phone_number,name, birth_date, gender, qna_question, qna_answer, created_at, updated_at, deleted_at FROM user WHERE id = ?", id).Scan(&getUUID, &user.ID, &user.Email, &user.PhoneNumber, &user.Name, &user.BirthDate, &user.Gender, &user.Qna.Question, &user.Qna.Answer, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	user.ID = getUUID.String()
	return &user, nil
}

func SelectUserByEmail(db *sql.DB, email string) (*model.UserResponse, error) {
	var user model.UserResponse
	var getUUID uuid.UUID
	err := db.QueryRow("SELECT id, nick_name, email, phone_number, name, birth_date, gender, qna_question, qna_answer, created_at, updated_at, deleted_at FROM user WHERE email = ?", email).Scan(&getUUID, &user.ID, &user.Email, &user.PhoneNumber, &user.Name, &user.BirthDate, &user.Gender, &user.Qna.Question, &user.Qna.Answer, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	user.ID = getUUID.String()
	return &user, nil
}

func UpdateUser(db *sql.DB, id string, user *model.User) error {
	stmt, err := db.Prepare("UPDATE user SET email = ?, name = ?, birth_date = ?, gender = ?, password = ?, qna_question = ?, qna_answer = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	/*result*/
	_, execErr := stmt.Exec(user.Email, user.Name, user.BirthDate, user.Gender, user.Password, user.Qna.Question, user.Qna.Answer, id)
	if execErr != nil {
		return execErr
	}
	return nil
}

func DeleteUser(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM user WHERE id = ?")
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
