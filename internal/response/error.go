package response

// ErroBody example
type ErroBody struct {
	Code    string `json:"code" example:"01"`
	Message string `json:"message" example:"status bad request"`
}
