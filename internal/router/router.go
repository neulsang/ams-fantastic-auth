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

	auth := v1.Group("/auth", middleware) // /api/v1/auth/
	auth.Post("/signin", handler.Signin)
	auth.Delete("/signout", handler.Signout)

	v1.Post("/signup", handler.Singup)
	v1.Get("/users", handler.Users)
	v1.Get("/users/:id", handler.User)
	v1.Put("/users/:id", handler.PutUser)
	v1.Delete("/users/:id", handler.DeleteUser)
}
