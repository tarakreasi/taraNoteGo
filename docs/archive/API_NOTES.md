# Notes API Documentation

**Version:** 1.0  
**Base URL:** `/api/v1/admin/notes`  
**Authentication:** Required (Bearer Token)



## Table of Contents

1. [Overview](#overview)
2. [Authentication](#authentication)
3. [Endpoints](#endpoints)
4. [Data Models](#data-models)
5. [Error Handling](#error-handling)
6. [Examples](#examples)
7. [Testing](#testing)



## Overview

The Notes API provides endpoints for managing user notes in TaraNote. Notes are rich-text documents that can be organized into notebooks, have different statuses (DRAFT/PUBLISHED), and support full-text search.

**Content Strategy:**
- **In `TaraNote.vue` (Workspace):** Consumed as hierarchical "Notes" for structured learning (OneNote-style).
- **In Dashboard/Portfolio (Digital Garden):** The same records are presented as polished "Articles" or "Digital Garden Nodes".

### Key Features

- Create, read, update, delete notes
- Rich text content (HTML)
- Status management (DRAFT/PUBLISHED)
- Notebook organization
- Full-text search
- Auto-save support
- User authorization



## Authentication

All endpoints require authentication using Laravel Sanctum.

### Headers

```http
Authorization: Bearer {token}
Content-Type: application/json
Accept: application/json
```

### Getting a Token

```bash
# Login to get token
POST /login
{
  "email": "user@example.com",
  "password": "password"
}
```



## Endpoints

### 1. List Notes

Get a paginated list of user's notes with optional filtering.

**Endpoint:** `GET /api/v1/admin/notes`

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `search` | string | No | Search in title and content |
| `status` | string | No | Filter by status (DRAFT/PUBLISHED) |
| `notebook_id` | integer | No | Filter by notebook |
| `page` | integer | No | Page number (default: 1) |
| `per_page` | integer | No | Items per page (default: 15) |

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "title": "My First Note",
      "slug": "my-first-note",
      "content": "<p>Note content...</p>",
      "excerpt": "Note excerpt...",
      "status": "PUBLISHED",
      "notebook_id": 1,
      "user_id": 1,
      "published_at": "2025-12-20T00:00:00.000000Z",
      "created_at": "2025-12-19T00:00:00.000000Z",
      "updated_at": "2025-12-20T00:00:00.000000Z",
      "notebook": {
        "id": 1,
        "name": "Web Development"
      }
    }
  ],
  "links": {...},
  "meta": {...}
}
```



### 2. Create Note

Create a new note.

**Endpoint:** `POST /api/v1/admin/notes`

**Request Body:**

```json
{
  "title": "Untitled Page",
  "content": "",
  "excerpt": "",
  "status": "DRAFT",
  "notebook_id": 1
}
```

**Validation Rules:**

| Field | Rules | Description |
|-------|-------|-------------|
| `title` | string, max:255 | Note title |
| `content` | string, nullable | HTML content |
| `excerpt` | string, nullable, max:500 | Short description |
| `status` | in:DRAFT,PUBLISHED | Note status |
| `notebook_id` | required, exists:notebooks,id | Must be user's notebook |

**Response:** `201 Created`

```json
{
  "data": {
    "id": 2,
    "title": "Untitled Page",
    "slug": "untitled-page",
    "content": "",
    "status": "DRAFT",
    "notebook_id": 1,
    "user_id": 1,
    "published_at": null,
    "created_at": "2025-12-20T01:00:00.000000Z",
    "updated_at": "2025-12-20T01:00:00.000000Z"
  }
}
```



### 3. Get Single Note

Retrieve a specific note.

**Endpoint:** `GET /api/v1/admin/notes/{id}`

**Response:** `200 OK`

```json
{
  "data": {
    "id": 1,
    "title": "My First Note",
    "slug": "my-first-note",
    "content": "<p>Full content here...</p>",
    "excerpt": "Excerpt...",
    "status": "PUBLISHED",
    "notebook_id": 1,
    "user_id": 1,
    "published_at": "2025-12-20T00:00:00.000000Z",
    "created_at": "2025-12-19T00:00:00.000000Z",
    "updated_at": "2025-12-20T00:00:00.000000Z",
    "notebook": {
      "id": 1,
      "name": "Web Development",
      "notes_count": 5
    }
  }
}
```



### 4. Update Note

Update an existing note. Supports partial updates.

**Endpoint:** `PUT /api/v1/admin/notes/{id}`

**Request Body:**

```json
{
  "title": "Updated Title",
  "content": "<p>Updated content...</p>",
  "status": "PUBLISHED"
}
```

**Auto-save Example:**

```json
{
  "content": "<p>User is typing...</p>"
}
```

**Response:** `200 OK`

```json
{
  "data": {
    "id": 1,
    "title": "Updated Title",
    "content": "<p>Updated content...</p>",
    "status": "PUBLISHED",
    "published_at": "2025-12-20T01:05:00.000000Z",
    "updated_at": "2025-12-20T01:05:00.000000Z"
  }
}
```



### 5. Delete Note

Permanently delete a note.

**Endpoint:** `DELETE /api/v1/admin/notes/{id}`

**Response:** `200 OK`

```json
{
  "message": "Note deleted successfully"
}
```



## Data Models

### Note Model

```php
{
  "id": integer,
  "title": string,
  "slug": string,
  "content": string (HTML),
  "excerpt": string,
  "status": enum('DRAFT', 'PUBLISHED'),
  "notebook_id": integer|null,
  "user_id": integer,
  "published_at": datetime|null,
  "created_at": datetime,
  "updated_at": datetime
}
```

### Relationships

- **Belongs To:** User
- **Belongs To:** Notebook (optional)

### Status Values

| Status | Description |
|--------|-------------|
| `DRAFT` | Note is not published, only visible to author |
| `PUBLISHED` | Note is published and visible publicly |



## Error Handling

### Error Response Format

```json
{
  "message": "Error message",
  "errors": {
    "field_name": [
      "Validation error message"
    ]
  }
}
```

### Common HTTP Status Codes

| Code | Meaning | Description |
|------|---------|-------------|
| 200 | OK | Request successful |
| 201 | Created | Resource created successfully |
| 400 | Bad Request | Invalid request data |
| 401 | Unauthorized | Authentication required |
| 403 | Forbidden | User not authorized for this resource |
| 404 | Not Found | Resource not found |
| 422 | Unprocessable Entity | Validation failed |
| 500 | Server Error | Internal server error |

### Example Error Responses

#### Validation Error (422)

```json
{
  "message": "The given data was invalid.",
  "errors": {
    "notebook_id": [
      "The notebook id field is required."
    ],
    "title": [
      "The title must not be greater than 255 characters."
    ]
  }
}
```

#### Authorization Error (403)

```json
{
  "message": "This action is unauthorized."
}
```

#### Not Found (404)

```json
{
  "message": "Note not found."
}
```



## Examples

### Example 1: Create and Publish Note

```javascript
// 1. Create draft note
const createResponse = await axios.post('/api/v1/admin/notes', {
  title: 'My New Article',
  content: '<p>Initial content</p>',
  status: 'DRAFT',
  notebook_id: 1
});

const noteId = createResponse.data.data.id;

// 2. Auto-save while editing (every 1.5s)
await axios.put(`/api/v1/admin/notes/${noteId}`, {
  content: '<p>Updated content while typing...</p>'
});

// 3. Publish when ready
await axios.put(`/api/v1/admin/notes/${noteId}`, {
  status: 'PUBLISHED'
});
```

### Example 2: Search and Filter

```javascript
// Search for notes containing "Laravel"
const searchResults = await axios.get('/api/v1/admin/notes', {
  params: {
    search: 'Laravel',
    status: 'PUBLISHED',
    notebook_id: 1
  }
});

console.log(searchResults.data.data);
```

### Example 3: Dashboard Implementation

```javascript
// Fetch notes for dashboard
const fetchNotes = async () => {
  const response = await axios.get('/api/v1/admin/notes', {
    params: {
      search: searchQuery,
      status: filterStatus !== 'ALL' ? filterStatus : null,
      notebook_id: selectedNotebookId
    }
  });
  
  notes.value = response.data.data;
};

// Auto-save implementation
let timeout = null;
watch([editorTitle, editorContent], () => {
  clearTimeout(timeout);
  timeout = setTimeout(async () => {
    await axios.put(`/api/v1/admin/notes/${noteId}`, {
      title: editorTitle.value,
      content: editorContent.value
    });
  }, 1500); // Auto-save after 1.5s
});
```



## Testing

### Running Tests

```bash
# Run all tests
php artisan test

# Run specific test file
php artisan test tests/Feature/AdminNoteTest.php

# Run specific test
php artisan test --filter test_user_can_create_draft_note
```

### Test Coverage

#### Unit Tests (`tests/Unit/NoteTest.php`)

- [x] Note belongs to notebook
- [x] Note belongs to user
- [x] Published_at casts to datetime
- [x] Note has slug
- [x] Note can be published/draft
- [x] Note has content and excerpt
- [x] Note can have null notebook

#### Feature Tests (`tests/Feature/AdminNoteTest.php`)

- [x] User can create draft note
- [x] User can update note
- [x] User can delete note
- [x] User can list their notes
- [x] User can filter by status
- [x] User can filter by notebook
- [x] User can search notes
- [x] Create requires notebook
- [x] User cannot update others' notes
- [x] User cannot delete others' notes
- [x] Guest cannot access API

### Example Test

```php
/** @test */
public function user_can_create_draft_note(): void
{
    $user = User::factory()->create();
    $notebook = Notebook::factory()->create(['user_id' => $user->id]);

    $response = $this->actingAs($user)->postJson('/api/v1/admin/notes', [
        'title' => 'Untitled Page',
        'notebook_id' => $notebook->id,
        'content' => '',
        'status' => 'DRAFT'
    ]);

    $response->assertStatus(201)
             ->assertJsonFragment([
                 'title' => 'Untitled Page',
                 'status' => 'DRAFT'
             ]);
}
```



## Rate Limiting

API endpoints are rate-limited to prevent abuse.

**Limits:**
- 60 requests per minute for authenticated users
- 10 requests per minute for guests

**Headers:**

```http
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 59
X-RateLimit-Reset: 1703030400
```



## Best Practices

### 1. Auto-save Implementation

```javascript
// Debounce auto-save to avoid excessive requests
const AUTOSAVE_DELAY = 1500; // 1.5 seconds

let saveTimeout = null;
const autoSave = (noteId, data) => {
  clearTimeout(saveTimeout);
  saveTimeout = setTimeout(async () => {
    try {
      await axios.put(`/api/v1/admin/notes/${noteId}`, data);
      showStatus('✓ Saved');
    } catch (error) {
      showStatus('✗ Error saving');
    }
  }, AUTOSAVE_DELAY);
};
```

### 2. Error Handling

```javascript
try {
  const response = await axios.post('/api/v1/admin/notes', noteData);
  return response.data;
} catch (error) {
  if (error.response?.status === 422) {
    // Validation errors
    console.error('Validation failed:', error.response.data.errors);
  } else if (error.response?.status === 403) {
    // Authorization error
    console.error('Unauthorized');
  } else {
    // Other errors
    console.error('Failed to create note:', error.message);
  }
  throw error;
}
```

### 3. Optimistic Updates

```javascript
// Update UI immediately, rollback on error
const deleteNote = async (noteId) => {
  const originalNotes = [...notes.value];
  
  // Optimistic update
  notes.value = notes.value.filter(n => n.id !== noteId);
  
  try {
    await axios.delete(`/api/v1/admin/notes/${noteId}`);
  } catch (error) {
    // Rollback on error
    notes.value = originalNotes;
    alert('Failed to delete note');
  }
};
```



## Changelog

### v1.0 (2025-12-20)
- Initial API release
- CRUD operations
- Search and filtering
- Auto-save support
- Authorization
- Comprehensive tests



**Maintained by:** TaraKreasi Team  
**Last Updated:** December 27, 2025  
**API Version:** 1.0
