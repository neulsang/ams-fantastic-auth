package model

type Login struct {
	Email    string `json:"email" maxLength:"255" example:"dgkwon90@naver.com"`
	Password string `json:"password" minLength:"8" maxLength:"255" example:"test1234"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
