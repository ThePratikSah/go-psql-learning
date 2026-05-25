# Agent Handoff Note — NexTask

> **Written:** 2026-05-25
> **Project:** `github.com/thepratiksah/nextask`
> **Language:** Go 1.25 + PostgreSQL 15+
> **Status:** Phase 2 (Dev Tools & Workspace) complete, ready for Phase 1 (Go Foundations)

---

## What Is This Project?

**NexTask** — a production-grade project management API (Trello/Linear clone) being built to teach Go + PostgreSQL to an experienced Node.js developer. Every feature is designed around the Node.js → Go mental bridge.

Built as part of the **Go + PostgreSQL Mastery** program (skill: `go-pg-mastery`). The learner's GitHub username is `thepratiksah`.

---

## What Was Accomplished

### 1. Project Initialized (2 commits)

| Commit | Message |
|--------|---------|
| `bca6c30` | `initial commit: enterprise project structure for NexTask` |
| `636d8f8` | `chore: remove deprecated chi middleware.RealIP` |

### 2. Directory Structure (Enterprise Layout)

```
go-psql-learning/
├── cmd/api/
│   └── main.go                  # Entry point — chi router + graceful shutdown
├── internal/
│   ├── config/
│   │   └── config.go           # Loads .env → typed Config struct
│   ├── domain/
│   │   ├── task.go            # Task struct, TaskStatus enum
│   │   ├── user.go            # User struct, UserRole enum
│   │   ├── project.go         # Project struct
│   │   └── workspace.go        # Workspace struct
│   ├── handler/
│   │   └── task_handler.go     # CRUD handlers (placeholder bodies)
│   ├── service/
│   │   └── task_service.go     # TaskService with TaskRepository interface
│   ├── repository/
│   │   └── task_repository.go  # Stub implementations (→ sqlc Phase 4)
│   ├── middleware/
│   │   └── middleware.go       # RequireAuth + LoggerMiddleware stubs
│   └── sse/
│       └── broadcaster.go      # SSE event broadcaster (sync.RWMutex + channels)
├── db/
│   ├── migrations/             # EMPTY — awaiting Phase 4
│   └── queries/                # EMPTY — awaiting Phase 4
├── .env + .env.example
├── .golangci.yml               # Linter config (errcheck, staticcheck, revive, etc.)
├── .air.toml                     # Hot-reload config (builds to ./tmp/main)
├── sqlc.yaml                     # sqlc config (uuid → uuid.UUID, timestamptz → time.Time)
├── go.mod + go.sum               # Module + dependencies
└── README.md                       # Project docs with quick start
```

### 3. Dependencies (go.mod)

| Package | Purpose |
|---------|---------|
| `github.com/go-chi/chi/v5` | HTTP router |
| `github.com/google/uuid` | UUID types |
| `github.com/joho/godotenv` | .env loading |
| `github.com/jackc/pgx/v5` | PostgreSQL driver (imported but not yet used — Phase 4) |
| `github.com/golang-jwt/jwt/v5` | JWT auth (imported but not yet used — Phase 5) |
| `github.com/go-playground/validator/v10` | Struct validation (imported but not yet used — Phase 5) |
| `github.com/stretchr/testify` | Testing (imported but not yet used — Phase 6) |

> **Note:** `go mod tidy` pruned pgx, jwt, validator, testify from `go.mod` because they aren't imported yet — only chi, uuid, godotenv remain as direct dependencies. When you implement phases that use these packages, they'll be added back automatically.

### 4. Configuration Files

| File | Purpose | Key Details |
|------|---------|-------------|
| `.env` | Environment variables | `DATABASE_URL`, `JWT_SECRET`, `JWT_EXPIRY=15m`, `PORT=8080`, `ENV=development` |
| `.golangci.yml` | Linter rules | Enforces errcheck, staticcheck, gofmt, goimports, revive |
| `.air.toml` | Hot reload | `go build -o ./tmp/main ./cmd/api`, watches `.go` and `.sql` |
| `sqlc.yaml` | SQL codegen | Maps `uuid` → `uuid.UUID`, `timestamptz` → `time.Time` |
| `.gitignore` | Git exclusions | `tmp/`, `bin/`, `.env`, `air.log`, IDE dirs |

### 5. What Works

- ✅ `go build ./...` — compiles cleanly, zero errors
- ✅ `go run ./cmd/api` — starts the HTTP server on :8080 with chi middleware
- ✅ Graceful shutdown on SIGINT/SIGTERM
- ✅ `/health` endpoint (chi heartbeat middleware)
- ✅ Root endpoint returns "NexTask API"

### 6. What Does Not Work Yet (By Design)

