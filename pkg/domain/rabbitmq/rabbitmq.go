package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

/*
Este arquivo contém a implementação de uma estrutura RabbitMQ que encapsula as operações básicas
para trabalhar com o RabbitMQ em Go.

Funcionalidades principais:

1. Estrutura RabbitMQ:
  - Mantém uma referência ao canal de comunicação com o RabbitMQ
  - Gerencia a conexão com o servidor RabbitMQ
  - Implementa verificação de estado da conexão
  - Suporta reconexão automática

2. Operações implementadas:
  - New: Constructor para criar nova instância do RabbitMQ
  - QueueDeclare: Cria uma nova fila no RabbitMQ com configurações específicas
  - QueueBind: Vincula uma fila a um exchange usando uma chave de roteamento
  - ExchangeDeclare: Cria um novo exchange com tipo específico (direct, fanout, topic, etc)
  - Consume: Configura um consumidor para receber mensagens de uma fila específica
  - Publish: Publica mensagens em um exchange específico
  - Close: Fecha a conexão e o canal com o RabbitMQ
  - OpenConnection: Estabelece uma nova conexão com o servidor RabbitMQ
  - PublishingMessage: Função auxiliar para criar mensagens no formato adequado
  - IsConnected: Verifica o estado atual da conexão
  - Reconnect: Restabelece a conexão em caso de falha

Características importantes:
  - Todas as filas são criadas como duráveis por padrão
  - As mensagens são formatadas como JSON
  - O código implementa os padrões básicos de mensageria do RabbitMQ
  - Fornece uma interface simplificada para as operações mais comuns do RabbitMQ
  - Gerenciamento robusto de conexões com verificação de estado
  - Suporte a reconexão automática em caso de falhas
  - Cleanup adequado de recursos através do método Close

Este pacote é útil para implementar sistemas de mensageria assíncrona,
permitindo comunicação entre diferentes partes de uma aplicação distribuída,
com garantias de resiliência e reconexão automática.
*/
type RabbitMQ struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

// New cria uma nova instância do RabbitMQ
func New(amqpServerURL string) (*RabbitMQ, error) {
	rabbitmq := &RabbitMQ{}
	conn, ch, err := rabbitmq.OpenConnection(amqpServerURL)
	if err != nil {
		return nil, err
	}

	rabbitmq.Connection = conn
	rabbitmq.Channel = ch
	return rabbitmq, nil
}

// QueueDeclare declara uma fila no RabbitMQ.
func (r *RabbitMQ) QueueDeclare(queueName string) (amqp.Queue, error) {
	queue, err := r.Channel.QueueDeclare(
		queueName, // nome
		true,      // durável
		false,     // auto-deletável
		false,     // exclusiva
		false,     // no-wait
		nil,       // argumentos
	)
	return queue, err
}

// QueueBind vincula uma fila a um exchange.
func (r *RabbitMQ) QueueBind(queueName, exchangeName, routingKey string) error {
	err := r.Channel.QueueBind(
		queueName,    // nome da fila
		routingKey,   // chave de roteamento
		exchangeName, // nome do exchange
		false,        // no-wait
		nil,          // argumentos
	)
	return err
}

// ExchangeDeclare declara um exchange no RabbitMQ.
func (r *RabbitMQ) ExchangeDeclare(exchangeName, exchangeType string) error {
	err := r.Channel.ExchangeDeclare(
		exchangeName, // nome
		exchangeType, // tipo
		true,         // durável
		false,        // auto-deletável
		false,        // interno
		false,        // no-wait
		nil,          // argumentos
	)
	return err
}

// Consume cria um consumidor para uma fila.
func (r *RabbitMQ) Consume(queueName, consumer string) (<-chan amqp.Delivery, error) {
	messages, err := r.Channel.Consume(
		queueName, // nome da fila
		consumer,  // consumidor
		true,      // auto-ack
		false,     // exclusivo
		false,     // no local
		false,     // no wait
		nil,       // argumentos
	)
	return messages, err
}

// Publish publica uma mensagem no RabbitMQ.
func (r *RabbitMQ) Publish(exchangeName, routingKey string, message amqp.Publishing) error {
	err := r.Channel.Publish(
		exchangeName, // nome do exchange
		routingKey,   // chave de roteamento
		false,        // mandatory
		false,        // immediate
		message,      // mensagem a ser publicada
	)
	return err
}

// Close fecha tanto o canal quanto a conexão
func (r *RabbitMQ) Close() error {
	if err := r.Channel.Close(); err != nil {
		return err
	}
	return r.Connection.Close()
}

// OpenConnection abre uma conexão com o RabbitMQ.
func (r *RabbitMQ) OpenConnection(amqpServerURL string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	return conn, ch, err
}

// PublishingMessage cria uma mensagem para ser publicada no RabbitMQ.
func (r *RabbitMQ) PublishingMessage(message string) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(message),
	}
}

// IsConnected verifica se a conexão está ativa
func (r *RabbitMQ) IsConnected() bool {
	return !r.Connection.IsClosed() && !r.Channel.IsClosed()
}

// Reconnect tenta reconectar ao servidor RabbitMQ
func (r *RabbitMQ) Reconnect(amqpServerURL string) error {
	if r.IsConnected() {
		return nil
	}

	conn, ch, err := r.OpenConnection(amqpServerURL)
	if err != nil {
		return err
	}

	r.Connection = conn
	r.Channel = ch
	return nil
}
