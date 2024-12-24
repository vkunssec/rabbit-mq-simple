# RabbitMQ Simple Service

Este é um projeto de exemplo que demonstra a implementação de um sistema de mensageria usando RabbitMQ com Go (Golang). O projeto consiste em dois serviços principais: um sender (produtor) e um consumer (consumidor).

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
git clone https://github.com/seu-usuario/rabbit-mq-simple.git
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
# Executar apenas o sender
make sender

# Executar apenas o consumer
make consumer
```

Para desenvolvimento com hot-reload:
```bash
make dev
```

## 📦 Estrutura do Projeto

```
.
├── cmd/
│   ├── consumer/     # Serviço consumidor
│   └── sender/       # Serviço produtor
├── internal/         # Código interno da aplicação
├── pkg/              # Pacotes reutilizáveis
├── docker-compose.yml
├── Dockerfile.consumer
├── Dockerfile.sender
└── Makefile
```

## 🔍 Endpoints da API

A documentação completa da API está disponível através do Swagger UI:
```
http://localhost:3000/swagger/
```

## 🐰 Configuração RabbitMQ

O RabbitMQ está configurado com as seguintes definições padrão:

- URL: `amqp://guest:guest@rabbitmq:5672/%2f`
- Interface de gerenciamento: `http://localhost:15672`
- Usuário padrão: `guest`
- Senha padrão: `guest`

## 🛠️ Comandos Make Disponíveis

- `make run`: Compila e executa todos os serviços
- `make build`: Compila os serviços
- `make swagger`: Gera a documentação Swagger
- `make dev`: Inicia o ambiente de desenvolvimento com hot-reload
- `make sender`: Executa apenas o serviço sender
- `make consumer`: Executa apenas o serviço consumer
- `make env`: Cria o arquivo de variáveis de ambiente

## 🔒 Git Hooks

O projeto utiliza Husky para gerenciar git hooks. Antes de cada commit, são executadas as seguintes verificações:

- `go mod tidy`
- `go fmt ./...`
- `go vet ./...`
- `golangci-lint run ./...`

## 🐳 Containers Docker

O projeto inclui três containers principais:

1. **sender**: Serviço produtor de mensagens
   - Porta: 3000
   - Dockerfile: `Dockerfile.sender`

2. **consumer**: Serviço consumidor de mensagens
   - Dockerfile: `Dockerfile.consumer`

3. **rabbitmq**: Servidor RabbitMQ
   - Portas: 5672 (AMQP), 15672 (Management UI)
   - Imagem: rabbitmq:3.11-management
