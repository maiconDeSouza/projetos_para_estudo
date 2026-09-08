# Gamer Deals Hub API

🎮 **Gamer Deals Hub API** é uma API backend desenvolvida em Go que consome uma API externa de ofertas de jogos, aplica regras de negócio para calcular descontos e ofertas especiais, e persiste os dados em um banco de dados relacional (PostgreSQL). O projeto segue os princípios da **Clean Architecture** para garantir manutenibilidade, testabilidade e separação clara de responsabilidades.

---

## 🚀 Tecnologias

- **Go** (1.21+)
- **PostgreSQL** 15 (executado via Docker)
- **Docker** e **Docker Compose**
- **Clean Architecture** (camadas: handlers, services, repository, models, database)
- **Gorilla Mux** ou **net/http** (para roteamento – a escolha é sua)
- **SQL** (para queries) e **pgx** (driver PostgreSQL)

---

## 📁 Estrutura de Pastas

A estrutura segue o padrão recomendado para projetos Go:

```
gamer-deals-hub/
├── cmd/
│   └── api/                # Ponto de entrada da aplicação (main.go)
├── internal/               # Código privado da aplicação
│   ├── database/           # Configuração e conexão com o banco
│   ├── models/             # Definições das entidades (structs)
│   ├── repository/         # Camada de acesso a dados (CRUD)
│   ├── services/           # Camada de lógica de negócio e regras
│   └── handlers/           # Manipuladores HTTP (controllers)
├── docker-compose.yml      # Orquestração do PostgreSQL
├── go.mod                  # Módulo Go
└── README.md
```

---

## 🐳 Pré-requisitos

- [Go](https://go.dev/dl/) (versão 1.21 ou superior)
- [Docker](https://www.docker.com/get-started) e [Docker Compose](https://docs.docker.com/compose/install/)
- (Opcional) [Make](https://www.gnu.org/software/make/) para automação de tarefas

---

## 🛠️ Configuração e Execução

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/gamer-deals-hub.git
cd gamer-deals-hub
```

### 2. Inicialize o módulo Go (se ainda não estiver feito)

```bash
go mod init gamer-deals-hub
```

### 3. Configure o banco de dados com Docker Compose

Crie um arquivo `docker-compose.yml` na raiz com o seguinte conteúdo:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: gamer_deals_postgres
    environment:
      POSTGRES_USER: gamer_user
      POSTGRES_PASSWORD: gamer_pass
      POSTGRES_DB: gamer_deals
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  postgres_data:
```

Suba o container:

```bash
docker-compose up -d
```

### 4. Execute as migrações (se houver)

Você pode usar ferramentas como `migrate` ou executar scripts SQL manualmente. Por enquanto, o banco estará vazio.

### 5. Execute a aplicação

```bash
go run cmd/api/main.go
```

O servidor será iniciado na porta padrão (ex: `8080`).

---

## 🧱 Arquitetura Limpa (Clean Architecture)

O projeto está organizado em camadas concêntricas:

- **Handlers** (camada de apresentação): recebem as requisições HTTP, validam os dados de entrada e chamam os serviços.
- **Services** (camada de negócio): contêm a lógica de aplicação, como cálculo de descontos, regras de ofertas e orquestração de dados.
- **Repository** (camada de persistência): abstraem o acesso ao banco de dados, implementando operações CRUD para as entidades.
- **Models**: definem as estruturas de dados que transitam entre as camadas.
- **Database**: configuração da conexão com o PostgreSQL.

Essa separação permite que cada camada seja testada independentemente e facilita futuras alterações (ex: trocar o banco de dados ou o framework web).

---

## 🔧 Variáveis de Ambiente

Recomenda-se o uso de um arquivo `.env` para configurações sensíveis. Exemplo:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=gamer_user
DB_PASSWORD=gamer_pass
DB_NAME=gamer_deals
API_PORT=8080
EXTERNAL_DEALS_API_URL=https://api.exemplo.com/deals
```

Carregue essas variáveis no seu código com `os.Getenv` ou com pacotes como `godotenv`.

---

## 📦 Dependências

As principais dependências (gerenciadas via `go.mod`) podem incluir:

- `github.com/lib/pq` ou `github.com/jackc/pgx/v5` (driver PostgreSQL)
- `github.com/gorilla/mux` (roteador HTTP)
- `github.com/joho/godotenv` (carregamento de variáveis de ambiente)

Instale-as com:

```bash
go get github.com/gorilla/mux
go get github.com/jackc/pgx/v5
```

---

## 🧪 Testes

Para rodar os testes unitários (quando implementados):

```bash
go test ./...
```

---

## 🤝 Como Contribuir

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas alterações (`git commit -m 'Adiciona nova feature'`).
4. Push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 👨‍💻 Autor

[Seu Nome] – [seu@email.com](mailto:seu@email.com)

---

**Divirta-se codando e caçando as melhores ofertas!** 🎮🛒