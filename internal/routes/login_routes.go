package routes

import (
	"ams-fantastic-auth/internal/handlers"
	"ams-fantastic-auth/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func LoginApi(app *fiber.App) {
	routeV1 := app.Group("/api/v1")
	routeV1.Post("/login", handlers.Login)
	routeV1.Delete("/logout/:tokon_uuid", middleware.BearerAuthReq2, handlers.Logout)
	routeV1.Get("/me", middleware.BearerAuthReq2, handlers.Me)
	routeV1.Get("/refresh", middleware.BearerAuthReq2, handlers.Refresh)
}
