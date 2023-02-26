package model

// User struct
type User struct {
	UUID      string `json:"uuid"`
	ID        string `json:"id" maxLength:"36" example:"dgkwon90"`
	Email     string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name      string `json:"name" maxLength:"255" example:"권대근"`
	BirthDate string `json:"birthDate" example:"1990-07-29"`
	Gender    string `json:"gender" enums:"male, female, other" example:"male"`
	Password  string `json:"password" maxLength:"255" example:"test1234"`
	QnA       QnA    `json:"qna"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// Question and Answer struct
type QnA struct {
	Question string `json:"question" maxLength:"255" example:"What is my favorite color?"`
	Answer   string `json:"answer" maxLength:"255" example:"blue"`
}

// // Gender struct
// type Gender struct {
// 	ID    int    `json:"id" example:"1"`
// 	Label string `json:"label" maxLength:"8" enums:"남성, 여성, 기타" example:"남성"`
// 	Code  int    `json:"code" example:"1"`
// }

// // Question struct
// type Question struct {
// 	ID    int    `json:"id" example:"1"`
// 	Label string `json:"label" maxLength:"255" enums:"What is my favorite color?, 고향은?, 별명은?" example:"What is my favorite color?"`
// 	Code  int    `json:"code" example:"1"`
// }
