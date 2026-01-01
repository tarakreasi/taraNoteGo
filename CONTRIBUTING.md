# Contributing to TaraNote Go

## 1. Development Setup

### Prerequisites
- Go 1.23+
- Node.js 18+
- GCC (for SQLite CGO)

### Quick Start
```bash
# 1. Install Dependencies
go mod download
npm install

# 2. Setup Environment
cp .env.example .env

# 3. Migrate Database
make migrate

# 4. Build Frontend
npm run build

# 5. Run Server
make run
```

## 2. Directory Structure
- **Handlers**: `internal/handlers` (Controllers)
- **Models**: `internal/models` (Database Structs)
- **Routes**: `cmd/server/main.go`
- **Frontend**: `resources/`

## 3. Workflow
- Always format code with `go fmt ./...` before committing.
- Ensure `make build` passes.
- For frontend changes, run `npm run build` to update `manifest.json`.

## 4. Documentation
- Update `docs/API_REFERENCE.md` if new routes are added.
- Add migration notes to `CHANGELOG.md`.

## 5. Docker
To test the production build locally:
```bash
docker build -t taranote .
docker run -p 3000:3000 taranote
```
