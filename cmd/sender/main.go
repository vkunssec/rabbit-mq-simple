package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vkunssec/rabbit-mq-simple/pkg/middleware"
	"github.com/vkunssec/rabbit-mq-simple/pkg/routes"

	_ "github.com/vkunssec/rabbit-mq-simple/docs"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.RegisterRoutes(app)
	middleware.SwaggerMiddleware(app)

	log.Fatal(app.Listen(":3000"))
}
