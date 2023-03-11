package model

// User Request struct
type User struct {
	Email       string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name        string `json:"name" maxLength:"255" example:"권대근"`
	NickName    string `json:"nickName" maxLength:"36" example:"dgkwon90"`
	PhoneNumber string `json:"phoneNumber" maxLength:"20" example:"01012344321"`
	BirthDate   string `json:"birth_date" example:"1990-07-29"`
	Gender      string `json:"gender" enums:"male, female, other" example:"male"`
	Password    string `json:"password" maxLength:"255" example:"test1234"`
	Qna         Qna    `json:"qna"`
}

// Question and Answer struct
type Qna struct {
	Question string `json:"question" maxLength:"255" example:"What is my favorite color?"`
	Answer   string `json:"answer" maxLength:"255" example:"blue"`
}

// User Response struct
type UserResponse struct {
	ID          string `json:"id" example:"{uuid}"`
	Email       string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name        string `json:"name" maxLength:"255" example:"권대근"`
	NickName    string `json:"nickName" maxLength:"36" example:"dgkwon90"`
	PhoneNumber string `json:"phoneNumber" maxLength:"20" example:"01012344321"`
	BirthDate   string `json:"birthDate" example:"1990-07-29"`
	Gender      string `json:"gender" enums:"male, female, other" example:"male"`
	Qna         Qna    `json:"qna"`
	CreatedAt   string `json:"created_at" example:"2023-02-27 17:03:20"`
	UpdatedAt   string `json:"updated_at" example:"2023-02-27 17:03:20"`
	DeletedAt   string `json:"deleted_at" example:"2023-02-27 17:03:20"`
}

// Register Request struct
type RegisterRequest struct {
	Email       string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name        string `json:"name" maxLength:"255" example:"권대근"`
	NickName    string `json:"nickName" maxLength:"36" example:"dgkwon90"`
	PhoneNumber string `json:"phoneNumber" maxLength:"20" example:"01012344321"`
	BirthDate   string `json:"birth_date" example:"1990-07-29"`
	Gender      string `json:"gender" enums:"male, female, other" example:"male"`
	Password    string `json:"password" maxLength:"255" example:"test1234"`
	Qna         Qna    `json:"qna"`
}
