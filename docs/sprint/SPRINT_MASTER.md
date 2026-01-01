# SPRINT MASTER PLAN: Refactoring & Optimization (Revised)

**Objective**: Elevate `taraNote_go` to a production-grade codebase by improving architecture, testability, performance, and maintainability without breaking existing features.

> [!IMPORTANT]
> **Refactoring Strategy**: We will adhere to the "Strangler Fig" inspired approach where applicable—adding tests first, then refactoring module by module. We have split the original plan into **5 Sprints** to prevent the "Big Bang Refactor" risk.

## Roadmap

### Sprint 1: Foundation & Quality Assurance (QA)
**Focus**: Safety Net & Observability.
**Why**: We cannot verify our refactors without a baseline of testing.
- **Goals**:
    - **Test Suite**: Initialize `testify` and write helper utilities for Mock DB.
    - **Coverage**: Write Unit Tests for critical Handlers (`Auth`, `Note`) to lock in current behavior.
    - **Logging**: Implement structured logging (Zap or Slog) to replace standard `log`. Standardize Error responses.
    - **Deliverable**: A test suite that passes green and meaningful logs.

### Sprint 2: Backend Core Architecture
**Focus**: Decoupling & Clean Code.
**Why**: Current Handlers mix HTTP logic with DB queries. This is hard to maintain and test.
- **Goals**:
    - **Service Layer**: Extract business logic from `internal/handlers` into `internal/services`.
    - **Repository Pattern**: (Optional but recommended) Extract DB queries into `internal/repositories` or use distinct scopes.
    - **Validation**: Implement `go-playground/validator` for all Request structs. Stop validating manually in handlers.
    - **Deliverable**: `handlers/` only handle HTTP; `services/` handle logic.

### Sprint 3: Backend Performance & Enhancements
**Focus**: Scalability & Standardization.
**Why**: Client-side search is unscalable. Route definitions are scattered.
- **Goals**:
    - **Server-Side Search**: Implement robust SQL search (`WHERE LIKE` optimized or FTS5) in `internal/services/note_service.go`.
    - **Route Cleanup**: Centralize route names and definitions. Make sure API resource names are consistent.
    - **Middleware**: Review and harden Auth middleware.
    - **Deliverable**: Performant Search API and standardized Routing.

### Sprint 4: Frontend Architecture
**Focus**: Maintainability & Composition.
**Why**: Large Vue components (`TaraNote.vue`, `Docs.vue`) are hard to read and reuse.
- **Goals**:
    - **Composables**: Extract logic into `useNotes`, `useNotebooks`, `useAuth`.
    - **Component Split**: Break down giant "Monolith Components" into smaller functional generic components.
    - **Type Safety**: Improve TypeScript definitions for Props and API responses.
    - **Deliverable**: Smaller, cleaner Vue components.

### Sprint 5: Frontend DX, Polish & Release
**Focus**: Developer Experience & Final Polish.
**Why**: Hardcoded URLs in Frontend are fragile. UI needs to match "Zen Glass".
- **Goals**:
    - **Route Helper**: Implement formatted route passing (Go -> Inertia) to avoid hardcoded strings (Ziggy-lite).
    - **UI Audit**: Verify "Zen Glass" aesthetics, dark mode consistency, and Mobile responsiveness.
    - **Docker**: Final Multi-stage build optimization.
    - **Release**: Tag v1.1.0.
