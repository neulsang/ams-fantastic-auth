package route

import (
	_ "ams-fantastic-auth/docs"

	"github.com/gofiber/fiber/v2"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func AddSwaggerRotue(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", fiberSwagger.WrapHandler)
}
