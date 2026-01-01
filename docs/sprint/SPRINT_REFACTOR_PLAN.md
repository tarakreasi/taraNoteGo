# SPRINT 1: Foundation & Quality Assurance

**Goal**: Establish a robust testing and logging foundation. **No feature refactoring** happens here; only characterizing existing behavior with tests.

## 1. Testing Framework Setup
> [!NOTE]
> We will use `testify` for assertions and mocking, as it is the industry standard for Go.

- [ ] **Install Dependencies**:
  ```bash
  go get github.com/stretchr/testify
  ```
- [ ] **Test Utilities (`internal/testutils/`)**:
    - `setup.go`: Function to spin up an in-memory SQLite DB (`:memory:`) and run auto-migrations.
    - `request.go`: Helper function `MakeRequest(method, url, body, token)` to simplify handler testing.

## 2. Characterization Tests (Backend)
*Write tests that pass with the CURRENT code to document behavior.*
- [ ] **Auth Tests** (`internal/handlers/auth_test.go`):
    - [ ] `POST /login`: Successful login returns token/session.
    - [ ] `POST /login`: Invalid password returns 401/422.
- [ ] **Note Tests** (`internal/handlers/note_test.go`):
    - [ ] `GET /api/notes`: Returns list of notes (check JSON structure).
    - [ ] `POST /api/notes`: Creates a note and returns 201.
    - [ ]Check that `user_id` is enforced (cannot read others' notes).

## 3. Infrastructure & Reliability
- [ ] **Structured Logging**:
    - Replace `log.Println` with a custom Logger or `slog` (Go 1.21+ built-in).
    - Ensure logs include: `time`, `level`, `method`, `path`, `status`, `duration`.
- [ ] **Global Error Handler**:
    - In `cmd/server/main.go`, ensure `fiber.Config{ErrorHandler: ...}` is set.
    - Catch panics and return generic 500 JSON to client (hide stack trace in production).

## 4. Configuration Guardrails
- [ ] **Env Validation**:
    - Create `internal/config/validate.go`.
    - Fail startup if `DB_NAME` or `APP_KEY` is missing.
