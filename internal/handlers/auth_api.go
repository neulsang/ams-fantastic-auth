package handlers

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	jwttool "ams-fantastic-auth/pkg/jwt"
	"ams-fantastic-auth/pkg/password"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// RegisterUser
//
// @Summary Create a new user.
// @Description Create a new user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.User true "users infomation"
// @Success 201 {object} model.UserResponse
// @Router /api/v1/auth/register [post]
func RegisterUser(c *fiber.Ctx) error {
	user := new(model.User)
	if parseErr := c.BodyParser(user); parseErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, parseErr.Error())
	}

	log.Printf("nickName: %v", user.NickName)
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

	genPassword, genErr := password.Generate(user.Password)
	if genErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, genErr.Error())
	}
	user.Password = genPassword

	insertErr := database.InsertUser(db, user)
	if insertErr != nil && strings.Contains(insertErr.Error(), "duplicate key value violates unique") {
		return response.NewError(c, fiber.StatusConflict, "User with that id, email already exists")
	} else if insertErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, "Something bad happened")
	}

	userRes, selectErr := database.SelectUserByEmail(db, user.Email)
	if selectErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, "Something bad happened")
	}
	return c.Status(fiber.StatusCreated).JSON(userRes)
}

// Login
//
// @Summary Login.
// @Description Login.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.Login true "Login infomation"
// @Success 200 {object} model.Token
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	loginInfo := new(model.Login)
	if parseErr := c.BodyParser(loginInfo); parseErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, parseErr.Error())
	}

	log.Printf("email: %v", loginInfo.Email)
	log.Printf("password: %v", loginInfo.Password)

	if len(loginInfo.Email) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "email is nil")
	}

	if len(loginInfo.Password) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "password is nil")
	}

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	userPassword, selectErr := database.SelectUserPassword(db, loginInfo.Email)
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if len(userPassword) <= 0 {
		return response.NewError(c, fiber.StatusBadRequest, "not found user")
	}

	compareErr := password.CompareHashAndPassword(userPassword, loginInfo.Password)
	if compareErr != nil {
		return response.NewError(c, fiber.StatusBadRequest, "Invalid email or Password")
	}

	accessJwtExpiresIn := time.Minute * 3
	accessJwtSercret := "AmsAccessJwtSecret"
	accessJwtMaxAge := 60

	accessToken, signErr := jwttool.GenerateNewToken(accessJwtExpiresIn, accessJwtSercret, loginInfo.Email)
	if signErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, fmt.Sprintf("generating JWT Token failed: %v", signErr))
	}

	refreshJwtExpiresIn := time.Minute * 5
	refreshJwtSercret := "AmsRefreshJwtSecret"
	refreshJwtMaxAge := 60

	refreshToken, signErr := jwttool.GenerateNewToken(refreshJwtExpiresIn, refreshJwtSercret, loginInfo.Email)
	if signErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, fmt.Sprintf("generating JWT Token failed: %v", signErr))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

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

	tokenRes := model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(tokenRes)
}

// Logout
//
// @Summary Logout.
// @Description Logout.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/auth/logout [get]
func Logout(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "false",
		Expires: expired,
	})
	return c.SendStatus(fiber.StatusOK)
}

// Refresh
//
// @Summary Request a new access token.
// @Description Request a new access token.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.Token
// @Security ApiKeyAuth
// @Router /api/v1/auth/refresh [get]
func Refresh(c *fiber.Ctx) error {
	var refreshToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		refreshToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("refresh_token") != "" {
		refreshToken = c.Cookies("refresh_token")
	}

	if refreshToken == "" {
		return response.NewError(c, fiber.StatusUnauthorized, "could not refresh access token")
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
		return response.NewError(c, fiber.StatusUnauthorized, fmt.Sprintf("invalidate token(%v)", parseErr))
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return response.NewError(c, fiber.StatusUnauthorized, "invalid token claim")
	}

	// get user info
	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	user, selectErr := database.SelectUserByEmail(db, fmt.Sprint(claims["sub"]))
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if user.Email != claims["sub"] {

		return response.NewError(c, fiber.StatusForbidden, "the user belonging to this token no logger exists")
	}

	accessJwtExpiresIn := time.Minute * 3
	accessJwtSercret := "AmsAccessJwtSecret"
	accessJwtMaxAge := 60

	accessToken, signErr := jwttool.GenerateNewToken(accessJwtExpiresIn, accessJwtSercret, user.Email)
	if signErr != nil {
		return response.NewError(c, fiber.StatusBadGateway, fmt.Sprintf("generating JWT Token failed: %v", signErr))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   accessJwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	tokenRes := model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(tokenRes)
}
