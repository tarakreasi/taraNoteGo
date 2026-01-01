# API Specification - Taranote V1

**Base URL:** `/api/v1`
**Content-Type:** `application/json`
**Rate Limits:**
- Public: 60 requests/min
- Login: 5 attempts/min

## A. Public Module (For Blog Readers)
No login required.

### 1. Get All Published Notes
* **Endpoint:** `GET /public/notes`
* **Query Params:**
    * `notebook_slug` (optional): Filter by category.
* **Response (200 OK):**
    ```json
    {
      "data": [
        {
          "title": "How to Install Laravel",
          "slug": "how-to-install-laravel",
          "excerpt": "Complete installation guide...",
          "notebook": "Programming",
          "cover_image": "/storage/img1.jpg",
          "published_at": "2025-10-10 10:00:00"
        }
      ]
    }
    ```

### 2. Read Single Note
* **Endpoint:** `GET /public/notes/{slug}`
* **Response (200 OK):**
    ```json
    {
      "data": {
        "title": "How to Install Laravel",
        "content": "<h1>Hello</h1><p>Article content...</p>",
        "author": "Tara",
        "notebook": "Programming"
      }
    }
    ```



## B. Admin Module (For Writers/You)
**Auth:** Bearer Token (Sanctum).

### 1. Get Notebooks
* **Endpoint:** `GET /admin/notebooks`
* **Purpose:** Retrieve list of categories for dropdown when writing.

### 2. Create Note
* **Endpoint:** `POST /admin/notes`
* **Body:**
    ```json
    {
      "title": "New Article",
      "notebook_id": 1,
      "content": "Draft content...",
      "status": "DRAFT"
    }
    ```
    * **Authorization:** Authenticated users can only create/manage their *own* notes.
    * **Errors:**
        * `403 Forbidden`: Attempting to manage another user's note.
        * `404 Not Found`: Attempting to preview a draft without ownership.

### 3. Upload Image
* **Endpoint:** `POST /admin/upload`
* **Body:** `image` (Multipart Form Data)
* **Response:** `{ "url": "/storage/uploads/image1.jpg" }`