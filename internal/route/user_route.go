package route

import (
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/handler"
	"ams-fantastic-auth/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddUserRoute(app *fiber.App, db *database.Database, middleJwtAuth *middleware.JWTAuth) {
	routeV1 := app.Group("/api/v1")

	userApiHandler := handler.NewUser(db, middleJwtAuth)
	routeV1.Get("/users/me", middleJwtAuth.BearerAuthReq, userApiHandler.Me)
	routeV1.Get("/users", middleJwtAuth.BearerAuthReq, userApiHandler.GetUsers)
	routeV1.Get("/users/:id", middleJwtAuth.BearerAuthReq, userApiHandler.GetUser)
	routeV1.Patch("/users/:id", middleJwtAuth.BearerAuthReq, userApiHandler.UpdateUser)
	routeV1.Delete("/users/:id", middleJwtAuth.BearerAuthReq, userApiHandler.DeleteUser)
}
