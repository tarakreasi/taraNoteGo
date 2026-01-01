# QA Strategy: TaraNote Go (v1.0.0)

This document outlines the Quality Assurance strategy for the **TaraNote Go** port. Unlike the migration plan (`qa.md`), this document focuses on validating the *production-readiness* of the implemented Go/Fiber/Inertia stack.

## [Philosophy] "Strictness meets Speed"
We have moved from a dynamic, interpreted runtime (PHP/Laravel) to a static, compiled, and concurrent runtime (Go/Fiber). 
-   **Strictness:** Every type must match. Logic errors are caught at compile time or strictly at runtime.
-   **Speed:** HTML rendering and JSON serialization are microseconds, not milliseconds.

## [Architecture] Verification
Before functional testing, verify the architectural pillars:

| Pillar | Technology | Status Check |
| :--- | :--- | :--- |
| **Protocol** | Inertia.js (JSON/HTML Hybrid) | Must handle `X-Inertia` headers correctly (Generic JSON vs HTML Wrapper). |
| **Backend** | Go (Fiber v2) + GORM | Middleware chain (CORS, Recover, Logger) must be robust. |
| **Database** | SQLite + GORM | Relationships (Notebook -> Notes) must include Foreign Keys constraints. |
| **Frontend** | Vue 3 + Tailwind CSS | "Zen Glass" design system must render consistently across routes. |

---

## [Plan] Master Test Plan

### [Phase 1] Core Integrity (The Foundation)
**Goal:** Ensure the backend reliably serves data and handles the Inertia protocol without "white screens" or 500 errors.

| ID | Scenario | Steps | Expected Result |
| :-- | :--- | :--- | :--- |
| **C1-01** | **Inertia Protocol Compliance** | 1. Visit `/` (HTML request).<br>2. Click internal link (AJAX request). | 1. Returns full HTML document.<br>2. Returns JSON with `X-Inertia: true`. |
| **C1-02** | **Authentication State** | 1. Login.<br>2. Navigate to Home (Public).<br>3. Check user avatar/floating dock. | Public pages must detect the authenticated session and show "Dashboard" button, NOT "Login". |
| **C1-03** | **Data Relationships** | 1. Create a Note inside "Project TaraNote".<br>2. Check Dashboard filter.<br>3. Check DB (`sqlite3`). | Note appears ONLY in "Project TaraNote" list. Database `notebook_id` matches. |
| **C1-04** | **Error Handling** | 1. Go to `/docs/invalid-page`. | Returns custom 404 "Zen Glass" page, NOT a raw JSON error or standard text. |

### [Phase 2] UI/UX (The "Zen Glass" Experience)
**Goal:** Verify visual fidelity and interaction implementation.

| ID | Scenario | Steps | Expected Result |
| :-- | :--- | :--- | :--- |
| **U2-01** | **Dashboard Filtering** | 1. Click "Design Journey" in Sidebar.<br>2. Click "ALL Notes". | Middle column filters instantly. "ALL" shows total count. |
| **U2-02** | **Public Sidebar Counts** | 1. Create a DRAFT note in "Design Journey".<br>2. Logout/Go to Public TaraNote page. | Sidebar count for "Design Journey" should NOT include the Draft. |
| **U2-03** | **Editor Persistance** | 1. Type in Tiptap editor.<br>2. Wait 2s (Auto-save).<br>3. Refresh. | Content persists exactly as typed (HTML safe). |
| **U2-04** | **Responsive Mobile** | 1. Shrink browser to <768px.<br>2. Open Sidebar Drawer. | Drawer slides in smoothly. "Glass" effect remains performant. |

### [Phase 3] Performance & Concurrency
**Goal:** Validate Go's ability to handle load better than the previous stack.

| ID | Scenario | Steps | Expected Result |
| :-- | :--- | :--- | :--- |
| **P3-01** | **Zero-Downtime Reload** | 1. Edit code.<br>2. `go run` recompiles.<br>3. Refresh page immediately. | Startup time < 500ms. No "Waiting for server" lag. |
| **P3-02** | **Concurrent Reads** | 1. Open 10 tabs of `/taranote` simultaneously. | All load instantly (< 100ms backend time). No DB locks. |
| **P3-03** | **Memory Stability** | 1. Check RAM usage (`top`). | Should be < 20MB for the compiled binary (vs >100MB for PHP-FPM). |

### [Phase 4] Deployment Readiness
**Goal:** Ensure the artifacts are ready for production (VPS).

| ID | Scenario | Steps | Expected Result |
| :-- | :--- | :--- | :--- |
| **D4-01** | **Build Binary** | Run `go build -o bin/server cmd/server/main.go`. | Creates a standalone executable. Zero compiler errors. |
| **D4-02** | **Environment Config** | Run binary with `.env.production`. | Respects `PORT`, `DB_DATABASE`, and `APP_ENV` settings. |
| **D4-03** | **Static Assets** | Verify `public/build/manifest.json`. | Vite manifest maps JS/CSS chunks correctly for production mode. |

---

## [Tools] Testing Tools & Commands

### 1. Functional Testing (Manual)
Run the diverse development server:
```bash
npm start
# (Runs 'go run cmd/server/main.go' AND 'npm run dev')
```

### 2. Database Inspection
Use the custom inspection tool for quick sanity checks:
```bash
go run cmd/inspect/main.go
# (Note: Ensure cmd/inspect exists or create it using the task script)
```

### 3. Production Simulation
Test the binary build (closest to VPS reality):
```bash
go build -o bin/server cmd/server/main.go
./bin/server
```

### 4. Code Quality
Run the Go linter/vet:
```bash
go vet ./...
```

---

**Last Updated:** January 1, 2026
**Target Release:** v1.0.0
