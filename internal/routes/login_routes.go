package routes

import (
	"ams-fantastic-auth/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func LoginApi(app *fiber.App) {
	routeV1 := app.Group("/api/v1")
	routeV1.Post("/login", handlers.Login)
	routeV1.Delete("/logout/:tokon_uuid", handlers.Logout)
	routeV1.Get("/me", handlers.Me)
	routeV1.Get("/refresh", handlers.Refresh)
}
