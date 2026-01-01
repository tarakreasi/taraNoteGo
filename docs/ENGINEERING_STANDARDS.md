# Engineering Standards & Manifesto
**"Code is a liability. Less code, better logic."**

This document outlines the golden rules for transitioning from a "Get it working" mindset (Junior) to a "Get it maintainable" mindset (Senior).

---

## 1. Zero Trust Architecture (Defensive Coding)

### The Rule
Never trust input, never trust the frontend, never trust yourself.

### The Practice
*   **Strict Types**: Always use `declare(strict_types=1);` in PHP. If a function expects an `int`, enforce it. PHP is loose; you shouldn't be.
*   **Return Types**: Every function MUST declare what it returns (`: JsonResponse`, `: void`, `: array`).
*   **Validations**: Never use `$request->all()`. Always allow-list inputs via Form Requests.

**Why?**
Type errors caught in the IDE effectively prevent runtime production crashes.

---

## 2. Commit Like a Surgeon (Git Hygiene)

### The Rule
Your commit history is a story. Don't write a messy one.

### The Practice
*   **Atomic Commits**: One feature/fix = One commit. Do not combine "Fix Login" and "Update CSS" in the same commit.
*   **Commit First**: Run `git init` before you write the first line of code.
*   **Branching**: Never push directly to `main`. Use `feature/login-fix` or `fix/typo`. Merge via Pull Request.
*   **Message**: `Fix bugs` is forbidden. Use `Fix: Prevent null pointer exception in AuthController`.

**Why?**
When everything breaks, `git bisect` is your only friend. It only works if commits are atomic.

---

## 3. Death to Magic Systems

### The Rule
Explicit is better than implicit.

### The Practice
*   **No Magic Strings**: Don't type `'PUBLISHED'` or `'admin'` in your code.
    *   *Bad*: `if ($role == 'admin')`
    *   *Good*: `if ($role === User::ROLE_ADMIN)`
*   **No Magic Numbers**: `if ($status == 1)` -> What is 1? Active? Pending? Deleted?
    *   *Good*: `if ($status === Order::STATUS_ACTIVE)`
*   **No "God" Controllers**: If a Controller has more than 5-7 methods, split it. (e.g., `AdminNoteController` vs `AdminNotebookController`).

**Why?**
Searching for usage of `User::ROLE_ADMIN` finds 100% of cases. Searching for `'admin'` finds comments, unrelated strings, and garbage.

---

## 4. The "Boy Scout" Rule

### The Rule
Always leave the code cleaner than you found it.

### The Practice
*   Refactor *as* you work, not *after* you finish (because you never will).
*   If you see inconsistent indentation, fix it.
*   If you see a variable named `$d`, rename it to `$date`.

**Why?**
Technical debt accumulates like dust. You must sweep daily.

---

## 5. Documentation is Not Optional

### The Rule
Code tells you *how*, comments tell you *why*. Documentation tells you *what*.

### The Practice
*   **Self-Documenting Code**: `calculate($a, $b)` is bad. `calculateTotalTax($subtotal, $taxRate)` is good.
*   **The "Bus Factor"**: If you get hit by a bus tomorrow, can someone else deploy your app? If no, write a `README.md` and `DEPLOYMENT.md`.
*   **Zero Emotion**: Keep docs professional. Future you (or your boss) doesn't care if you were "excited" or "frustrated". They just need the facts.

---

## 6. Design First, Code Second

### The Rule
Typing is the most expensive way to think.

### The Practice
*   **Mockups**: Don't build UI in code until you've sketched it (even on paper). "Bento Grid" vs "List View" decisions should happen before writing CSS.
*   **Schema**: Plan your database tables on a whiteboard. Renaming columns after production deployment is a nightmare.

---

## 7. Respect the Folder Structure

### The Rule
A place for everything, and everything in its place.

### The Practice
*   **Assets**: Images go in `/images`, not root.
*   **Logic**: Business logic goes in Services/Models, not Controllers.
*   **Views**: Don't mix admin views with public views. `pages/Admin/` vs `pages/Public/`.

---

**Signed,**


---

## 8. Technical Standards Reference

### PHP / Laravel (PSR-12)
*   **Strict Types**: `declare(strict_types=1);` on every file.
*   **Naming**: `PascalCase` for Classes, `camelCase` for methods/variables, `snake_case` for database columns.
*   **Models**: Use `$fillable` (never `$guarded`). Use Constants for rigid states (e.g. `Note::STATUS_DRAFT`).

### Vue 3 / TypeScript
*   **Composition API**: Use `<script setup lang="ts">`.
*   **Components**: `PascalCase` filenames (e.g., `ArticleCard.vue`).
*   **Props**: Define explicit interfaces.
    ```typescript
    interface Props {
        note: Note;
        editable?: boolean;
    }
    ```

### Git Protocol (Conventional Commits)
*   `feat:` New features
*   `fix:` Bug fixes
*   `refactor:` Code change that neither fixes a bug nor adds a feature
*   `docs:` Documentation only changes
