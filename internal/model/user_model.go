package model

// User Request struct
type User struct {
	Email     string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name      string `json:"name" maxLength:"255" example:"권대근"`
	NickName  string `json:"nick_name" maxLength:"36" example:"dgkwon90"`
	BirthDate string `json:"birth_date" example:"1990-07-29"`
	Gender    string `json:"gender" enums:"male, female, other" example:"male"`
	Password  string `json:"password" maxLength:"255" example:"test1234"`
	QnA       QnA    `json:"qna"`
}

// Question and Answer struct
type QnA struct {
	Question string `json:"question" maxLength:"255" example:"What is my favorite color?"`
	Answer   string `json:"answer" maxLength:"255" example:"blue"`
}

// User Response struct
type UserResponse struct {
	ID        string `json:"id" example:"{uuid}"`
	Email     string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name      string `json:"name" maxLength:"255" example:"권대근"`
	NickName  string `json:"nick_name" maxLength:"36" example:"dgkwon90"`
	BirthDate string `json:"birthDate" example:"1990-07-29"`
	Gender    string `json:"gender" enums:"male, female, other" example:"male"`
	QnA       QnA    `json:"qna"`
	CreatedAt string `json:"created_at" example:"2023-02-27 17:03:20"`
	UpdatedAt string `json:"updated_at" example:"2023-02-27 17:03:20"`
	DeletedAt string `json:"deleted_at" example:"2023-02-27 17:03:20"`
}
