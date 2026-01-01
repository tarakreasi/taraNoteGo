# System Architecture

**Version:** 1.0 (Production)
**Design Philosophy:** "Monolithic Integrity" - Everything needed to run the system is contained within.

---

## Architectural Overview

TaraNote follows a **Modern Monolith** architecture. Unlike microservices which introduce network latency and distributed complexity, TaraNote maximizes the speed and type-safety of a single codebase.

### The Stack: "The Integrated Circuit"
*   **Backend (Logic):** Laravel 12 (PHP 8.2+)
*   **Frontend (UI):** Vue 3 (Composition API)
*   **Bridge:** Inertia.js (Removes the need for a REST API layer)
*   **Database:** MySQL / SQLite (Data Integrity)

---

## Backend Design: The "Zero Trust" Model

As a System Integrator, I treat user input as "untrusted external signals". The backend enforces strict boundaries.

### 1. Dedicated Controllers (Separation of Concern)
We avoid "God Classes". Logic is strictly strictly separated:
*   **`AdminNoteController`**: Handles individual Note logic.
*   **`AdminNotebookController`**: Handles Notebook (Category) logic.
*   **`PublicController`**: Read-only access for the blog interface.

### 2. Strict Type Safety
We do not rely on PHP's loose typing. Every file enables strict types to catch errors at the "compiler" level (IDE analysis) rather than runtime.
```php
declare(strict_types=1);

public function store(StoreNoteRequest $request): JsonResponse
{
    // ...
}
```

### 3. Data Integrity & Validation
*   **FormRequest Objects:** Controllers never see raw `$request->all()`. Data must pass a `FormRequest` validation gate first.
*   **Constants over Magic Strings:** We use `Note::STATUS_PUBLISHED` instead of `'PUBLISHED'` to ensure extensive find-and-replace safety.

---

## Database Schema (Relation Map)

The database is designed for **Referential Integrity** using Foreign Keys with Cascade rules where appropriate.

### Core Entities

#### 1. Users (`users`)
The central entity. Currently designed for a single-tenant (Portfolio) or multi-tenant (SaaS) future.

#### 2. Notebooks (`notebooks`)
*   **Role:** The "Folder" metaphor.
*   **Relation:** A User has minimal 1 Notebook.
*   **Constraint:** Deleting a Notebook is **restricted** if it contains Notes (Data Safety precaution).

#### 3. Notes (`notes`)
*   **Role:** The content unit (Article).
*   **Relation:** Belongs to 1 Notebook.
*   **Key Fields:**
    *   `content` (JSON/LongText): Stores the Tiptap structure.
    *   `status`: State machine (`DRAFT` -> `PUBLISHED` -> `ARCHIVED`).
    *   `slug`: Unique index for SEO-friendly URLs.

---

## The "Headless" Editor Strategy

We utilize **Tiptap** as a headless editor framework.

*   **Why?** Traditional editors (like CKEditor) save rigid HTML blobs.
*   **Benefit:** Tiptap saves a JSON schema. This allows us to render the content differently on the "Admin Editor" (dense, technical) vs the "Public Blog" (spacious, readable) by interpreting the JSON data differently.

---

## Public vs Private Separation

The application has two distinct "Zones":

1.  **The Workshop (`/admin`)**:
    *   Protected by Laravel Sanctum/Auth.
    *   Dense UI (Data tables, forms).
    *   Vue Components: `AdminLayout`, `NoteEditor`.

2.  **The Showroom (`/`)**:
    *   Publicly cacheable.
    *   Focus on Typography and Readability.
    *   Vue Components: `BlogLayout`, `ArticleCard`.

This separation ensures that administrative weight does not slow down the public reading experience.

---

## Performance Metrics (Benchmarked 2025-12-27)

*   **Test Suite:** 64 Passing Tests (Critical Path 100% Covered)
*   **Backend Latency:** ~20ms (Local Bare Metal)
*   **Frontend Assets:** < 300kb Gzipped (Vue 3 Tree-Shaking)
*   **Database Queries:** N+1 Free (Enforced by strict Eager Loading)
