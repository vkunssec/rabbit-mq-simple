package repository

import (
	"log"
	"os"

	"github.com/vkunssec/rabbit-mq-simple/internal/config/environment"
	rabbitmq "github.com/vkunssec/rabbit-mq-simple/pkg/domain/rabbitmq"
)

const (
	ExchangeName = "ExchangeService1" // ExchangeName é o nome do exchange
	ExchangeType = "direct"           // ExchangeType é o tipo de exchange

	QueueName1 = "QueueService1" // QueueName1 é o nome da fila 1
	QueueName2 = "QueueService2" // QueueName2 é o nome da fila 2

	RoutingKey1 = "route.service1" // RoutingKey1 é a chave de rota para a fila 1
	RoutingKey2 = "route.service2" // RoutingKey2 é a chave de rota para a fila 2
)

// SendMessageRabbitMQ publica uma mensagem no RabbitMQ com uma chave de rota específica.
func SendMessageRabbitMQ(m string, routingKey string) error {
	environment.LoadEnv()
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Criar nova instância do RabbitMQ
	rabbitmq, err := rabbitmq.New(amqpServerURL)
	if err != nil {
		return err
	}
	defer rabbitmq.Close()

	// Declarar Exchange
	if err := rabbitmq.ExchangeDeclare(ExchangeName, ExchangeType); err != nil {
		return err
	}

	// Publicar mensagem
	message := rabbitmq.PublishingMessage(m)
	if err := rabbitmq.Publish(ExchangeName, routingKey, message); err != nil {
		return err
	}

	return nil
}

// ReceiveMessageFromQueue1 recebe mensagens da fila 1.
func ReceiveMessageFromQueue1() error {
	return receiveMessages(QueueName1, RoutingKey1)
}

// ReceiveMessageFromQueue2 recebe mensagens da fila 2.
func ReceiveMessageFromQueue2() error {
	return receiveMessages(QueueName2, RoutingKey2)
}

// receiveMessages é uma função auxiliar para configurar e receber mensagens de uma fila específica.
func receiveMessages(queueName, routingKey string) error {
	environment.LoadEnv()
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Criar nova instância do RabbitMQ
	rabbitmq, err := rabbitmq.New(amqpServerURL)
	if err != nil {
		return err
	}
	defer rabbitmq.Close()

	// Declarar Exchange
	if err := rabbitmq.ExchangeDeclare(ExchangeName, ExchangeType); err != nil {
		return err
	}

	// Declarar Queue
	queue, err := rabbitmq.QueueDeclare(queueName)
	if err != nil {
		return err
	}

	// Vincular fila ao exchange
	if err := rabbitmq.QueueBind(queue.Name, ExchangeName, routingKey); err != nil {
		return err
	}

	// Consumir mensagens
	messages, err := rabbitmq.Consume(queueName, "")
	if err != nil {
		return err
	}

	log.Printf("Conectado com sucesso ao RabbitMQ - Escutando na fila: %s\n", queueName)
	log.Printf("Aguardando mensagens com routing key: %s\n", routingKey)

	// Criar canal para aguardar mensagens indefinidamente
	forever := make(chan bool)

	// Iniciar goroutine para consumir mensagens
	go func() {
		for message := range messages {
			log.Printf("[%s] Mensagem recebida: %s\n", queueName, message.Body)
		}
	}()

	// Aguardar mensagens indefinidamente
	<-forever

	return nil
}
