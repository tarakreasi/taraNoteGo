# API Reference

**TaraNote API Documentation (Beta 0.9.80)**  
**UI Design System: V6 "Electric Earth"**  
This document provides a comprehensive reference for integrating with TaraNote's public and authenticated APIs.



## Base URL

```
https://your-domain.com/api/v1
```

## Authentication

Most endpoints require authentication via Laravel Sanctum tokens.

### Obtaining a Token

```http
POST /api/v1/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password"
}
```

**Response:**
```json
{
  "token": "1|abc123...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "user@example.com"
  }
}
```

**Usage:**
```
Authorization: Bearer {token}
```



## Public Endpoints

### 1. Newsletter Subscription

Subscribe an email address to the newsletter.

**Endpoint:** `POST /newsletter/subscribe`

**Request:**
```json
{
  "email": "subscriber@example.com"
}
```

**Response (Success - New):**
```json
{
  "message": "Thanks for subscribing! You are now on the list."
}
```

**Response (Info - Duplicate):**
```json
{
  "message": "You are already subscribed. Thanks for checking in!"
}
```

**Validation Errors:**
```json
{
  "message": "The email field is required.",
  "errors": {
    "email": ["The email field is required."]
  }
}
```



### 2. Get Published Articles

Retrieve a paginated list of published articles.

**Endpoint:** `GET /api/v1/articles`

**Query Parameters:**
- `page` (int): Page number (default: 1)
- `per_page` (int): Items per page (default: 15, max: 50)
- `notebook` (string): Filter by notebook slug
- `search` (string): Search in title and content

**Example:**
```
GET /api/v1/articles?page=1&per_page=10&search=laravel
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "title": "Getting Started with Laravel",
      "slug": "getting-started-with-laravel",
      "excerpt": "A beginner's guide to Laravel...",
      "cover_image": "/storage/covers/laravel.jpg",
      "published_at": "2025-12-01T10:00:00Z",
      "notebook": {
        "id": 1,
        "name": "Tutorials",
        "slug": "tutorials"
      },
      "user": {
        "id": 1,
        "name": "John Doe",
        "username": "johndoe"
      }
    }
  ],
  "links": {...},
  "meta": {...}
}
```



### 3. Get Single Article

**Endpoint:** `GET /api/v1/articles/{slug}`

**Example:**
```
GET /api/v1/articles/getting-started-with-laravel
```

**Response:**
```json
{
  "id": 1,
  "title": "Getting Started with Laravel",
  "slug": "getting-started-with-laravel",
  "content": "<p>Full HTML content...</p>",
  "excerpt": "A beginner's guide...",
  "cover_image": "/storage/covers/laravel.jpg",
  "published_at": "2025-12-01T10:00:00Z",
  "notebook": {...},
  "user": {...}
}
```



## Authenticated Endpoints

### 4. Get User's Notes (Private)

Retrieve the authenticated user's notes (including drafts).

**Endpoint:** `GET /api/v1/admin/notes`

**Headers:**
```
Authorization: Bearer {token}
```

**Query Parameters:**
- `status` (string): Filter by status (`draft`, `published`, `archived`)
- `notebook_id` (int): Filter by notebook
- `search` (string): Search query

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "title": "Draft Article",
      "slug": "draft-article",
      "status": "draft",
      "created_at": "2025-12-20T14:30:00Z",
      "updated_at": "2025-12-21T09:00:00Z"
    }
  ]
}
```



### 5. Create Note

**Endpoint:** `POST /api/v1/admin/notes`

**Headers:**
```
Authorization: Bearer {token}
Content-Type: application/json
```

**Request:**
```json
{
  "notebook_id": 1,
  "title": "My New Article",
  "content": "<p>Article content...</p>",
  "excerpt": "Short summary",
  "status": "draft"
}
```

**Response:**
```json
{
  "id": 42,
  "title": "My New Article",
  "slug": "my-new-article",
  "status": "draft",
  "created_at": "2025-12-21T09:22:00Z"
}
```



### 6. Update Note

**Endpoint:** `PUT /api/v1/admin/notes/{id}`

**Request:** Same as Create Note

**Response:** Updated note object



### 7. Delete Note

**Endpoint:** `DELETE /api/v1/admin/notes/{id}`

**Response:**
```json
{
  "message": "Note deleted successfully"
}
```



## Notebook Management

### 8. Get Notebooks (Spaces)

**Endpoint:** `GET /api/v1/admin/notebooks`

**Response:** List of notebooks with note counts.

### 9. Create Notebook (Space)

**Endpoint:** `POST /api/v1/admin/notebooks`

**Request:**
```json
{
  "name": "My New Space"
}
```

### 10. Update Notebook

**Endpoint:** `PUT /api/v1/admin/notebooks/{id}`

**Request:**
```json
{
  "name": "New Name",
  "slug": "new-slug" // Optional
}
```

### 11. Delete Notebook (Space)

**Endpoint:** `DELETE /api/v1/admin/notebooks/{id}`

**Rules:**
- Can only delete **empty** notebooks (0 notes).
- Returns 403 if notebook is not empty.

**Response:**
```json
{
  "message": "Notebook deleted"
}
```



### 12. Upload Image

**Endpoint:** `POST /api/v1/admin/settings/upload`

**Headers:**
```
Authorization: Bearer {token}
Content-Type: multipart/form-data
```

**Request:**
```
file: (binary)
type: "avatar" | "hero" | "cover"
```

**Response:**
```json
{
  "url": "/storage/uploads/avatar_123.jpg"
}
```



## Settings API

### 13. Get Settings

**Endpoint:** `GET /api/v1/admin/settings`

**Response:**
```json
{
  "welcome_hero_title": "Deep dives into Code & Creation",
  "welcome_hero_subtitle": "A living collection...",
  "footer_copyright_text": "© 2026 TaraKreasi"
}
```



### 14. Update Settings

**Endpoint:** `POST /api/v1/admin/settings`

**Request:**
```json
{
  "settings": {
    "welcome_hero_title": "New Title",
    "footer_copyright_text": "© 2026 New Text"
  }
}
```



## Error Responses

All errors follow this structure:

**400 Bad Request:**
```json
{
  "message": "Validation error",
  "errors": {
    "field": ["Error message"]
  }
}
```

**401 Unauthorized:**
```json
{
  "message": "Unauthenticated."
}
```

**403 Forbidden:**
```json
{
  "message": "This action is unauthorized."
}
```

**404 Not Found:**
```json
{
  "message": "Resource not found."
}
```

**500 Server Error:**
```json
{
  "message": "Server error.",
  "error": "Detailed error message (dev mode only)"
}
```



## Rate Limiting

- **Public Endpoints:** 60 requests per minute
- **Authenticated Endpoints:** 120 requests per minute
- **Newsletter Subscription:** 10 requests per minute per IP

Rate limit headers:
```
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 59
X-RateLimit-Reset: 1640000000
```



## Webhooks (Future)

Webhook support is planned for:
- New subscriber notifications
- Article published events
- Comment notifications



## SDKs & Libraries

Official libraries:
- **JavaScript/Node.js:** Coming soon
- **PHP:** Use Laravel HTTP Client with Bearer tokens



## Support

- **Issues:** [GitHub Issues](#)
- **Email:** ajarsinau@gmail.com
- **Docs:** [https://tarakreasi.com/docs](https://tarakreasi.com/docs)



**Last Updated:** December 27, 2025 (Beta 0.9.80 - UI Design V6.5)
