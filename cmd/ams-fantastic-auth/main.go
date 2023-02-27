package main

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/database/schema"
	"ams-fantastic-auth/internal/middleware"
	"ams-fantastic-auth/internal/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	// GitCommit, BuildTime, Get infos at build time the golang.
	GitCommit string
	BuildTime string
)

func buildInfoPrint() {
	log.Printf("Build Information : %v at %v\n", GitCommit, BuildTime)
	log.Println("Started at :", time.Now().Format(time.RFC3339))
}

// @title AMS Fantastic Auth Swagger API
// @version 1.0
// @description This is a Test auth api server
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url https://github.com/neulsang
// @contact.email dgkwon90@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	buildInfoPrint()
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
		if initTableErr := schema.CreateUsersTable(db); initTableErr != nil {
			log.Fatal(initTableErr)
		}
		// Not Used
		// if initTableErr := schema.CreateTokenTable(db); initTableErr != nil {
		// 	log.Fatal(initTableErr)
		// }
	} else {
		log.Fatal(dbErr)
	}

	app := fiber.New(fiberConfig)

	// Middlewares.
	middleware.Fiber(app)

	// Swagger
	routes.Swagger(app)
	routes.AuthApi(app)
	routes.UserApi(app)

	log.Fatal(app.Listen(":9090"))
}
