# TaraNote Go Migration - Reference Context

## Project Overview
Migrating **TaraNote** (Laravel 12 + Vue 3/Inertia) to **Go** (Fiber + GORM + Inertia Adapter).
The goal is to replace the backend while preserving the existing Vue frontend with minimal changes.

## Architecture
- **Language**: Go 1.23+
- **Framework**: Fiber v2
- **ORM**: GORM (SQLite for dev, Postgres capable)
- **Frontend**: Vue 3 (Options/Composition API) + Inertia.js
- **Template Engine**: `html/v2` (Go templates serving Inertia root)

## Current Status
**Sprint 1 Completed**: Foundation
- Go module initialized (`github.com/tarakreasi/taraNote_go`)
- Project structure set (`cmd`, `internal`, `resources`)
- Database connection established (SQLite)

**Sprint 2 Completed**: Frontend Pipeline
- Assets copied from `taraNote` (`resources/js`, `css`, `vite.config.js`)
- Dependencies installed (Vite downgraded to v6)
- Build verified (`npm run build` works)
- Ziggy (Laravel Routes) temporarily disabled in `app.js`

## Immediate Next Steps (Sprint 3)
**Goal**: Authentication & Authorization
1.  Implement Session Management (Fiber Session/Cookies).
2.  Create Auth Controller (Login, Logout).
3.  Implement Middleware to guard protected routes.
4.  Re-enable/Mock necessary frontend auth props (`auth.user`).

## Known Issues / Tech Debt
- **Ziggy**: `resources/js/app.js` has `ZiggyVue` commented out. Needs a Go-compatible solution for named routes.
- **Vite Dev Server**: `views/app.html` hardcodes localhost:5173. Needs logic to switch between dev (Vite server) and prod (manifest.json).

## Directory Structure
```
taraNote_go/
├── cmd/server/main.go      # Entry point
├── internal/
│   ├── config/             # Env handling
│   ├── database/           # GORM connection
│   └── handlers/           # HTTP Handlers
├── resources/
│   ├── js/                 # Copied Vue app
│   └── css/                # Copied styles
├── views/
│   └── app.html            # Inertia Root Template
├── public/                 # Static assets & build output
├── go.mod                  # Go dependencies
├── package.json            # Node dependencies
└── vite.config.js          # Build config
```
