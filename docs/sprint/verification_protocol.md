# Verification Protocol

**Rule #1: Code is Truth.**
Before refactoring any documentation, you must **READ THE CODE** first. Do not hallucinate features.

---

## Verification Steps

### 1. Check Routes
Read `routes/web.php` and `routes/api.php`.
*   *Check:* Is the endpoint actually there?
*   *Check:* Does it return JSON or a View (Inertia)?

### 2. Check Controllers
Read `app/Http/Controllers/*.php`.
*   *Check:* What validation rules exist? (These are your required fields).
*   *Check:* What serves the response? (e.g., `Inertia::render` vs `response()->json`).

### 3. Check Models
Read `app/Models/*.php`.
*   *Check:* What is in `$fillable`? (These are your writable attributes).

### 4. Check Cross-File Consistency
*   *Check:* Contact info (email, LinkedIn, GitHub) must be identical across ALL files.
*   *Check:* Version numbers mentioned in multiple docs must match.
*   *Check:* Tech stack badges in README must match `composer.json` / `package.json`.

### 5. Check Tech Stack Versions
*   *Check:* Read `composer.json` for actual PHP & Laravel versions.
*   *Check:* Read `package.json` for actual React, Vite, Tailwind versions.
*   *Check:* Compare badges in README/docs against actual installed versions.
*   *Example:* If `package.json` says `"react": "^19.0.0"`, docs must say "React 19", not "React 18".

### 6. Date Preservation
*   **DO NOT** change existing dates (e.g., "Last Updated: December 21, 2025") unless explicitly instructed.
*   If you find conflicting dates across documents, **STOP** and ask the user which date is correct.
*   New documents you create should use the current date from the system.

### 7. Style Compliance (Strict)
*   **NO EMOJIS:** Ensure no emojis exist in titles or body text.
*   **NO FILLER:** Remove distinct "Corporate Fluff". Keep it direct and technical.
*   **Professional Tone:** Use "We" or "I" consistently (Engineers Journey = "I").

---

## Action on Discrepancy

If the old documentation says "X" but the code says "Y":
*   **Update the Documentation to match "Y" (Real Code).**
*   *Exception:* If the discrepancy seems like a bug, flag it in a `NOTE` block.
