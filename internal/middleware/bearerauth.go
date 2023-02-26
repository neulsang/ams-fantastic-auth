package middleware

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/response"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Config struct {
	// BodyKey defines the key to use when searching for the bearer token inside the
	// request's body.
	// Optional. Default: "access_token".
	BodyKey string

	// HeaderKey defines the prefix of the Authorization header's value, used when
	// searching for the bearer token inside the request's headers.
	// Optional. Default: "Bearer".
	HeaderKey string

	// QueryKey defines the key to use when searching for the bearer token inside the
	// request's query parameters.
	// Optional. Default: "access_token".
	QueryKey string

	// RequestKey defines the name of the local variable that will be created in the
	// request's context, which will contain the bearer token extracted from the
	// request.
	// Optional. Default: "token".
	RequestKey string
}

func BearerAuthNew(config *Config) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		log.Println("BearerAuthNew")
		log.Printf("Header: %v", c.GetReqHeaders())
		var token string
		var errored bool = false

		// query parameter
		queryValue := c.Query(config.QueryKey)
		if len(queryValue) > 0 {
			token = queryValue
			log.Println("queryValue: ", queryValue)
		}

		// body parameter
		bodyValue := c.Body()
		if len(bodyValue) > 0 {
			log.Println("bodyValue: ", bodyValue)
			if len(token) > 0 {
				errored = true
			}

			token = string(bodyValue)
		}

		// request authorization header
		headerValue := c.Get("Authorization")
		if len(headerValue) > 0 {
			log.Println("headerValue: ", headerValue)
			components := strings.SplitN(headerValue, " ", 2)
			log.Println("components: ", components)

			if len(components) == 2 && components[0] == config.HeaderKey {
				if len(token) > 0 {
					errored = true
				}
				token = components[1]
				log.Println("token: ", components)

			}
		}

		// check user test1234
		if "test1234" != token {
			errored = true
		}

		if errored {
			c.Status(401)
			return c.JSON(fiber.Map{
				"message": "Missing Authorization header",
			})
		} else {
			//ctx.Locals(config.RequestKey, token)
			return c.Next()
		}
	}
}

func GetBeareToken(c *fiber.Ctx) (string, int, error) {
	// Get the Authorization header value
	authHeader := c.Get("Authorization")

	// Check if the header is empty or not
	if len(authHeader) == 0 {
		// Return a 401 Unauthorized status code if the header is empty
		return "", fiber.StatusUnauthorized, errors.New("Missing Authorization header")
	}

	log.Println("authHeader: ", authHeader)

	// Check if the header starts with "Bearer "
	if authHeader[:7] != "Bearer " {
		// Return a 401 Unauthorized status code if the header is invalid
		return "", fiber.StatusUnauthorized, errors.New("Invalid Authorization header format")
	}

	// Get the token from the header
	token := authHeader[7:]

	log.Println("authHeader: ", token)
	return token, fiber.StatusOK, nil
}

// BearerAuthReq Middleware to handle bearer token authentication
func BearerAuthReq(c *fiber.Ctx) error {

	token, status, getErr := GetBeareToken(c)
	if getErr != nil {
		c.Status(status)
		return c.JSON(fiber.Map{
			"message": getErr.Error(),
		})
	}

	// case 1. token is tester
	// TODO: Token for tester. So you have to delete it later.
	if token == "ams-tester" {
		return c.Next()
	}

	// Find token in db
	db, newErr := database.New(configs.Database())
	if newErr != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error(" + newErr.Error() + ")",
		})
	}

	tokenInfo, selErr := database.SelectToken(db, token)
	if selErr != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Invalid token(" + selErr.Error() + ")",
		})
	}

	// case 2. not found
	if tokenInfo == nil {
		// Return a 401 Unauthorized status code if the token is invalid
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// case 3. time expire
	// TODO:

	// Continue with the request chain if the token is valid
	return c.Next()
}

// BearerAuthReq2 Middleware to handle bearer token authentication
func BearerAuthReq2(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		tokenString = c.Cookies("access_token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	accessJwtSercret := "AmsAccessJwtSecret"

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(accessJwtSercret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

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

	c.Locals("user", user)

	// Continue with the request chain if the token is valid
	return c.Next()
}
