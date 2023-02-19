package database

import (
	"ams-fantastic-auth/internal/model"
	"database/sql"
)

func InsertUser(db *sql.DB, user *model.User) error {
	stmt, err := db.Prepare("INSERT INTO users (id, email, name, birth_year, birth_month, birth_day, gender, password, qna_question, qna_answer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(user.ID, user.Email, user.Name, user.BirthDate.Year, user.BirthDate.Month, user.BirthDate.Day, user.Gender, user.Password, user.QnA.Question, user.QnA.Answer)
	return execErr
}

func SelectUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT id, email, name, birth_year, birth_month, birth_day, gender, password, qna_question, qna_answer FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		scanErr := rows.Scan(&user.ID, &user.Email, &user.Name, &user.BirthDate.Year, &user.BirthDate.Month, &user.BirthDate.Day, &user.Gender, &user.Password, &user.QnA.Question, &user.QnA.Answer)
		if scanErr != nil {
			return nil, scanErr
		}
		users = append(users, user)
	}
	return users, nil
}

func SelectUser(db *sql.DB, id string) (*model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, email, name, birth_year, birth_month, birth_day, gender, password, qna_question, qna_answer FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Name, &user.BirthDate.Year, &user.BirthDate.Month, &user.BirthDate.Day, &user.Gender, &user.Password, &user.QnA.Question, &user.QnA.Answer)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *sql.DB, user *model.User) error {
	stmt, err := db.Prepare("UPDATE users SET email = ?, name = ?, birth_year = ?, birth_month = ?, birth_day = ?, gender = ?, password = ?, qna_question = ?, qna_answer = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(user.Email, user.Name, user.BirthDate.Year, user.BirthDate.Month, user.BirthDate.Day, user.Gender, user.Password, user.QnA.Question, user.QnA.Answer, user.ID)
	if execErr != nil {
		return execErr
	}
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
