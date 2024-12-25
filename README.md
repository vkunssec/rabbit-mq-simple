# RabbitMQ Simple Service

Este é um projeto de exemplo que demonstra a implementação de um sistema de mensageria usando RabbitMQ com Go (Golang). O projeto consiste em um serviço sender (produtor) e dois consumers (consumidores) que utilizam routing keys diferentes.

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação
- [RabbitMQ](https://www.rabbitmq.com/) - Message Broker
- [Docker](https://www.docker.com/) - Containerização
- [Fiber](https://gofiber.io/) - Framework Web
- [Swagger](https://swagger.io/) - Documentação da API
- [Air](https://github.com/cosmtrek/air) - Live Reload para desenvolvimento
- [Husky](https://typicode.github.io/husky/) - Git Hooks

## 📋 Pré-requisitos

- Go 1.16+
- Docker e Docker Compose
- Make (opcional, mas recomendado)

## 🔧 Configuração do Ambiente

1. Clone o repositório:

```
git clone https://github.com/vkunssec/rabbit-mq-simple.git
cd rabbit-mq-simple
```

2. Configure as variáveis de ambiente:
```bash
make env
```

3. Instale as dependências:
```bash
go mod download
```

## 🚀 Executando o Projeto

### Usando Docker Compose

```bash
docker-compose up -d
```

### Usando Make

Para executar todos os serviços:
```bash
make run
```

Para executar serviços individualmente:
```bash
# Executar o sender
make sender

# Executar o consumer 1
make consumer-1

# Executar o consumer 2
make consumer-2

# Executar ambos os consumers
make consumers
```

## 📦 Estrutura do Projeto

```
.
├── cmd/
│   ├── consumer-1/    # Serviço consumidor 1
│   ├── consumer-2/    # Serviço consumidor 2
│   └── sender/        # Serviço produtor
├── pkg/
│   ├── domain/        # Domínio da aplicação
│   │   └── rabbitmq/  # Implementação base do RabbitMQ
│   └── repository/    # Camada de repositório
├── internal/          # Código interno da aplicação
├── docker-compose.yml
└── Makefile
```

## 🔍 Implementação do RabbitMQ

### Domain (pkg/domain/rabbitmq)
O domínio implementa a estrutura base do RabbitMQ com as seguintes funcionalidades:

- Gerenciamento de conexões e canais
- Verificação de estado da conexão
- Reconexão automática
- Declaração de exchanges e filas
- Publicação e consumo de mensagens
- Cleanup adequado de recursos

### Repository (pkg/repository)
O repositório implementa a lógica de negócio específica:

- Configuração de exchanges e filas
- Roteamento de mensagens usando routing keys
- Gerenciamento de múltiplos consumers
- Logging de mensagens

### Routing Keys
O sistema utiliza as seguintes routing keys:
- `route.service1`: Para mensagens destinadas ao Consumer 1
- `route.service2`: Para mensagens destinadas ao Consumer 2

## 🐰 Configuração RabbitMQ

O RabbitMQ está configurado com as seguintes definições:

- Exchange: `ExchangeService1` (tipo: direct)
- Filas: 
  - `QueueService1`: Vinculada à routing key `route.service1`
  - `QueueService2`: Vinculada à routing key `route.service2`
- URL: `amqp://guest:guest@rabbitmq:5672/%2f`
- Interface de gerenciamento: `http://localhost:15672`
  - Usuário: `guest`
  - Senha: `guest`

## 🛠️ Comandos Make Disponíveis

- `make run`: Compila e executa todos os serviços
- `make build`: Compila os serviços
- `make swagger`: Gera a documentação Swagger
- `make dev`: Inicia o ambiente de desenvolvimento com hot-reload
- `make sender`: Executa apenas o serviço sender
- `make consumer-1`: Executa o consumer 1
- `make consumer-2`: Executa o consumer 2
- `make consumers`: Executa ambos os consumers
- `make env`: Cria o arquivo de variáveis de ambiente

## 🔒 Git Hooks

O projeto utiliza Husky para gerenciar git hooks. Antes de cada commit, são executadas as seguintes verificações:

- `go mod tidy`
- `go fmt ./...`
- `go vet ./...`
- `golangci-lint run ./...`

## 🐳 Containers Docker

O projeto inclui quatro containers principais:

1. **sender**: Serviço produtor de mensagens
   - Porta: 3000
   - Dockerfile: `Dockerfile.sender`

2. **consumer-1**: Primeiro serviço consumidor
   - Dockerfile: `Dockerfile.consumer-1`

3. **consumer-2**: Segundo serviço consumidor
   - Dockerfile: `Dockerfile.consumer-2`

4. **rabbitmq**: Servidor RabbitMQ
   - Portas: 5672 (AMQP), 15672 (Management UI)
   - Imagem: rabbitmq:3.11-management
