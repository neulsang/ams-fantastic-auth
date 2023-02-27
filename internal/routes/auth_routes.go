package routes

import (
	"ams-fantastic-auth/internal/handlers"
	"ams-fantastic-auth/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthApi(app *fiber.App) {
	routeV1 := app.Group("/api/v1")

	routeV1.Post("/auth/register", handlers.RegisterUser)
	routeV1.Post("/auth/login", handlers.Login)
	routeV1.Get("/auth/refresh", handlers.Refresh)
	routeV1.Get("/auth/logout", middleware.BearerAuthReq, handlers.Logout)
}
