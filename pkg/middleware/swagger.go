package middleware

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SwaggerMiddleware(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
