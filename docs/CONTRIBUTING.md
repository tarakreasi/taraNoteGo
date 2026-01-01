# Contributing to TaraNote

Thank you for your interest in contributing!

> **Contribution Policy**: This is a **Portfolio Project** actively maintained by the author. We value quality over speed.
 > - **High Standards**: We maintain a strict "Green Build" policy (all 64 tests must pass).
 > - **Thoughtful Review**: PRs are reviewed for architectural consistency, not just functionality.
 > - **No Pressure**: This is an open-source tool for the community. There are no deadlines, only shared goals.

If you're okay with these limitations, we'd love your help! This guide sets the standards for contributing to the TaraNote UI Design V6 codebase.



## Development Environment

### Prerequisites
- PHP 8.2+
- Node.js 18+
- SQLite (for local dev)

### Setup Commands
```bash
# 1. Clone & Install
git clone <repo>
composer install
npm install

# 2. Environment
cp .env.example .env
touch database/database.sqlite
php artisan key:generate
php artisan migrate --seed

# 3. Running
# Terminal 1:
php artisan serve
# Terminal 2:
npm run dev
```



## Coding Standards (V6 "Electric Earth")

### Laravel (Backend)
- **Controllers:** Keep lean. Business logic for complex items goes into Services or Models.
- **Naming:** `PascalCase` for Classes, `camelCase` for variables/methods.
- **Type Hinting:** Strictly use Return Types (`: View`, `: RedirectResponse`).

### Vue.js (Frontend)
- **Composition API:** Use `<script setup>` syntax exclusively.
- **Props:** Define generic props clearly. Use `defineProps` with type validation.
- **Tailwind:** Use utility classes. Avoid custom CSS unless for animations/noise.
- **Icons:** Use `Material Symbols` via span classes, NOT SVG imports (ADR-004 compliance).



## Testing Policy
We maintain a strict "Green Build" policy.
- **Run Tests:** `php artisan test`
- **Requirement:** All PRs must pass the existing 73 Unit/Feature tests.
- **New Features:** Must include a matching Feature Test case (e.g., `tests/Feature/NewThingTest.php`).



## Pull Request Process
1.  **Branch:** `feature/your-feature-name`
2.  **Commit:** Conventional Commits (`feat: add avatar upload`, `fix: header z-index`)
3.  **Review:** Request review from @SeniorEng.



Happy Coding!
