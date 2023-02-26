package handlers

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/middleware"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Login
//
// @Summary Login.
// @Description Login.
// @Tags Login
// @Accept json
// @Produce json
// @Param user body model.Login true "Login infomation"
// @Success 200 {object} model.TokenResponse
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/login [post]
func Login(c *fiber.Ctx) error {
	loginInfo := new(model.Login)
	if parseErr := c.BodyParser(loginInfo); parseErr != nil {
		response.NewError(c, fiber.StatusBadRequest, parseErr)
		return parseErr
	}

	log.Printf("id: %v", loginInfo.ID)
	log.Printf("password: %v", loginInfo.Password)

	if len(loginInfo.ID) <= 0 {
		nilErr := errors.New("id is nil")
		response.NewError(c, fiber.StatusBadRequest, nilErr)
		return nilErr
	}
	if len(loginInfo.Password) <= 0 {
		nilErr := errors.New("password is nil")
		response.NewError(c, fiber.StatusBadRequest, nilErr)
		return nilErr
	}

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	user, selectErr := database.SelectUser(db, loginInfo.ID)
	if selectErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, selectErr)
		return selectErr
	}

	if len(user.ID) <= 0 {
		notFoundErr := errors.New("not found user")
		response.NewError(c, fiber.StatusBadRequest, notFoundErr)
		return notFoundErr
	}

	if user.Password != loginInfo.Password {
		passErr := errors.New("password wang")
		response.NewError(c, fiber.StatusBadRequest, passErr)
		return passErr
	}

	// create refresh token
	expiresAt := time.Now().Add(1 * time.Minute)
	newRefreshToken := &model.Token{UserID: user.ID, ExpiresAt: expiresAt}
	if insertErr := database.InsertToken(db, newRefreshToken); insertErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, insertErr)
		return insertErr
	}

	// create access token
	// expiresAt = time.Now().Add(5 * time.Minute).String()
	// newAccessToken := &model.Token{UUID: uuid.New().String(), UserID: user.ID, ExpiresAt: expiresAt}

	// output
	tokenRes := model.TokenResponse{
		RefreshToken: *newRefreshToken,
		//AccessToken:  *newAccessToken,
	}

	return c.JSON(tokenRes)
}

// Logout
//
// @Summary Logout.
// @Description Logout.
// @Tags Logout
// @Accept json
// @Produce json
// @Param tokon_uuid path string true "uuid of the token"
// @Success 204 {string} status "ok"
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Security ApiKeyAuth
// @Router /v1/logout/{tokon_uuid} [delete]
func Logout(c *fiber.Ctx) error {
	tokenUUID := c.Params("tokon_uuid")
	if len(tokenUUID) <= 0 {
		nilErr := errors.New("uuid is nil")
		response.NewError(c, fiber.StatusBadRequest, nilErr)
		return nilErr
	}
	log.Println("path tokon_uuid: ", tokenUUID)

	token, status, getErr := middleware.GetBeareToken(c)
	if getErr != nil {
		c.Status(status)
		return c.JSON(fiber.Map{
			"message": getErr.Error(),
		})
	}

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, newErr)
		return newErr
	}

	delErr := database.DeleteToken(db, token)
	if delErr != nil {
		response.NewError(c, fiber.StatusInternalServerError, delErr)
		return delErr
	}
	return c.SendStatus(fiber.StatusNoContent)
}
