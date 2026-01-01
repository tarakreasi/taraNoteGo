# Database Schema & Data Dictionary
**Engine:** SQLite/MySQL | **Design:** Monolith Modular | **Version:** 6.0

## 1. Table: `users`
Blog owner / Admin data. Includes profile and hero customization fields.

| Column | Type | Description |
| :--- | :--- | :--- |
| `id` | Integer (PK) | Unique ID. |
| `name` | String | Author display name. |
| `email` | String | Login credential (Unique). |
| `password` | String | Hashed password. |
| `username` | String (Null) | **[New V6]** Custom handle for portfolio URL (Unique). |
| `bio` | Text (Null) | Short author description for article footer. |
| `avatar` | String (Null)| Path/URL to profile photo. |
| `hero_tagline` | String (Null) | **[New V6]** Job Title / Tagline for Bio section. |
| `hero_image` | String (Null) | **[New V6]** Cover image URL for Portfolio Hero. |

## 2. Table: `settings` (Key-Value Store)
Dynamic site-level settings for public-facing pages.

| Column | Type | Description |
| :--- | :--- | :--- |
| `id` | `bigint unsigned` | Primary key |
| `key` | `string(255)` | Setting key (unique) |
| `value` | `text` | Setting value |
| `created_at` | `timestamp` | Record creation |
| `updated_at` | `timestamp` |
## 6. Table: `subscribers`
Email addresses collected from newsletter subscription forms.

| Column | Type | Description |
| :--- | :--- | :--- |
| `id` | `bigint unsigned` | Primary key |
| `email` | `string(255)` | Subscriber email (unique) |
| `status` | `string` | Subscription status (default: 'active') |
| `verified_at` | `timestamp` | Email verification timestamp (nullable) |
| `created_at` | `timestamp` | Subscription date |
| `updated_at` | `timestamp` | Last update |



**Last Updated:** 2025-12-27 (Beta 0.9.80 (UI Design V6.5))

### Key Setting Keys:
- **Welcome:** `welcome_hero_title`, `welcome_hero_subtitle`, `welcome_notes_title`
- **Articles:** `articles_search_placeholder`, `articles_hero_title`
- **Footer:** `footer_copyright_text`, `footer_social_twitter`

## 3. Table: `notebooks`
Functions as Main Category. Relation: One-to-Many with Notes.

| Column | Type | Description |
| :--- | :--- | :--- |
| `id` | Integer (PK) | Unique ID. |
| `user_id` | Integer (FK)| Relation to Users. |
| `name` | String | Category Name (e.g., "SysAdmin Diary"). |
| `slug` | String | URL friendly name (Unique). |
| `description`| Text (Null) | Meta description for category page. |

## 4. Table: `notes`
Stores article content. Core Blog Logic is here.

| Column | Type | Attributes | Description |
| :--- | :--- | :--- | :--- |
| `id` | Integer | PK | Unique ID. |
| `user_id` | Integer | FK | Relation to Users (Required). `OnDelete: Cascade`. |
| `notebook_id` | Integer | FK | Relation to Notebook (Required). |
| `title` | String | Not Null | Article Title (H1). |
| `slug` | String | **Unique, Index** | URL Identifier (`/articles/great-title`). |
| `excerpt` | Text | Nullable | Summary for front page display. |
| `content` | LongText | Nullable | Main article content (HTML). |
| `cover_image` | String | Nullable | Thumbnail image path (local). |
| `status` | String | Default: 'DRAFT' | Options: `DRAFT`, `PUBLISHED`, `ARCHIVED`. |
| `published_at`| Timestamp | Nullable | Date article was published (for sorting). |
| `views` | Integer | Default: 0 | Simple reader counter. |
| `created_at` | Timestamp | | Record creation date. |
| `updated_at` | Timestamp | | Record update date. |

## 5. Business Logic Constraints
1.  **Deletion:** If a `Notebook` is deleted, all `Notes` inside it are restricted (deletion rejected) if they exist. *Decision: Restrict (For data safety).*
2.  **Public Access:** Query for public page MUST use `WHERE status = 'PUBLISHED'` filter.