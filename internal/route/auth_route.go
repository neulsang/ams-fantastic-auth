package route

import (
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/handler"
	"ams-fantastic-auth/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddAuthRoute(app *fiber.App, db *database.Database, middleJwtAuth *middleware.JWTAuth) {
	routeV1 := app.Group("/api/v1")

	authApiHansler := handler.NewAuth(db, middleJwtAuth)
	routeV1.Post("/auth/register", authApiHansler.RegisterUser)
	routeV1.Post("/auth/login", authApiHansler.Login)
	routeV1.Get("/auth/refresh", authApiHansler.Refresh)
	routeV1.Get("/auth/logout", middleJwtAuth.BearerAuthReq, authApiHansler.Logout)
}
