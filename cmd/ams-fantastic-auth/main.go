package main

import (
	"ams-fantastic-auth/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.Add(app)
	log.Fatal(app.Listen(":9090"))
}
