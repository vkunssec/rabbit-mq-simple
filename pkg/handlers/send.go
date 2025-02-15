package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/rabbit-mq-simple/pkg/repository"

	_ "github.com/vkunssec/rabbit-mq-simple/docs"
)

// Payload é a estrutura que representa o payload da requisição
// @Description Payload da requisição
type Payload struct {
	Message    string `json:"message"`
	RoutingKey string `json:"routing_key"`
}

// Response é a estrutura que representa a resposta da requisição
// @Description Resposta da requisição
type Response struct {
	Message string `json:"message" example:"message sent"`
}

// BadRequestError é a estrutura que representa o erro de requisição
// @Description Erro de requisição
type BadRequestError struct {
	Error string `json:"error" example:"message is required"`
}

// @Summary Envia uma mensagem para o RabbitMQ
// @Description Endpoint para enviar mensagens ao RabbitMQ. Requer uma mensagem e uma routing key válida. A mensagem será publicada no exchange configurado usando a routing key especificada.
// @Tags Send Message RabbitMQ
// @Accept json
// @Produce json
// @Param payload body Payload true "Payload"
// @Success 200 {object} Response
// @Failure 400 {object} BadRequestError
// @Router /send [post]
func SendMessageRabbitMQHandler(ctx *fiber.Ctx) error {
	payload := new(Payload)
	if err := ctx.BodyParser(payload); err != nil || payload.Message == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(BadRequestError{Error: "message is required"})
	}

	if payload.RoutingKey == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(BadRequestError{Error: "routing key is required"})
	}

	err := repository.SendMessageRabbitMQ(payload.Message, payload.RoutingKey)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(BadRequestError{Error: "failed to send message"})
	}

	return ctx.Status(fiber.StatusOK).JSON(Response{Message: "message sent"})
}
