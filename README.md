# NexTask

A production-grade project management API built with Go and PostgreSQL. Built as a learning project to master idiomatic Go and PostgreSQL by building a real application — from zero to deployed.

## Features

- **User Authentication** — JWT access tokens + refresh tokens
- **Workspaces → Projects → Tasks** hierarchy
- **Real-time Updates** — Server-Sent Events (SSE) with goroutines
- **File Attachments** — local storage → S3-compatible
- **Role-Based Access Control** (RBAC)
- **Activity Feed & Audit Log**
- **Full Observability** — structured logging, metrics

## Tech Stack

| Layer | Choice | Why |
|---|---|---|
| Router | `chi` | Lightweight, stdlib-compatible, idiomatic |
| DB Driver | `pgx/v5` | Gold-standard Postgres driver |
| SQL | `sqlc` | Type-safe queries, no ORM magic |
| Migrations | `golang-migrate` | Battle-tested, CLI + programmatic |
| Config | `godotenv` + `viper` | Simple → powerful progression |
| Validation | `go-playground/validator` | Struct tags, like Joi but typed |
| Auth | `golang-jwt/jwt` | Minimal, explicit |
| Logging | `slog` (stdlib) | Standard, structured, modern |
| Testing | `testify` + `mockery` | Industry standard |
| Container | Docker multi-stage | Tiny final image (~15MB) |

## Project Structure

```
nextask/
├── cmd/api/                  # Entry point — wires deps, starts server
├── internal/
│   ├── domain/              # Pure types: Task, User, Project
│   ├── handler/            # HTTP handlers (thin — delegate to service)
│   ├── service/            # Business logic
│   ├── repository/         # DB queries (sqlc generated)
│   ├── middleware/         # Auth, logging, recovery
│   └── config/             # App configuration
├── db/
│   ├── migrations/         # .sql migration files
│   └── queries/           # .sql query files for sqlc
├── configs/                # Environment-specific configs
├── sqlc.yaml               # sqlc code generation config
├── .golangci.yml          # Linter configuration
├── .air.toml               # Hot-reload configuration
└── .env                    # Environment variables
```

## Quick Start

### Prerequisites

- Go 1.22+
- PostgreSQL 15+
- Docker (optional, for running Postgres)

### Local Development

```bash
# 1. Clone and set up
git clone https://github.com/thepratiksah/nextask.git
cd nextask
cp .env.example .env

# 2. Start PostgreSQL (Docker)
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=nextask postgres:15

# 3. Run migrations
migrate -path db/migrations -database "$DATABASE_URL" up

# 4. Generate sqlc code
sqlc generate

# 5. Run the server
go run ./cmd/api

# or with hot-reload (install air first)
air
```

## Quick Reference

```bash
# Run the app
go run ./cmd/api

# Run all tests
go test ./... -v

# Generate sqlc code
sqlc generate

# Run migrations
migrate -path db/migrations -database "$DATABASE_URL" up

# Lint
golangci-lint run

# Build production binary
go build -o bin/nextask ./cmd/api

# Docker
docker build -t nextask .
docker run -p 8080:8080 nextask
```

## Phases

1. **Go Foundations** — Domain models, structs, pointers, error handling
2. **Dev Tools & Workspace** — Project structure, linting, hot-reload
3. **Concurrency** — Goroutines, channels, SSE real-time events
4. **PostgreSQL** — Schema design, sqlc, migrations, connection pooling
5. **REST API** — chi handlers, middleware, validation, auth
6. **Architecture & Testing** — Clean architecture, DI, table-driven tests
7. **Deployment** — Docker, CI/CD, Cloud Run

## Learning

This project is built as part of the **Go + PostgreSQL Mastery** program. Every feature is designed to teach idiomatic Go patterns by comparing them to the Node.js equivalents you already know.

Coming from Node.js? Here's the mental model:

| Node.js | Go |
|---|---|
| `package.json` | `go.mod` |
| `express()` | `chi.NewRouter()` |
| `async/await` | goroutines + channels |
| `try/catch` | `if err != nil` |
| `mongoose.Schema` | sqlc generated structs |
| `jest` | `testing` (stdlib) |

## Author

Built by **[Pratik Sah](https://github.com/thepratiksah)** — experienced Node.js developer learning Go the right way: by building something real.
