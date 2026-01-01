# AI Context & Prompt Reference: TaraNote Go

**Current Date**: January 1, 2026
**Project Status**: Migration Complete (v1.0.0)

## 1. Project Overview
**TaraNote Go** is a high-performance port of the original Laravel-based `taraNote` application.
- **Goal**: Provide a lightweight, fast, and binary-distributable note-taking application.
- **State**: The backend has been fully rewritten in Go. The frontend (Vue 3 + Inertia) has been preserved and integrated successfully.

## 2. Technical Architecture
- **Language**: Go 1.23+
- **Framework**: [Fiber v2](https://gofiber.io/) (Express-style)
- **Database**: SQLite (via GORM)
    - *Note*: Production builds use `go-sqlite3` (CGO required).
- **Frontend Integration**: Inertia.js (Monolithic).
    - The Go server renders a root HTML template (`views/app.html`) and passes initial state via JSON.
    - No strict API separation; it behaves like a monolithic MVC app.
- **Routing**: Defined in `cmd/server/main.go`.
- **Assets**: Vite v6 builds to `public/build`.

## 3. Completed Sprints (Migration Phase)
The following sprints have been completed effectively:
- **Sprint 1**: Setup & Foundation (Fiber, GORM).
- **Sprint 2**: Frontend integration (Asset copying, Vite config adjustment).
- **Sprint 3**: Authentication (Sessions, Bcrypt, Middleware).
- **Sprint 4**: Notebooks Domain (CRUD).
- **Sprint 5**: Notes Domain (CRUD, Image Uploads).
- **Sprint 6**: Public Views (Home, Article, Browser).
- **Sprint 7**: Settings Domain (Key-Value Store).
- **Sprint 8**: Deployment Config (Docker, Makefile).

## 4. Coding Standards
When writing new code for this project, adhere to these rules:
- **Handlers**: Place in `internal/handlers/`. Group by domain (e.g., `note.go`, `auth.go`).
- **Models**: Place in `internal/models/`. Use GORM tags.
- **Routes**: Register in `main.go`. Group protected routes under `api` or middleware blocks.
- **Error Handling**: Return JSON for API/Inertia errors (`c.Status(500).JSON(...)`).
- **Inertia**: Use `utils.CreateInertiaPage` to generate the correct JSON structure for the frontend.

## 5. Development Workflow
- **Run Dev**: `make run` (Starts server on port 3000).
- **Build**: `make build` (Outputs to `bin/`).
- **Migrate**: `make migrate` (Updates DB schema).
- **Frontend**: `npm run build` (Required to generate `manifest.json` for asset loading).

## 6. Known Issues / Roadmap
- **Ziggy Routes**: The frontend originally used Laravel's `route()` helper (Ziggy). This is currently mocked/hardcoded. A future task is to implement a robust JS route generator for Go.
- **Search**: The current search implementation in `PublicList` is a basic SQL `LIKE`. Consider full-text search (FTS5) for SQLite later.
- **Tests**: Zero unit tests currently exist. Prioritize adding `_test.go` for handlers.

## 7. Directory Structure Reference
```
taraNote_go/
├── cmd/
│   ├── server/   (main.go entry)
│   └── migrate/  (schema update)
├── internal/
│   ├── config/   (session, env)
│   ├── database/ (sqlite connect)
│   ├── handlers/ (controllers)
│   ├── models/   (structs)
│   └── utils/    (inertia helper)
├── resources/    (vue/js sources)
└── views/        (html templates)
```

**Use this file to prime your context before starting new tasks.**
