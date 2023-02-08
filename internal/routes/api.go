package routes

import (
	"ams-fantastic-auth/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	routeV1 := app.Group("/api/v1")

	routeV1.Post("/signin", handlers.Signin)
	routeV1.Delete("/signout", handlers.Signout)

	routeV1.Post("/signup", handlers.Singup)

	routeV1.Get("/users", handlers.GetUsers)

	routeV1.Get("/users/:id", handlers.GetUser)
	routeV1.Put("/users/:id", handlers.PutUser)
	routeV1.Delete("/users/:id", handlers.DeleteUser)
}
