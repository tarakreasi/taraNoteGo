# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased] - Sprint 6 (Bug Fixes & Refinement)
### Fixed
- **UI Navigation**: Resolved stacking issue by wrapping page templates in single root element.
- **Notebooks**: Fixed `unique constraint` error by auto-generating slugs.
- **Frontend**: Fixed `route()` helper error by adding global mixin.
- **Assets**: Replaced broken placeholder images with `placehold.co`.
- **Accessibility**: Added missing form attributes (`id`/`name`) to inputs.
- **Data Integrity**: Fixed "Uncategorized Notes" bug by ensuring `notebook_id` is saved on note creation.
- **Data Integrity**: Fixed empty slugs in legacy notebooks.
- **Dashboard**: Fixed note filtering by `notebook_id` and `status`.
- **Auth**: Fixed incorrect guest state on public pages (Home/TaraNote).
- **Docs**: Resolved 404 errors for README and unreachable code warning.

## [Unreleased] - Sprint 5 (Polish & Deployment)
### Added
- Frontend `route()` helper (Ziggy-lite).
- Replaced hardcoded URLs with named routes.

## [Unreleased] - Sprint 4 (Frontend Architecture)
### Changed
- Extract Note logic to `useNotes` composable.
- Optimized `TaraNote.vue` component.

## [Unreleased] - Sprint 3 (Enhancements)
### Added
- Server-side Search for Notes (`LIKE %query%`).
- Route cleanup.

## [Unreleased] - Sprint 2 (Core Architecture)
### Added
- Service Layer (`internal/services`).
- Request Validation (`go-playground/validator`).
- Refactored `Note` and `Auth` logic to services.

## [Unreleased] - Sprint 1 (Foundation)
### Added
- Testing framework (`testify`) and helper utilities (`internal/testutils`).
- Characterization tests for Authentication (`internal/handlers/auth_test.go`).
- Global Error Handler in `cmd/server/main.go`.

## [v1.0.0] - 2026-01-01
### Added
- **Migration**: Complete rewrite of backend from Laravel to Go (Fiber).
- **Frontend**: Successfully ported Vue 3 + Inertia.js frontend.
- **Docker**: Added `Dockerfile` for lightweight production builds.
- **Models**: `User`, `Notebook`, `Note`, `Setting` (GORM).
- **Public**: Home, Article, and TaraNote Viewers.
- **Admin**: Full CRUD for Notes and Notebooks.
- **Settings**: System-wide key-value settings management.

### Changed
- **Session**: Replaced Laravel sessions with Fiber `session` (cookie-based).
- **Build**: Added `Makefile` for streamlined development commands.
- **Database**: Switched to SQLite as the primary supported database for this release.
