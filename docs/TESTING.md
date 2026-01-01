# Testing Guide

This project uses **PHPUnit** for backend testing. We have a mix of Unit tests and Feature tests to ensure system stability.

## Prerequisites
Ensure detailed dependencies are installed:
```bash
composer install
```

## Running Tests
To run the full test suite:
```bash
php artisan test
# Or use PHPUnit directly
./vendor/bin/phpunit --testdox
```

## Test Structure

### 1. Feature Tests (`tests/Feature`)
Integration tests that verify full HTTP requests and responses.
- **`NoteAuthorizationTest.php`**: Crucial for security. Verifies:
    - Users can only edit *their own* notes.
    - Public cannot view `DRAFT` notes.
    - Owners *can* preview `DRAFT` notes.
- **`AdminNotebookTest.php`**: Verifies notebook (section) creation logic.
- **`AdminNoteTest.php`**: Verifies draft note creation logic.
- **`ColabTest.php`**: Verifies Colab Deck creation, updates, and token-based access.
### `PortfolioTest.php`
- **Covers:** Portfolio page access via custom username or slug.
- **Tests:**
  - Portfolio can be accessed via username
  - Portfolio falls back to slugified name
  - 404 is returned for non-existent users

### `SubscriptionTest.php`
- **Covers:** Newsletter subscription functionality.
- **Tests:**
  - User can subscribe with valid email
  - Duplicate emails show info message
  - Invalid emails are rejected
  - Email field is required
- **`Auth/`**: Verifies Login, Registration, and Password Reset flows.

### 2. Unit Tests (`tests/Unit`)
Isolated tests for specific methods or classes.
- **`NoteTest.php`**: Verifies model relationships and casting.

## Continuous Integration (CI)
It is recommended to run `php artisan test` before every commit or deploy.

## Current Test Coverage

**68 tests, 185 assertions**

### Recent Improvements
- Enhanced Note model with query scopes
- Improved validation in AdminNoteController
- Better status transition handling (DRAFT <-> PUBLISHED)
- Image upload MIME type restrictions

## Resources
- [Laravel Testing Docs](https://laravel.com/docs/12.x/testing)
- [PHPUnit Manual](https://phpunit.de/manual/current/en/index.html)
