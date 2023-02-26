package handlers

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser
//
// @Summary Create a new user.
// @Description Create a new book.
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "users infomation"
// @Success 201 {object} model.User
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	if parseErr := c.BodyParser(user); parseErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, parseErr.Error())
	}

	log.Printf("id: %v", user.ID)
	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate)
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.QnA.Question)
	log.Printf("qna answer: %v", user.QnA.Answer)

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, hashErr.Error())
	}
	user.Password = string(hashedPassword)

	insertErr := database.InsertUser(db, user)
	if insertErr != nil && strings.Contains(insertErr.Error(), "duplicate key value violates unique") {
		return response.NewError(c, fiber.StatusConflict, "User with that email already exists")
	} else if insertErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, "Something bad happened")
	}

	if user == nil {
		return response.NewError(c, fiber.StatusNoContent, "data is nil")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUsers
//
// @Summary Get all exists users.
// @Description Get all exists users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Security ApiKeyAuth
// @Router /v1/users [get]
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
// @Summary Get user by given ID.
// @Description Get user by given ID.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 200 {object} model.User
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users/{id} [get]
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

	user, selectErr := database.SelectUser(db, id)
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
// @Summary Update user.
// @Description Update user.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Param user body  model.User true "users infomation"
// @Success 201 {string} status "ok"
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users/{id} [patch]
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

	log.Printf("id: %v", user.ID)
	user.ID = id
	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate)
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.QnA.Question)
	log.Printf("qna answer: %v", user.QnA.Answer)

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	updateErr := database.UpdateUser(db, user)
	if updateErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, updateErr.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

// DeleteUser
//
// @Summary Delete user by given ID.
// @Description Delete user by given ID.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 204 {string} status "ok"
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users/{id} [delete]
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
