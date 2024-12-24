package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/rabbit-mq-simple/pkg/repository"

	_ "github.com/vkunssec/rabbit-mq-simple/docs"
)

// Payload é a estrutura que representa o payload da requisição
// @Description Payload da requisição
type Payload struct {
	Message string `json:"message"`
}

// Response é a estrutura que representa a resposta da requisição
// @Description Resposta da requisição
type Response struct {
	Message string `json:"message" example:"message sent"`
}

// @Summary Envia uma mensagem para o RabbitMQ
// @Description Envia uma mensagem para o RabbitMQ
// @Tags send
// @Accept json
// @Produce json
// @Param payload body Payload true "Payload"
// @Success 200 {object} Response
// @Router /send [post]
func SendMessageRabbitMQHandler(ctx *fiber.Ctx) error {
	payload := new(Payload)
	if err := ctx.BodyParser(payload); err != nil || payload.Message == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "message is required"})
	}

	repository.SendMessageRabbitMQ(payload.Message)
	return ctx.Status(fiber.StatusOK).JSON(Response{Message: "message sent"})
}