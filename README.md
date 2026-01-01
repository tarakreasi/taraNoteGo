# TaraNote Go

**Current Version**: v1.0.0  
**Project Status**: Migration Complete

---

## 1. Project Overview

**TaraNote Go** is a high-performance port of the original Laravel-based `taraNote` application.

- **Goal**: Provide a lightweight, fast, and binary-distributable note-taking application.
- **State**: The backend has been fully rewritten in Go. The frontend (Vue 3 + Inertia) has been preserved and integrated successfully.

---

## 2. Technical Architecture

| Component | Technology |
| :--- | :--- |
| **Language** | Go 1.23+ |
| **Framework** | [Fiber v2](https://gofiber.io/) (Express-style) |
| **Database** | SQLite (via GORM) |
| **Frontend** | Vue 3 + Inertia.js (Monolithic) |
| **Assets** | Vite v6 → `public/build` |
| **Routing** | Defined in `cmd/server/main.go` |

> **Note**: Production builds use `go-sqlite3` which requires CGO.

The Go server renders a root HTML template (`views/app.html`) and passes initial state via JSON. No strict API separation; it behaves like a monolithic MVC app.

---

## 3. Completed Sprints (Migration Phase)

| Sprint | Description |
| :--- | :--- |
| **Sprint 1** | Setup & Foundation (Fiber, GORM) |
| **Sprint 2** | Frontend integration (Asset copying, Vite config) |
| **Sprint 3** | Authentication (Sessions, Bcrypt, Middleware) |
| **Sprint 4** | Notebooks Domain (CRUD) |
| **Sprint 5** | Notes Domain (CRUD, Image Uploads) |
| **Sprint 6** | Public Views (Home, Article, Browser) |
| **Sprint 7** | Settings Domain (Key-Value Store) |
| **Sprint 8** | Deployment Config (Docker, Makefile) |

---

## 4. Coding Standards

When writing new code for this project, adhere to these rules:

- **Handlers**: Place in `internal/handlers/`. Group by domain (e.g., `note.go`, `auth.go`).
- **Models**: Place in `internal/models/`. Use GORM tags.
- **Routes**: Register in `main.go`. Group protected routes under `api` or middleware blocks.
- **Error Handling**: Return JSON for API/Inertia errors (`c.Status(500).JSON(...)`).
- **Inertia**: Use `utils.CreateInertiaPage` to generate the correct JSON structure for the frontend.

---

## 5. Development Workflow

### Prerequisites
- **Go**: v1.23 or higher
- **Node.js**: v18+ (for building frontend assets)
- **GCC**: Required for CGO (SQLite driver)

### Commands

| Command | Description |
| :--- | :--- |
| `make run` | Starts server on port 3000 |
| `make build` | Outputs binaries to `bin/` |
| `make migrate` | Updates DB schema |
| `npm run build` | Generates `manifest.json` for asset loading |

### Quick Start

```bash
# 1. Clone & Setup
git clone https://github.com/tarakreasi/taraNoteGo.git
cd taraNoteGo
cp .env.example .env

# 2. Install Dependencies
go mod download
npm install

# 3. Migrate Database
make migrate

# 4. Build Frontend
npm run build

# 5. Run Server
make run
```

---

## 6. Known Issues / Roadmap

- **Ziggy Routes**: The frontend originally used Laravel's `route()` helper (Ziggy). This is currently mocked/hardcoded. A future task is to implement a robust JS route generator for Go.
- **Search**: The current search implementation in `PublicList` is a basic SQL `LIKE`. Consider full-text search (FTS5) for SQLite later.
- **Tests**: Zero unit tests currently exist. Prioritize adding `_test.go` for handlers.

---

## 7. Directory Structure Reference

```
taraNote_go/
├── cmd/
│   ├── server/       # main.go entry point
│   └── migrate/      # Database schema update utility
├── internal/
│   ├── config/       # Session, Environment configuration
│   ├── database/     # SQLite connection logic
│   ├── handlers/     # HTTP Controllers (Auth, Notes, etc.)
│   ├── models/       # GORM Data Models (structs)
│   └── utils/        # Inertia helper functions
├── resources/        # Vue/JS frontend source
├── public/           # Static assets & compiled build
└── views/            # HTML templates (Inertia root)
```

---

## 8. Docker Deployment

A multi-stage `Dockerfile` is included for lightweight deployment (Alpine Linux).

```bash
# Build Image
docker build -t taranote-go .

# Run Container
docker run -p 3000:3000 -v $(pwd)/database:/app/database taranote-go
```

---

## 9. License

This project is open-source and available under the [MIT License](LICENSE).

---

**Use `docs/sprint/prompt_ref.md` to prime your context before starting new tasks.**
