package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthErrorCode int

const (
	AuthUnknownErrorCode AuthErrorCode = iota

	AuthBodyParseErrorCode
	AuthPasswordGenerateErrorCode
	AuthDuplicateDataErrorCode
	AuthInsertFailErrorCode
	AuthLoginReqNilErrorCode

	AuthSelectFailErrorCode
	AuthNotFoundUserErrorCode
	AuthInvalidPasswordErrorCode
	AuthTokenGenerateErrorCode
	AuthRefreshReqNilErrorCode

	AuthInvalidTokenErrorCode
	AuthNotFoundOwnUserErrorCode
	AuthUnauthorizedErrorCode
)

var ErrorCodes = [...]string{
	"unknown error.",

	"body parsing error.",
	"password genrate error.",
	"user with that id, email already exists.",
	"insult faild.", // "Something bad happened"
	"email or password is nil.",

	"select faild.",
	"user not found.",
	"invalid email or password.",
	"generating JWT Token failed.",
	"could not refresh access token.",

	"invalid token.",
	"the user belonging to this token no logger exists.",
	"you are not logged in.",
}

func (s AuthErrorCode) String() string {
	return ErrorCodes[s]
}

// AuthError example
func AuthError(ctx *fiber.Ctx, status int, code AuthErrorCode, detailedMsg string) error {
	msg := code.String()
	if len(detailedMsg) > 0 {
		msg = fmt.Sprintf("%v(%v)", msg, detailedMsg)
	}

	err := ErroBody{
		Code:    fmt.Sprintf("AU%02d", code),
		Message: msg,
	}

	return ctx.Status(status).JSON(err)
}
