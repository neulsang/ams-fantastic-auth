package routes

import (
	"ams-fantastic-auth/internal/handlers"
	"ams-fantastic-auth/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserApi(app *fiber.App) {

	// BasicAuth
	//
	//routeV1 := app.Group("/api/v1", middleware.BasicAuthReq())

	// BasicAuthExt
	//
	//routeV1 := app.Group("/api/v1", middleware.BasicAuthExtReq())

	// NoneAuth
	//
	routeV1 := app.Group("/api/v1")
	routeV1.Get("/users/me", middleware.BearerAuthReq, handlers.Me)
	routeV1.Get("/users", middleware.BearerAuthReq, handlers.GetUsers)
	routeV1.Get("/users/:id", middleware.BearerAuthReq, handlers.GetUser)
	routeV1.Patch("/users/:id", middleware.BearerAuthReq, handlers.UpdateUser)
	routeV1.Delete("/users/:id", middleware.BearerAuthReq, handlers.DeleteUser)
}
