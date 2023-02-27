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

	// BearerAuthNew
	//
	// routeV1 := app.Group("/api/v1", middleware.BearerAuthNew(&middleware.Config{
	// 	BodyKey:    "access_token",
	// 	HeaderKey:  "Bearer",
	// 	QueryKey:   "access_token",
	// 	RequestKey: "token",
	// }))

	// BearerAuthReq
	//
	routeV1 := app.Group("/api/v1")

	// NoneAuth
	//
	//routeV1 := app.Group("/api/v1")

	routeV1.Post("/users", handlers.CreateUser)
	routeV1.Get("/users", middleware.BearerAuthReq2, handlers.GetUsers)
	routeV1.Get("/users/:id", middleware.BearerAuthReq2, handlers.GetUser)
	routeV1.Patch("/users/:id", middleware.BearerAuthReq2, handlers.UpdateUser)
	routeV1.Delete("/users/:id", middleware.BearerAuthReq2, handlers.DeleteUser)
}
