package handler

import (
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/middleware"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	"ams-fantastic-auth/pkg/jwttool"
	"ams-fantastic-auth/pkg/password"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserAPI struct {
	db           *database.Database
	accessToken  *jwttool.Token
	refreshToken *jwttool.Token
}

func NewUser(db *database.Database, middleJwtAuth *middleware.JWTAuth) *UserAPI {
	return &UserAPI{
		db:           db,
		accessToken:  middleJwtAuth.AccessToken(),
		refreshToken: middleJwtAuth.RefreshToken(),
	}
}

// Me
//
// @Summary Get my user information.
// @Description Get my user information..
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.UserResponse
// @Security ApiKeyAuth
// @Router /api/v1/users/me [get]
func (u *UserAPI) Me(c *fiber.Ctx) error {
	var user model.UserResponse
	if c.Locals("user") == nil {
		return response.UserError(c, fiber.StatusNoContent, response.UserNotFoundUserErrorCode, "")
	}

	user = c.Locals("user").(model.UserResponse)
	return c.Status(fiber.StatusOK).JSON(user)
}

// GetUsers
//
// @Summary Get all exists users information.
// @Description Get all exists users information.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.UserResponse
// @Security ApiKeyAuth
// @Router /api/v1/users [get]
func (u *UserAPI) GetUsers(c *fiber.Ctx) error {
	users, selectErr := database.SelectUsers(u.db.Connect())
	if selectErr != nil {
		return response.UserError(c, fiber.StatusInternalServerError, response.UserSelectFailErrorCode, selectErr.Error())
	}

	if users == nil {
		return response.UserError(c, fiber.StatusNoContent, response.UserNotFoundUserErrorCode, "")
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUser
//
// @Summary Get user information by given ID.
// @Description Get user information by given ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 200 {object} model.UserResponse
// @Router /api/v1/users/{id} [get]
func (u *UserAPI) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		return response.UserError(c, fiber.StatusBadRequest, response.UserReqParamErrorCode, "")
	}
	log.Println("path id: ", id)

	user, selectErr := database.SelectUserById(u.db.Connect(), id)
	if selectErr != nil {
		return response.UserError(c, fiber.StatusInternalServerError, response.UserSelectFailErrorCode, selectErr.Error())
	}

	if user == nil {
		return response.UserError(c, fiber.StatusNoContent, response.UserNotFoundUserErrorCode, "")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser
//
// @Summary Update user information.
// @Description Update user information.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Param user body  model.User true "users infomation"
// @Success 200 {string} status "ok"
// @Router /api/v1/users/{id} [patch]
func (u *UserAPI) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		return response.UserError(c, fiber.StatusBadRequest, response.UserReqParamErrorCode, "")
	}
	log.Println("path id: ", id)

	user := new(model.User)
	if parseErr := c.BodyParser(user); parseErr != nil {
		return response.UserError(c, fiber.StatusBadRequest, response.UserBodyParseErrorCode, parseErr.Error())
	}

	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate)
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.Qna.Question)
	log.Printf("qna answer: %v", user.Qna.Answer)

	genPassword, genErr := password.Generate(user.Password)
	if genErr != nil {
		return response.UserError(c, fiber.StatusBadRequest, response.UserPasswordGenerateErrorCode, genErr.Error())
	}
	user.Password = genPassword

	updateErr := database.UpdateUser(u.db.Connect(), id, user)
	if updateErr != nil {
		return response.UserError(c, fiber.StatusInternalServerError, response.UserUpdateFailErrorCode, updateErr.Error())
	}

	return c.Status(fiber.StatusOK).SendString("ok")
}

// DeleteUser
//
// @Summary Delete user information by given ID.
// @Description Delete user information by given ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 204 {string} status "ok"
// @Router /api/v1/users/{id} [delete]
func (u *UserAPI) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		return response.UserError(c, fiber.StatusBadRequest, response.UserReqParamErrorCode, "")
	}
	log.Println("path id: ", id)

	deleteErr := database.DeleteUser(u.db.Connect(), id)
	if deleteErr != nil {
		return response.UserError(c, fiber.StatusInternalServerError, response.UseDeleteFailErrorCode, deleteErr.Error())
	}

	// Return status 204 no content.
	return c.Status(fiber.StatusNoContent).SendString("ok")
}
