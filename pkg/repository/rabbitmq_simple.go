package repository

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/vkunssec/rabbit-mq-simple/internal/config/environment"
)

// SendMessageRabbitMQ publica uma mensagem no RabbitMQ.
func SendMessageRabbitMQ(m string) error {
	environment.LoadEnv()
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, channelRabbitMQ, err := OpenConnectionRabbitMQ(amqpServerURL)
	if err != nil {
		return err
	}
	defer connectRabbitMQ.Close()

	// Cria uma mensagem para publicar
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(m),
	}

	// Tenta publicar a mensagem na fila
	if err := channelRabbitMQ.Publish(
		"ExchangeService1", // exchange
		"QueueService1",    // nome da fila
		false,              // mandatory
		false,              // immediate
		message,            // mensagem a ser publicada
	); err != nil {
		return err
	}

	return nil
}

// ReceiveMessageRabbitMQ recebe mensagens do RabbitMQ.
func ReceiveMessageRabbitMQ(channelRabbitMQ *amqp.Channel) {
	// Inscrevendo-se na QueueService1 para receber mensagens
	messages, err := channelRabbitMQ.Consume(
		"QueueService1", // nome da fila
		"",              // consumidor
		true,            // auto-ack
		false,           // exclusivo
		false,           // no local
		false,           // no wait
		nil,             // argumentos
	)
	if err != nil {
		log.Println(err)
	}

	// Constrói mensagem de boas-vindas
	log.Println("Conectado com sucesso ao RabbitMQ")
	log.Println("Aguardando mensagens")

	// Cria um canal para receber mensagens em loop infinito
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// Exemplo: mostra a mensagem recebida no console
			log.Printf(" > Mensagem recebida: %s\n", message.Body)
		}
	}()

	<-forever
}

// OpenConnectionRabbitMQ abre uma conexão com o RabbitMQ.
func OpenConnectionRabbitMQ(amqpServerURL string) (*amqp.Connection, *amqp.Channel, error) {
	// Cria uma nova conexão com o RabbitMQ
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return nil, nil, err
	}

	// Abre um canal na conexão que estabelecemos
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil, nil, err
	}

	// Declarar o Exchange
	err = channelRabbitMQ.ExchangeDeclare(
		"ExchangeService1", // nome
		"direct",           // tipo
		true,               // durável
		false,              // auto-deletável
		false,              // interno
		false,              // no-wait
		nil,                // argumentos
	)
	if err != nil {
		return nil, nil, err
	}

	// Declarar a Fila
	queue, err := channelRabbitMQ.QueueDeclare(
		"QueueService1", // nome
		true,            // durável
		false,           // auto-deletável
		false,           // exclusiva
		false,           // no-wait
		nil,             // argumentos
	)
	if err != nil {
		return nil, nil, err
	}

	// Criar o vínculo (binding) entre Exchange e Fila
	err = channelRabbitMQ.QueueBind(
		queue.Name,         // nome da fila
		"QueueService1",    // chave de roteamento
		"ExchangeService1", // nome do exchange
		false,              // no-wait
		nil,                // argumentos
	)
	if err != nil {
		return nil, nil, err
	}

	return connectRabbitMQ, channelRabbitMQ, nil
}
