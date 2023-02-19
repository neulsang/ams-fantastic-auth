package model

import "fmt"

// User struct
type User struct {
	ID        string    `json:"id" maxLength:"36" example:"dgkwon90"`
	Email     string    `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name      string    `json:"name" maxLength:"255" example:"권대근"`
	BirthDate BirthDate `json:"birthDate"`
	Gender    string    `json:"gender" enums:"male, female, other" example:"male"`
	Password  string    `json:"password" maxLength:"255" example:"test1234"`
	QnA       QnA       `json:"qna"`
}

func (u User) ValidateEmail() bool {
	return true
}

type BirthDate struct {
	Year  int `json:"year" example:"1990"`
	Month int `json:"month" example:"07"`
	Day   int `json:"day" example:"29"`
}

func (b BirthDate) String() string {
	return fmt.Sprintf("%04d/%02d/%02d", b.Year, b.Month, b.Day)
}

// Question and Answer struct
type QnA struct {
	Question string `json:"question" maxLength:"255" example:"What is my favorite color?"`
	Answer   string `json:"answer" maxLength:"255" example:"blue"`
}
