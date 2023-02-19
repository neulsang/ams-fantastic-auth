package routes

import (
	"ams-fantastic-auth/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	routeV1 := app.Group("/api/v1")

	routeV1.Post("/login", handlers.Login)
	routeV1.Delete("/logout", handlers.Logout)

	routeV1.Post("/users", handlers.CreateUser)
	routeV1.Get("/users", handlers.GetUsers)
	routeV1.Get("/users/:id", handlers.GetUser)
	routeV1.Put("/users/:id", handlers.UpdateUser)
	routeV1.Delete("/users/:id", handlers.DeleteUser)
}
