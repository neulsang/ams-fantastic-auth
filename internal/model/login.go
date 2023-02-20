package model

type Login struct {
	ID       string `json:"id" maxLength:"36" example:"dgkwon90"`
	Password string `json:"password" maxLength:"255" example:"test1234"`
}
