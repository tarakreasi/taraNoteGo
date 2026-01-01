# Backend Architecture Refactor
**Date:** December 27, 2025
**Scope:** API Controllers, Requests, and Type Safety

## Overview
This document details the architectural refactor executed to elevate the TaraNote backend to "Senior Engineering" standards. The primary goals were to enforce strict type safety, separate concerns between resources, and eliminate "magic" values in favor of defined constants.

## Key Changes

### 1. Separation of Concerns (Controllers)
Previously, `AdminNoteController` handled both Note and Notebook logic, leading to a "fat controller" anti-pattern.

**Refactor:**
- **`AdminNoteController`**: Now strictly handles `Note` CRUD operations.
- **`AdminNotebookController`**: A new, dedicated controller for `Notebook` (Space) management.

This separation improves readability and allows for independent scaling of logic.

### 2. Strict Type Safety
All refactored files now adhere to strict typing standards to prevent runtime type errors and improve IDE support.

- **Directive:** `declare(strict_types=1);` enabled at the top of every file.
- **Return Types:** All Controller methods now explicitly declare `: JsonResponse`.
- **Arguments:** All method arguments are type-hinted (e.g., `int $id`).

### 3. Form Request Validation
Inline validation (e.g., `$request->validate([...])`) inside controllers has been replaced with dedicated **Form Request** classes. This keeps controllers "thin" and centralizes validation rules.

**New Request Classes:**
- `App\Http\Requests\StoreNoteRequest`
- `App\Http\Requests\UpdateNoteRequest`
- `App\Http\Requests\StoreNotebookRequest`
- `App\Http\Requests\UpdateNotebookRequest`

### 4. Domain Constants
Hardcoded string literals (Magic Strings) were removed to prevent typos and enable easier refactoring.

**Model Updates (`App\Models\Note`):**
- `Note::STATUS_DRAFT` (Replaces 'DRAFT')
- `Note::STATUS_PUBLISHED` (Replaces 'PUBLISHED')
- `Note::STATUS_ARCHIVED` (Replaces 'ARCHIVED')

## Code Example

**Before (Junior Pattern):**
```php
public function store(Request $request) {
    if ($request->status == 'PUBLISHED') { ... } // Magic string
}
```

**After (Senior Pattern):**
```php
public function store(StoreNoteRequest $request): JsonResponse {
    $validated = $request->validated();
    if ($validated['status'] === Note::STATUS_PUBLISHED) { ... } // Constant
}
```

## Route Definitions
Routes in `routes/api.php` have been updated to point to the new split controllers.

```php
// Note Management
Route::get('/notes', [AdminNoteController::class, 'index']);

// Notebook Management (New)
Route::get('/notebooks', [AdminNotebookController::class, 'index']);
```
