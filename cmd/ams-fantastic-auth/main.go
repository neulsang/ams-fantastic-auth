package main

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/middleware"
	"ams-fantastic-auth/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title AMS Fantastic Auth Swagger API
// @version 1.0
// @description This is a Test auth api server

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/neulsang
// @contact.email dgkwon90@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("could not load .env file: %v", err)
	}

	// Define Fiber fiberConfig.
	fiberConfig := configs.Fiber()
	// Define Database DatabaseConfig.
	databaseConfig := configs.Database()

	// database init.
	if db, dbErr := database.New(databaseConfig); dbErr == nil {
		if initTableErr := database.CreateUsersTable(db); initTableErr != nil {
			log.Fatal(initTableErr)
		}
	} else {
		log.Fatal(dbErr)
		//log.Println(dbErr)
	}

	app := fiber.New(fiberConfig)

	// Middlewares.
	middleware.Fiber(app)

	// Swagger
	routes.Swagger(app)
	routes.LoginApi(app)
	routes.UserApi(app)

	log.Fatal(app.Listen(":9090"))
}
