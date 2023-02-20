package handlers

import "github.com/gofiber/fiber/v2"

// Login
//
// @Summary Login.
// @Description Login.
// @Tags Login
// @Accept json
// @Produce json
// @Param user body model.Login true "Login infomation"
// @Success 200 {object} model.User
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/Login [post]
func Login(c *fiber.Ctx) error {

	return nil
}

// Logout
//
// @Summary Logout.
// @Description Logout.
// @Tags Logout
// @Accept json
// @Produce json
// @Param user body model.Login true "Logout infomation"
// @Success 200 {object} model.User
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/Logout [post]
func Logout(c *fiber.Ctx) error {
	return nil
}
