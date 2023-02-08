package routes

import (
	_ "ams-fantastic-auth/docs"

	"github.com/gofiber/fiber/v2"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	//swagger "github.com/arsmn/fiber-swagger/v2"
)

func Swagger(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", fiberSwagger.WrapHandler)
	//route.Get("*", swagger.HandlerDefault)
}
