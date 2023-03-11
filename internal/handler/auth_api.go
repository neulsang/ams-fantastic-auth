package handler

import (
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/middleware"
	"ams-fantastic-auth/internal/model"
	"ams-fantastic-auth/internal/response"
	"ams-fantastic-auth/pkg/password"
	"fmt"
	"log"
	"strings"
	"time"

	"ams-fantastic-auth/pkg/jwttool"

	"github.com/gofiber/fiber/v2"
)

type AuthAPI struct {
	db           *database.Database
	accessToken  *jwttool.Token
	refreshToken *jwttool.Token
}

func NewAuth(db *database.Database, middleJwtAuth *middleware.JWTAuth) *AuthAPI {
	return &AuthAPI{
		db:           db,
		accessToken:  middleJwtAuth.AccessToken(),
		refreshToken: middleJwtAuth.RefreshToken(),
	}
}

// RegisterUser
//
// @Summary Create a new user.
// @Description Create a new user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "users infomation"
// @Success 201 {string} status "ok"
// @Router /api/v1/auth/register [post]
func (a *AuthAPI) RegisterUser(c *fiber.Ctx) error {
	user := new(model.RegisterRequest)
	if parseErr := c.BodyParser(user); parseErr != nil {
		return response.AuthError(c, fiber.StatusBadRequest, response.AuthBodyParseErrorCode, parseErr.Error())
	}

	log.Printf("nickName: %v", user.NickName)
	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("birtDate: %v", user.BirthDate)
	log.Printf("gender: %v", user.Gender)
	log.Printf("password: %v", user.Password)
	log.Printf("qna question: %v", user.Qna.Question)
	log.Printf("qna answer: %v", user.Qna.Answer)

	genPassword, genErr := password.Generate(user.Password)
	if genErr != nil {
		return response.AuthError(c, fiber.StatusBadRequest, response.AuthPasswordGenerateErrorCode, genErr.Error())
	}
	user.Password = genPassword

	insertErr := database.InsertUser(a.db.Connect(), user)
	if insertErr != nil && strings.Contains(insertErr.Error(), "duplicate key value violates unique") {
		return response.AuthError(c, fiber.StatusConflict, response.AuthDuplicateDataErrorCode, "")
	} else if insertErr != nil {
		return response.AuthError(c, fiber.StatusBadGateway, response.AuthInsertFailErrorCode, insertErr.Error())
	}
	return c.Status(fiber.StatusCreated).SendString("ok")
}

// Login
//
// @Summary Login.
// @Description Login.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.LoginRequest true "Login infomation"
// @Success 200 {object} model.LoginResponse
// @Router /api/v1/auth/login [post]
func (a *AuthAPI) Login(c *fiber.Ctx) error {
	loginInfo := new(model.LoginRequest)
	if parseErr := c.BodyParser(loginInfo); parseErr != nil {
		return response.AuthError(c, fiber.StatusBadRequest, response.AuthBodyParseErrorCode, parseErr.Error())
	}

	log.Printf("email: %v", loginInfo.Email)
	log.Printf("password: %v", loginInfo.Password)

	if len(loginInfo.Email) <= 0 || len(loginInfo.Password) <= 0 {
		return response.AuthError(c, fiber.StatusBadRequest, response.AuthLoginReqNilErrorCode, "")
	}

	userPassword, selectErr := database.SelectUserPassword(a.db.Connect(), loginInfo.Email)

	if selectErr != nil && strings.Contains(selectErr.Error(), "no rows in result set") {
		return response.AuthError(c, fiber.StatusInternalServerError, response.AuthNotFoundUserErrorCode, "")
	} else if selectErr != nil {
		return response.AuthError(c, fiber.StatusInternalServerError, response.AuthSelectFailErrorCode, selectErr.Error())
	}

	if len(userPassword) <= 0 {
		return response.AuthError(c, fiber.StatusBadRequest, response.AuthNotFoundUserErrorCode, "")
	}

	compareErr := password.CompareHashAndPassword(userPassword, loginInfo.Password)
	if compareErr != nil {
		return response.AuthError(c, fiber.StatusBadRequest, response.AuthInvalidPasswordErrorCode, "")
	}

	accessToken, signErr := a.accessToken.GenerateNewToken(loginInfo.Email)
	if signErr != nil {
		return response.AuthError(c, fiber.StatusBadGateway, response.AuthTokenGenerateErrorCode, signErr.Error())
	}

	refreshToken, signErr := a.refreshToken.GenerateNewToken(loginInfo.Email)
	if signErr != nil {
		return response.AuthError(c, fiber.StatusBadGateway, response.AuthTokenGenerateErrorCode, signErr.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   a.accessToken.Maxage * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   a.accessToken.Maxage * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   a.refreshToken.Maxage * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	userRes, selectErr := database.SelectUserByEmail(a.db.Connect(), loginInfo.Email)
	if selectErr != nil {
		return response.AuthError(c, fiber.StatusBadGateway, response.AuthSelectFailErrorCode, selectErr.Error())
	}

	tokenRes := model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ID:           userRes.ID,
		Email:        userRes.Email,
		Name:         userRes.Name,
		NickName:     userRes.NickName,
		PhoneNumber:  userRes.PhoneNumber,
		BirthDate:    userRes.BirthDate,
		Gender:       userRes.Gender,
		Qna:          userRes.Qna,
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
func (a *AuthAPI) Logout(c *fiber.Ctx) error {
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
	return c.Status(fiber.StatusOK).SendString("ok")
}

// Refresh
//
// @Summary Request a new access token.
// @Description Request a new access token.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.RefreshResponse
// @Security ApiKeyAuth
// @Router /api/v1/auth/refresh [get]
func (a *AuthAPI) Refresh(c *fiber.Ctx) error {
	var refreshToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		refreshToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("refresh_token") != "" {
		refreshToken = c.Cookies("refresh_token")
	}

	if len(refreshToken) <= 0 {
		return response.AuthError(c, fiber.StatusUnauthorized, response.AuthRefreshReqNilErrorCode, "")
	}

	// valid
	mapClaims, validErr := a.refreshToken.ValidToken(refreshToken)
	if validErr != nil {
		return response.AuthError(c, fiber.StatusUnauthorized, response.AuthInvalidTokenErrorCode, validErr.Error())
	}

	// get user info
	user, selectErr := database.SelectUserByEmail(a.db.Connect(), fmt.Sprint(mapClaims["sub"]))
	if selectErr != nil {
		return response.AuthError(c, fiber.StatusInternalServerError, response.AuthSelectFailErrorCode, selectErr.Error())
	}

	if user.Email != mapClaims["sub"] {
		return response.AuthError(c, fiber.StatusForbidden, response.AuthNotFoundOwnUserErrorCode, "")
	}

	accessToken, signErr := a.accessToken.GenerateNewToken(user.Email)
	if signErr != nil {
		return response.AuthError(c, fiber.StatusBadGateway, response.AuthTokenGenerateErrorCode, signErr.Error())
	}

	refreshToken, signErr = a.refreshToken.GenerateNewToken(user.Email)
	if signErr != nil {
		return response.AuthError(c, fiber.StatusBadGateway, response.AuthTokenGenerateErrorCode, signErr.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   a.accessToken.Maxage * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   a.accessToken.Maxage * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   a.refreshToken.Maxage * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})
	tokenRes := model.RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(tokenRes)
}
