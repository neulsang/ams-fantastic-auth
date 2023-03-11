package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserErrorCode int

const (
	UserUnknownErrorCode UserErrorCode = iota

	UserBodyParseErrorCode
	UserPasswordGenerateErrorCode
	UserDuplicateDataErrorCode
	UserInsertFailErrorCode
	UserLoginReqNilErrorCode

	UserSelectFailErrorCode
	UserNotFoundUserErrorCode
	UserInvalidPasswordErrorCode
	UserTokenGenerateErrorCode
	UserRefreshReqNilErrorCode

	UserInvalidTokenErrorCode
	UserNotFoundOwnUserErrorCode
	UserUnauthorizedErrorCode
	UserReqParamErrorCode
	UserUpdateFailErrorCode

	UseDeleteFailErrorCode
)

var UserErrorCodes = [...]string{
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
	"invalid request param.",
	"update faild.",

	"delete faild.",
}

func (s UserErrorCode) String() string {
	return UserErrorCodes[s]
}

// UserError example
func UserError(ctx *fiber.Ctx, status int, code UserErrorCode, detailedMsg string) error {
	msg := code.String()
	if len(detailedMsg) > 0 {
		msg = fmt.Sprintf("%v(%v)", msg, detailedMsg)
	}

	err := ErroBody{
		Code:    fmt.Sprintf("US%02d", code),
		Message: msg,
	}

	return ctx.Status(status).JSON(err)
}
