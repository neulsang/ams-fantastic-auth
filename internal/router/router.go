package router

import (
	"ams-fantastic-auth/internal/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

func middleware(c *fiber.Ctx) error {
	log.Println("middleware")
	return c.Next()
}

func Add(app *fiber.App) {
	api := app.Group("/api", middleware)

	v1 := api.Group("/v1", middleware) // /api/v1

	v1.Post("/signin", handler.Signin)
	v1.Delete("/signout", handler.Signout)

	v1.Post("/signup", handler.Singup)
	v1.Get("/users", handler.GetUsers)
	v1.Get("/users/:id", handler.GetUser)
	v1.Put("/users/:id", handler.PutUser)
	v1.Delete("/users/:id", handler.DeleteUser)
}
