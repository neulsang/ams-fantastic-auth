package swagger

import (
	_ "ams-fantastic-auth/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Add(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
