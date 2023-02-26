package handlers

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/model"
	jwttool "ams-fantastic-auth/internal/pkg/jwt"
	"ams-fantastic-auth/internal/response"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
		return response.NewError(c, fiber.StatusBadRequest, parseErr.Error())
	}

	log.Printf("id: %v", loginInfo.ID)
	log.Printf("password: %v", loginInfo.Password)

	if len(loginInfo.ID) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "id is nil")
	}
	if len(loginInfo.Password) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "password is nil")
	}

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	user, selectErr := database.SelectUser(db, loginInfo.ID)
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if len(user.ID) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "not found user")
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if compareErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, "Invalid email or Password")
	}

	// OLD!!
	// // create refresh token
	// expiresAt := time.Now().Add(1 * time.Minute)
	// newRefreshToken := &model.Token{UserID: user.ID, ExpiresAt: expiresAt}
	// if insertErr := database.InsertToken(db, newRefreshToken); insertErr != nil {
	// 	return response.NewError(c, fiber.StatusInternalServerError, insertErr.Error())
	// }

	// // create access token
	// // expiresAt = time.Now().Add(5 * time.Minute).String()
	// // newAccessToken := &model.Token{UUID: uuid.New().String(), UserID: user.ID, ExpiresAt: expiresAt}

	// // output
	// tokenRes := model.TokenResponse{
	// 	RefreshToken: *newRefreshToken,
	// 	//AccessToken:  *newAccessToken,
	// }

	// return c.JSON(tokenRes)

	accessJwtExpiresIn := time.Minute * 1
	accessJwtSercret := "AmsAccessJwtSecret"
	accessJwtMaxAge := 60

	accessToken, signErr := jwttool.GenerateNewToken(accessJwtExpiresIn, accessJwtSercret, user.ID)
	if signErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, fmt.Sprintf("generating JWT Token failed: %v", signErr))
	}

	refreshJwtExpiresIn := time.Minute * 1
	refreshJwtSercret := "AmsRefreshJwtSecret"
	refreshJwtMaxAge := 60

	refreshToken, signErr := jwttool.GenerateNewToken(refreshJwtExpiresIn, refreshJwtSercret, user.ID)
	if signErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, fmt.Sprintf("generating JWT Token failed: %v", signErr))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   refreshJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": accessToken})
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
	// OLD
	// tokenUUID := c.Params("tokon_uuid")
	// if len(tokenUUID) <= 0 {
	// 	return response.NewError(c, fiber.StatusBadRequest, "uuid is nil")
	// }
	// log.Println("path tokon_uuid: ", tokenUUID)

	// token, status, getErr := middleware.GetBeareToken(c)
	// if getErr != nil {
	// 	c.Status(status)
	// 	return c.JSON(fiber.Map{
	// 		"message": getErr.Error(),
	// 	})
	// }

	// db, newErr := database.New(configs.Database())
	// if newErr != nil {
	// 	return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	// }

	// delErr := database.DeleteToken(db, token)
	// if delErr != nil {
	// 	return response.NewError(c, fiber.StatusInternalServerError, delErr.Error())
	// }
	//return c.SendStatus(fiber.StatusNoContent)

	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "false",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	return c.SendStatus(fiber.StatusOK)
}

// Me
//
// @Summary Get user's credentials.
// @Description Get the login user’s credentials.
// @Tags Me
// @Accept json
// @Produce json
// @Param user body model.Login true "Login infomation"
// @Success 200 {object} model.TokenResponse
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users/me [get]
func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	return c.Status(fiber.StatusOK).JSON(user)
}

// Refresh
//
// @Summary Request a new access token.
// @Description Request a new access token.
// @Tags Refresh
// @Accept json
// @Produce json
// @Param user body model.Login true "Login infomation"
// @Success 200 {object} model.TokenResponse
// @Failure	400	{object} response.HTTPError
// @Failure	404	{object} response.HTTPError
// @Failure	500	{object} response.HTTPError
// @Router /v1/users/refresh [get]
func Refresh(c *fiber.Ctx) error {
	var refreshToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		refreshToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("refresh_token") != "" {
		refreshToken = c.Cookies("refresh_token")
	}
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "could not refresh access token"})
	}
	// get sub
	refreshJwtSercret := "AmsRefreshJwtSecret"
	tokenByte, parseErr := jwt.Parse(refreshToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(refreshJwtSercret), nil
	})

	if parseErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", parseErr)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}
	// get user info
	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	user, selectErr := database.SelectUser(db, fmt.Sprint(claims["sub"]))
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if user.ID != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	accessJwtExpiresIn := time.Minute * 1
	accessJwtSercret := "AmsAccessJwtSecret"
	accessJwtMaxAge := 60

	accessToken, signErr := jwttool.GenerateNewToken(accessJwtExpiresIn, accessJwtSercret, user.ID)
	if signErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, fmt.Sprintf("generating JWT Token failed: %v", signErr))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": accessToken})
}
