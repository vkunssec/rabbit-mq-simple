package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vkunssec/rabbit-mq-simple/pkg/middleware"
	"github.com/vkunssec/rabbit-mq-simple/pkg/routes"

	_ "github.com/vkunssec/rabbit-mq-simple/docs"
)

// @Title RabbitMQ Simple
// @Version 0.0.1
// @Description Este é um projeto de exemplo que demonstra a implementação de um sistema de mensageria usando RabbitMQ e Go
// @TermsOfService http://swagger.io/terms/

// @Contact.name API Support

// @License.name Apache 2.0
// @License.url http://www.apache.org/licenses/LICENSE-2.0.html

// @Host localhost:3000
// @BasePath /
// @Schemes http https
func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.RegisterRoutes(app)
	middleware.SwaggerMiddleware(app)

	log.Fatal(app.Listen(":3000"))
}
