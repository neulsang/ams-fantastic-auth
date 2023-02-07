package main

import (
	"ams-fantastic-auth/internal/router"
	"ams-fantastic-auth/internal/swagger"
	"log"

	"github.com/gofiber/fiber/v2"
)

// @title AMS Fantastic Auth Swagger API
// @version 1.0
// @description This is a Test auth api server

// @contact.name Request permission of Example API
// @contact.url https://github.com/neulsang
// @contact.email dgkwon90@gmail.com

// @host localhost:9090
// @BasePath /api/v1
func main() {
	app := fiber.New()
	router.Add(app)
	swagger.Add(app)
	log.Fatal(app.Listen(":9090"))
}
