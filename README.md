# Chat App Go

Uma aplicação de chat simples construída em Go utilizando Clean Architecture, JWT para autenticação e PostgreSQL como banco de dados.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Arquitetura](#arquitetura)
- [Tecnologias](#tecnologias)
- [Funcionalidades](#funcionalidades)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Executando a Aplicação](#executando-a-aplicação)
- [API Endpoints](#api-endpoints)
- [Migrações do Banco](#migrações-do-banco)
- [Estrutura do Banco de Dados](#estrutura-do-banco-de-dados)

## 🎯 Visão Geral

Este projeto é uma API REST para um sistema de chat que permite:

- Criação e autenticação de usuários
- Login com JWT tokens
- Criação de chats entre usuários
- Sistema de mensagens (estrutura definida, implementação em andamento)
- WebSocket para comunicação em tempo real (código comentado, pronto para implementação)

O projeto segue os princípios de Clean Architecture, separando claramente as responsabilidades em camadas.

## 🏗️ Arquitetura

O projeto utiliza **Clean Architecture** com as seguintes camadas:

```
┌─────────────────────────────────────────┐
│              Presentation               │
│        (Controllers & Routes)           │
├─────────────────────────────────────────┤
│            Application                  │
│        (Use Cases & Repositories)       │
├─────────────────────────────────────────┤
│              Domain                     │
│           (Entities & Rules)            │
├─────────────────────────────────────────┤
│            Infrastructure               │
│      (Database, Security, HTTP)         │
└─────────────────────────────────────────┘
```

### Camadas:

- **Domain**: Entidades de negócio (User, Chat, Message, Claims)
- **Application**: Casos de uso e interfaces de repositórios
- **Infrastructure**: Implementações concretas (PostgreSQL, REST, Security)
- **Presentation**: Controllers e rotas HTTP

## 🛠️ Tecnologias

- **Go 1.20**
- **Gin Gonic** - Framework web HTTP
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT-Go** - Autenticação com JSON Web Tokens
- **bcrypt** - Hash de senhas
- **UUID** - Identificadores únicos
- **Logrus** - Logging estruturado
- **golang-migrate** - Migrações de banco de dados
- **godotenv** - Gerenciamento de variáveis de ambiente

## ✨ Funcionalidades

### ✅ Implementadas
- [x] **Sistema de configuração por profiles** (local, development, production)
- [x] **Scripts de execução** para diferentes ambientes
- [x] Criação de usuários com validação
- [x] Hash seguro de senhas com bcrypt
- [x] Sistema de login com JWT
- [x] Middleware de autenticação
- [x] Rotas protegidas
- [x] Conexão com PostgreSQL via GORM
- [x] Sistema de migrações
- [x] **CRUD completo de usuários**
- [x] **Estrutura completa para criação de chats**
- [x] **Repositórios PostgreSQL** para usuários e chats
- [x] **Factory pattern** para injeção de dependências
- [x] **Docker Compose** com PostgreSQL e PgAdmin

### 🚧 Em Desenvolvimento
- [ ] WebSocket para mensagens em tempo real
- [ ] Sistema de mensagens
- [ ] Autenticação com Google OAuth
- [ ] Busca avançada de usuários
- [ ] Histórico de mensagens
- [ ] Notificações push

## 📁 Estrutura do Projeto

```
chat-app-go/
├── cmd/
│   └── main.go                    # Ponto de entrada da aplicação
├── domain/
│   └── entities/
│       ├── user.go               # Entidade User
│       ├── chat.go               # Entidade Chat
│       ├── message.go            # Entidade Message
│       └── claims.go             # Claims JWT
├── application/
│   ├── repositories/
│   │   ├── user-repository.go    # Interface UserRepository
│   │   └── chat-repository.go    # Interface ChatRepository
│   └── usecases/
│       ├── auth/
│       │   └── login_use_case.go # Caso de uso de login
│       ├── user/
│       │   ├── create_user_use_case.go
│       │   ├── find_users_by_ids_use_case.go
│       │   └── finding_user_by_email_use_case.go
│       └── chat/
│           └── create_chat_use_case.go
├── infra/
│   ├── config/
│   │   └── config.go             # Sistema de configuração por profiles
│   ├── postgreSQL/
│   │   ├── config.go             # Configuração do banco
│   │   ├── models/
│   │   │   ├── user-model.go     # Modelo do usuário
│   │   │   └── chat-model.go     # Modelo do chat
│   │   ├── repositories/
│   │   │   ├── user-postgre-repository.go
│   │   │   └── chat-postgre-repository.go
│   │   └── migrations/
│   │       ├── 000001_create_users_table.up.sql
│   │       └── 000001_create_users_table.down.sql
│   ├── rest/
│   │   ├── controllers/
│   │   │   ├── auth_controller.go
│   │   │   ├── user_controller.go
│   │   │   └── chat_controller.go
│   │   ├── routes/
│   │   │   ├── routes.go
│   │   │   ├── auth_routes.go
│   │   │   ├── user_routes.go
│   │   │   └── chat_routes.go
│   │   ├── factories/
│   │   │   ├── create-user-use-case-factory.go
│   │   │   └── create-chat-use-case-factory.go
│   │   └── middlewares/
│   │       └── validate_request_middleware.go
│   └── security/
│       ├── middlewares/
│       │   └── authorization_middleware.go
│       └── services/
│           ├── hash_service.go
│           └── token_service.go
├── docker-compose.yaml          # Configuração Docker
├── Dockerfile                   # Build da aplicação Go
├── .env.example                 # Template de configuração
├── .env.local                   # Configuração local (gitignored)
├── .env.development             # Configuração desenvolvimento (gitignored)
├── .env.production              # Configuração produção (gitignored)
├── init.sql                     # Script de inicialização do banco
├── run.sh                       # Script de execução Unix
├── run.bat                      # Script de execução Windows
├── create-migration.sh          # Script para criar migrações
├── migrate.sh                   # Script para executar migrações
├── go.mod
└── go.sum
```

## 🚀 Instalação

### Opção 1: Usando Docker (Recomendado)

#### Pré-requisitos
- Docker
- Docker Compose

#### 1. Clone o repositório

```bash
git clone https://github.com/karlgama/chat-app-go.git
cd chat-app-go
```

#### 2. Configure as variáveis de ambiente

```bash
cp .env.example .env
# Edite o arquivo .env conforme necessário
```

#### 3. Inicie os serviços com Docker

```bash
# Apenas PostgreSQL e PgAdmin
docker-compose up -d

# Para incluir a aplicação Go no Docker (descomente a seção no docker-compose.yaml)
# docker-compose up -d --build
```

#### 4. Execute as migrações

```bash
# Se estiver rodando a app localmente
./migrate.sh

# Ou se preferir usar docker para as migrações
docker-compose exec postgres psql -U chat_user -d chat_app_db -f /docker-entrypoint-initdb.d/init.sql
```

### Opção 2: Instalação Local

#### Pré-requisitos
- Go 1.20+
- PostgreSQL 12+
- golang-migrate CLI (para migrações)

#### 1. Clone o repositório

```bash
git clone https://github.com/karlgama/chat-app-go.git
cd chat-app-go
```

#### 2. Instale as dependências

```bash
go mod download
```

#### 3. Instale o golang-migrate (opcional, para migrações)

```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Windows
# Baixe o binário do GitHub releases
```

## ⚙️ Configuração

O projeto utiliza um **sistema de profiles** para diferentes ambientes, permitindo configurações específicas para cada contexto.

### Profiles Disponíveis

- **`local`**: Desenvolvimento local (padrão)
- **`development`**: Ambiente de desenvolvimento  
- **`production`**: Ambiente de produção

### 1. Arquivos de Configuração

O sistema carrega automaticamente o arquivo `.env` baseado no profile:

```bash
.env.local          # Profile local
.env.development    # Profile development  
.env.production     # Profile production
.env.example        # Template de configuração
```

### 2. Configure as variáveis de ambiente

Copie o arquivo exemplo e ajuste conforme necessário:

```bash
cp .env.example .env.local
```

**Exemplo de configuração (.env.local):**
```env
# Environment Profile
APP_ENV=local

# Banco de dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=chat_user
DB_PASSWORD=chat_password
DB_NAME=chat_app_db

# JWT Security  
JWT_SECRET=local_development_secret_key_change_in_production

# Application
APP_PORT=8080

# Logs
LOG_LEVEL=debug
```

### 3. Definindo o Profile

Existem três formas de definir o profile:

#### Opção 1: Variável de ambiente
```bash
export APP_ENV=development  # Linux/macOS
set APP_ENV=development     # Windows
go run cmd/main.go
```

#### Opção 2: Argumento da linha de comando
```bash
go run cmd/main.go --env production
# ou
go run cmd/main.go --profile development
```

#### Opção 3: Scripts de execução
```bash
# Linux/macOS
./run.sh local
./run.sh development
./run.sh production

# Windows
.\run.bat local
.\run.bat development  
.\run.bat production
```

### 4. Banco de dados

#### Com Docker:
O PostgreSQL será criado automaticamente com o `docker-compose up`

#### Sem Docker:
```sql
-- Para profile local
CREATE DATABASE chat_app_db;

-- Para profile development
CREATE DATABASE chat_app_dev;

-- Para profile production
CREATE DATABASE chat_app_production;
```

## 🏃‍♂️ Executando a Aplicação

### Método 1: Scripts de Execução (Recomendado)

#### Windows:
```bash
# Profile local (padrão)
.\run.bat

# Profiles específicos
.\run.bat local
.\run.bat development  
.\run.bat production
```

#### Linux/macOS:
```bash
# Profile local (padrão)
./run.sh

# Profiles específicos  
./run.sh local
./run.sh development
./run.sh production
```

### Método 2: Execução Direta

```bash
# Com argumentos
go run cmd/main.go --env local
go run cmd/main.go --profile development

# Com variável de ambiente
export APP_ENV=production  # Linux/macOS
set APP_ENV=production     # Windows
go run cmd/main.go
```

### Com Docker

#### 1. Inicie os serviços

```bash
# Inicia PostgreSQL e PgAdmin
docker-compose up -d

# Verifica se os containers estão rodando
docker-compose ps
```

#### 2. Execute as migrações

```bash
# Com a aplicação rodando localmente
./migrate.sh

# Ou execute as migrações manualmente no container
docker-compose exec postgres psql -U chat_user -d chat_app_db
```

#### 3. Inicie a aplicação Go

```bash
# Localmente (recomendado para desenvolvimento)
.\run.bat local

# Ou descomente a seção chat-app no docker-compose.yaml e execute:
# docker-compose up -d --build
```

#### 4. Acesse os serviços

- **API**: http://localhost:8080
- **PgAdmin**: http://localhost:5050
  - Email: admin@chat-app.com
  - Senha: admin123

A aplicação estará disponível em `http://localhost:8080`

## 📡 API Endpoints

### Autenticação

#### POST `/login`
Realiza login e retorna JWT token

```json
{
  "email": "user@example.com",
  "password": "senha123"
}
```

**Resposta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Usuários

#### POST `/users`
Cria um novo usuário

```json
{
  "name": "João Silva",
  "email": "joao@example.com",
  "password": "senha123"
}
```

**Resposta:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "João Silva",
  "email": "joao@example.com",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Rotas Protegidas

Todas as rotas abaixo requerem autenticação via JWT token.

**Headers obrigatórios:**
```
Authorization: Bearer <seu_jwt_token>
```

#### GET `/protected`
Rota de exemplo que requer autenticação

**Resposta:**
```json
{
  "message": "protected route"
}
```

#### POST `/chats`
Cria um novo chat entre usuários

**Payload:**
```json
{
  "user_ids": [
    "550e8400-e29b-41d4-a716-446655440000",
    "550e8400-e29b-41d4-a716-446655440001"
  ]
}
```

**Resposta:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Chats (Em Desenvolvimento)

Mais endpoints de chat serão adicionados:
- GET `/chats` - Listar chats do usuário
- GET `/chats/{id}` - Obter detalhes de um chat
- POST `/chats/{id}/messages` - Enviar mensagem
- GET `/chats/{id}/messages` - Listar mensagens do chat

## 🗄️ Migrações do Banco

### Sistema de Migrações com GORM

O projeto utiliza o sistema de **Auto Migration** do GORM, que:
- ✅ **Cria tabelas** automaticamente
- ✅ **Adiciona colunas** faltantes  
- ✅ **Cria índices** necessários
- ✅ **Mantém dados** existentes (não remove colunas)
- ✅ **Executa migrações customizadas** para casos específicos

### Comandos de Migração

#### Método 1: Via Aplicação Principal (Recomendado)

```bash
# Visualizar ajuda
go run cmd/main.go --help

# Executar todas as migrações pendentes
go run cmd/main.go migrate

# Remover todas as tabelas do banco
go run cmd/main.go drop

# Resetar banco (drop + migrate)
go run cmd/main.go reset
```

#### Método 2: Scripts

```bash
# Windows
.\migrate.bat

# Linux/macOS  
./migrate.sh
```

### Auto-Migração na Inicialização

Em ambientes de **desenvolvimento** (`local`, `development`), as migrações são executadas automaticamente quando a aplicação inicia:

```go
// Em cmd/main.go
if config.IsDevelopment() || config.IsLocal() {
    log.Println("Executando auto-migração das tabelas...")
    postgreSQL.AutoMigrate()
}
```

### Migrações Customizadas

Para casos que o GORM não cobre automaticamente, use migrações customizadas em `infra/postgreSQL/migrations.go`:

```go
func RunCustomMigrations() {
    // Índices específicos
    createIndexes()
    
    // Foreign keys
    createForeignKeys()
    
    // Dados iniciais
    seedData()
}
```

## 📊 Estrutura do Banco de Dados

### Tabela `users`

```sql
CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
    externalID UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP,
    updatedAt TIMESTAMP
);
```

**Campos:**
- `ID`: Chave primária auto-incremento
- `externalID`: UUID público do usuário
- `name`: Nome do usuário
- `email`: Email único do usuário
- `password`: Senha hasheada com bcrypt
- `createdAt`: Data de criação
- `updatedAt`: Data de última atualização

## 🔒 Segurança

- **Senhas**: Hasheadas com bcrypt (cost 14)
- **JWT**: Tokens com expiração de 5 minutos
- **CORS**: Configurado para desenvolvimento
- **Validação**: Validação de entrada em todos os endpoints

## 🚧 Próximos Passos

1. **WebSocket**: Implementar comunicação em tempo real
2. **Mensagens**: Sistema completo de envio/recebimento
3. **Busca**: Implementar busca real de usuários
4. **OAuth**: Integração com Google
5. **Testes**: Adicionar testes unitários e de integração
6. **Docker**: Containerização da aplicação
7. **CI/CD**: Pipeline de deploy automatizado

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👨‍💻 Autor

**Karl Gama**
- GitHub: [@karlgama](https://github.com/karlgama)

---

**Nota**: Este projeto está em desenvolvimento ativo. Algumas funcionalidades podem estar incompletas ou em fase de implementação.
