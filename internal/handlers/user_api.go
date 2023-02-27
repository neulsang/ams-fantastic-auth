package handlers

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	"ams-fantastic-auth/pkg/password"
	"log"

	"github.com/gofiber/fiber/v2"
)

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
func Me(c *fiber.Ctx) error {
	var user model.UserResponse
	if c.Locals("user") == nil {
		return response.NewError(c, fiber.StatusNoContent, "data is nil")
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
func GetUsers(c *fiber.Ctx) error {
	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	users, selectErr := database.SelectUsers(db)
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if users == nil {
		return response.NewError(c, fiber.StatusNoContent, "data is nil")
	}

	return c.JSON(users)
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
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "id is nil")
	}
	log.Println("path id: ", id)

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	user, selectErr := database.SelectUserById(db, id)
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if user == nil {
		return response.NewError(c, fiber.StatusNoContent, "data is nil")
	}

	return c.JSON(user)
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
// @Success 201 {string} status "ok"
// @Router /api/v1/users/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "id is nil")
	}
	log.Println("path id: ", id)

	user := new(model.User)
	if parseErr := c.BodyParser(user); parseErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, parseErr.Error())
	}

	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate)
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.QnA.Question)
	log.Printf("qna answer: %v", user.QnA.Answer)

	genPassword, genErr := password.Generate(user.Password)
	if genErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, genErr.Error())
	}
	user.Password = genPassword

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	updateErr := database.UpdateUser(db, id, user)
	if updateErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, updateErr.Error())
	}

	return c.SendStatus(fiber.StatusOK)
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
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "id is nil")
	}
	log.Println("path id: ", id)

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	updateErr := database.DeleteUser(db, id)
	if updateErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, updateErr.Error())
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
