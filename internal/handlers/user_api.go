package handlers

import (
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserAPI struct {
}

func Login(c *fiber.Ctx) error {
	log.Println("Signin")
	return nil
}

func Logout(c *fiber.Ctx) error {
	log.Println("Signout")
	return nil
}

// CreateUser
//
// @Summary Create a new user.
// @Description Create a new book.
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "users infomation"
// @Success 200 {object} model.User
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	if parseErr := c.BodyParser(user); parseErr != nil {
		response.NewError(c, fiber.StatusBadRequest, parseErr)
		return parseErr
	}

	log.Printf("id: %v", user.ID)
	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate.String())
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.QnA.Question)
	log.Printf("qna answer: %v", user.QnA.Answer)

	db, newErr := database.New()
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	insertErr := database.InsertUser(db, user)
	if insertErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, insertErr)
		return insertErr
	}

	return c.JSON(user)
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
// @Router /v1/users [get]
func GetUsers(c *fiber.Ctx) error {
	db, newErr := database.New()
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	users, selectErr := database.SelectUsers(db)
	if selectErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, selectErr)
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
		nilErr := errors.New("id is nil")
		response.NewError(c, fiber.StatusBadRequest, nilErr)
		return nilErr
	}
	log.Println("path id: ", id)

	db, newErr := database.New()
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	user, selectErr := database.SelectUser(db, id)
	if selectErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, selectErr)
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
// @Router /v1/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) <= 0 {
		nilErr := errors.New("id is nil")
		response.NewError(c, fiber.StatusBadRequest, nilErr)
		return nilErr
	}
	log.Println("path id: ", id)

	user := new(model.User)
	if parseErr := c.BodyParser(user); parseErr != nil {
		response.NewError(c, fiber.StatusBadRequest, parseErr)
		return parseErr
	}

	log.Printf("id: %v", user.ID)
	user.ID = id
	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate.String())
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.QnA.Question)
	log.Printf("qna answer: %v", user.QnA.Answer)

	db, newErr := database.New()
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	updateErr := database.UpdateUser(db, user)
	if updateErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, updateErr)
		return updateErr
	}

	return c.SendStatus(fiber.StatusCreated)
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
		nilErr := errors.New("id is nil")
		response.NewError(c, fiber.StatusBadRequest, nilErr)
		return nilErr
	}
	log.Println("path id: ", id)

	db, newErr := database.New()
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	updateErr := database.DeleteUser(db, id)
	if updateErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, updateErr)
		return updateErr
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
