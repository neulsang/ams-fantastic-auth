package model

import "time"

type Login struct {
	ID       string `json:"id" minLength:"1" maxLength:"36" example:"dgkwon90"`
	Password string `json:"password" minLength:"8" maxLength:"255" example:"test1234"`
}

type TokenResponse struct {
	AccessToken  Token `json:"access_token"`
	RefreshToken Token `json:"refresh_token"`
}

type Token struct {
	UUID      string    `json:"uuid"`
	UserID    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreateAt  string    `json:"create_at"`
}
