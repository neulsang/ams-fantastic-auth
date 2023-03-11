package main

import (
	"ams-fantastic-auth/internal/config"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/middleware"
	"ams-fantastic-auth/internal/route"
	"ams-fantastic-auth/internal/server"
	"ams-fantastic-auth/pkg/beautiprint"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	// GitCommit, BuildTime, Get infos at build time the golang.
	GitCommit string
	BuildTime string

	shutdowns []func() error
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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {

	// Print build information to the console.
	buildInfoPrint()

	// Print Logo to the console.
	beautiprint.Logo("AMS Fantastic-Auth")

	// Load env file .
	err := godotenv.Load()
	if err != nil {
		log.Printf("could not load .env file: %v", err)
	}

	// Load Config
	var (
		cfg      = config.LoadConfig()
		shutdown = make(chan struct{})
	)

	log.Println("server: ", cfg.Server())
	log.Println("logger: ", cfg.Logger())
	log.Println("rdb: ", cfg.RDB())
	log.Println("jwt: ", cfg.JWT())

	// database init.
	db, dbErr := database.Open(cfg.RDB())
	if dbErr == nil {
		if initTablesErr := db.InitTables(); initTablesErr != nil {
			log.Fatal(initTablesErr)
		}
	} else {
		log.Fatal(dbErr)
	}

	// new middleTwtAuth
	middleJwtAuth := middleware.New(db, cfg.JWT())

	authServer := server.New(cfg.Server())

	route.AddSwaggerRotue(authServer)
	route.AddAuthRoute(authServer, db, middleJwtAuth)
	route.AddUserRoute(authServer, db, middleJwtAuth)

	shutdowns = append(shutdowns, db.Shutdown)

	go gracefulShutdown(authServer, shutdown)

	if err := authServer.Listen(fmt.Sprintf(":%v", cfg.Server().Port)); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func buildInfoPrint() {
	log.Printf("Build Information : %v at %v\n", GitCommit, BuildTime)
	log.Println("Started at :", time.Now().Format(time.RFC3339))
}

func gracefulShutdown(server *fiber.App, shutdown chan struct{}) {
	var (
		sigint = make(chan os.Signal, 1)
	)

	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	s := <-sigint

	// defalut log
	log.Println("got system signal:", s)
	log.Println("shutting down server gracefully")

	if err := server.Shutdown(); err != nil {
		log.Println("shutdown err: ", err)
	}

	log.Println("shutdown server success")

	log.Println("shutting down other modules")
	// close any other modules.
	for i := range shutdowns {
		shutdowns[i]()
	}
	log.Println("shutdown other modules success")
	close(shutdown)
}