All handler/service/repository methods are **placeholders** returning stub values. This is intentional — they will be implemented phase by phase.

---

## Architecture Decision Notes

### Why This Layout?

The structure follows the [Go Project Layout](https://github.com/golang-standards/project-layout) convention:

- `cmd/api/` — only the entry point. No business logic. Wire deps, start server.
- `internal/` — Go-enforced private packages. Only importable within this module.
- `internal/domain/` — pure types, zero DB/HTTP imports. The "source of truth" for entities.
- `internal/handler/` — thin HTTP layer. Decode request, validate, delegate to service.
- `internal/service/` — business logic. Defines interfaces for dependencies.
- `internal/repository/` — DB queries. Implements interfaces defined by service.
- `internal/middleware/` — HTTP middleware (auth, logging).
- `internal/sse/` — real-time event system.

### Node.js → Go Mental Bridges to Reference

| Node.js | Go |
|---------|----|
| `package.json` | `go.mod` |
| `express()` | `chi.NewRouter()` |
| `async/await` | goroutines + channels |
| `try/catch` | `if err != nil` |
| `mongoose.Schema` | sqlc generated structs |
| `jest` | `testing` (stdlib) |
| `process.env` | `config.Load()` → explicit dependency injection |
| `app.use(middleware)` | `r.Use(func(http.Handler) http.Handler)` — decorator pattern |

### Key Patterns Enforced

1. **`context.Context` always first** — Every function touching the outside world takes `context.Context` as its first parameter.
2. **Error wrapping with `%w`** — Never `return err` without wrapping: `return fmt.Errorf("FindByID: %w", err)`
3. **Pointers for nullability** — `*uuid.UUID` = nullable field. Non-pointer `uuid.UUID` = required.
4. **`json:"-"` for sensitive fields** — e.g., `Password string \`json:"-"\``
5. **Interface in consumer package** — `TaskRepository` interface is defined in `service/`, not `repository/`. This is Go convention (implicit satisfaction).

---

## What's Next (Phase 1: Go Foundations)

### Priority 1: Deep-dive into Domain Models

The next session should be a teaching session. The learner understands the structure at a high level but needs a deep-dive into Go fundamentals through the lens of these domain models:

**Topics to cover:**
1. **Pointers vs values** — Why `AssigneeID *uuid.UUID` vs `ID uuid.UUID`? Walk through the `domain/task.go` struct field by field.
2. **Struct tags** — How `json:"id"`, `json:"assignee_id,omitempty"`, `json:"-"` work at runtime.
3. **Custom types** — `TaskStatus string` + const block. Compare to TypeScript discriminated unions.
4. **Error handling** — `if err != nil` pattern with error wrapping. No try/catch.
5. **Methods on structs** — `func (t *Task) IsComplete() bool` — method receivers vs JavaScript class methods.

### Priority 2: Add Behavior to Domain Types

Extend domain structs with methods:
- `func (t *Task) CanTransitionTo(status TaskStatus) bool` — business rules
- `func (u *User) HasPermission(action string) bool` — RBAC foundation
- Validation methods using `go-playground/validator`

### Priority 3: Wire Config to Main

Replace inline `os.Getenv("PORT")` in `main.go` with `cfg := config.Load()` and `cfg.Port`.

---

## Known Issues / Gotchas

1. **Unused imports:** `pgx/v5`, `jwt`, `validator`, `testify` are in the original `go get` call but `go mod tidy` pruned them because nothing imports them yet. They'll be re-added when used in later phases.
2. **`configs/` directory** — Was mentioned in README structure but not yet created with any files. Can be used for environment-specific configs (dev/prod).
3. **`go.mod` shows `go 1.25.0`** — Make sure the local Go version is 1.25+ or update the constraint.
4. **`middleware.RealIP` removed** — Deprecated in chi v5. If IP detection is needed, use `chi/middleware.StripSlashes` or cloud-specific detection.

---

## How to Run Everything

```bash
# Build and verify
go build ./...

# Run the server
go run ./cmd/api

# Lint
golangci-lint run

# Hot-reload (install air first: go install github.com/air-verse/air@latest)
air
```

---

## Teaching Notes

- The learner is an **experienced Node.js developer**. Always anchor new Go concepts to their Node.js equivalent.
- The project is built as part of the **Go + PostgreSQL Mastery** skill. Reference files at `~/.claude/skills/go-pg-mastery/references/` for detailed phase guides.
- The learner enjoys **deep, architectural walkthroughs** — don't skip the "why" when explaining patterns.
- Prefer **short, focused teaching sessions** over long monologues. Code snippets > prose.
- The learner has been working through this incrementally (created structure, then asked for a walkthrough, now wrapping up). Next session should resume with Phase 1 content.
