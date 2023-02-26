package response

import "github.com/gofiber/fiber/v2"

// NewError example
func NewError(ctx *fiber.Ctx, status int, message string) error {
	er := HTTPError{
		Code:    status,
		Message: message,
	}
	return ctx.Status(status).JSON(er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
