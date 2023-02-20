package configs

import (
	"ams-fantastic-auth/pkg/env"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func Fiber() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount := env.GetAsInt("SERVER_READ_TIMEOUT", 0)

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
