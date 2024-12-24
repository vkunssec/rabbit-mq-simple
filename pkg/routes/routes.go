package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/rabbit-mq-simple/pkg/handlers"

	_ "github.com/vkunssec/rabbit-mq-simple/docs"
)

// RegisterRoutes registra as rotas da aplicação
// @Summary Registra as rotas da aplicação
func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/send", handlers.SendMessageRabbitMQHandler)
}
