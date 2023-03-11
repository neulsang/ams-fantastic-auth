package model

type LoginRequest struct {
	Email    string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Password string `json:"password" minLength:"8" maxLength:"255" example:"test1234"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken" example:"{uuid}"`
	RefreshToken string `json:"refreshToken" example:"{uuid}"`
	ID           string `json:"id" example:"{uuid}"`
	Email        string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Name         string `json:"name" maxLength:"255" example:"권대근"`
	NickName     string `json:"nickName" maxLength:"36" example:"dgkwon90"`
	PhoneNumber  string `json:"phoneNumber" maxLength:"20" example:"01012344321"`
	BirthDate    string `json:"birth_date" example:"1990-07-29"`
	Gender       string `json:"gender" enums:"male, female, other" example:"male"`
	Qna          Qna    `json:"qna"`
}

type RefreshResponse struct {
	AccessToken  string `json:"accessToken" example:"{uuid}"`
	RefreshToken string `json:"refreshToken" example:"{uuid}"`
}
