# API Reference

## Authentication
| Method | Endpoint | Auth | Description |
| :--- | :--- | :--- | :--- |
| `GET` | `/login` | Guest | Show login page |
| `POST` | `/login` | Guest | Authenticate user |
| `POST` | `/logout` | User | Destroy session |

## Public Content
| Method | Endpoint | Auth | Description |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Guest | Home / Article List |
| `GET` | `/articles/:slug` | Guest | View single article |
| `GET` | `/taranote` | Guest | 3-Column Note Browser |

## Notebooks (Admin)
| Method | Endpoint | Auth | Description |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/v1/admin/notebooks` | User | List all notebooks |
| `POST` | `/api/v1/admin/notebooks` | User | Create new notebook |
| `PUT` | `/api/v1/admin/notebooks/:id` | User | Update notebook |
| `DELETE` | `/api/v1/admin/notebooks/:id` | User | Delete notebook |

## Notes (Admin)
| Method | Endpoint | Auth | Description |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/v1/admin/notes` | User | List notes (paginated) |
| `POST` | `/api/v1/admin/notes` | User | Create new note |
| `PUT` | `/api/v1/admin/notes/:id` | User | Update note |
| `DELETE` | `/api/v1/admin/notes/:id` | User | Delete note |
| `POST` | `/api/v1/admin/upload` | User | Upload image for editor |

## Settings (Admin)
| Method | Endpoint | Auth | Description |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/v1/admin/settings` | User | Get all settings |
| `POST` | `/api/v1/admin/settings` | User | Bulk update settings |
