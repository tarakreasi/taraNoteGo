# Changelog

All notable changes to this project will be documented in this file.

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
