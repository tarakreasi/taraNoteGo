# TaraNote Go 🚀

**TaraNote Go** is a high-performance port of the TaraNote application, rewritten in **Go (Golang)** using the **Fiber** web framework and **GORM**. It retains the original **Vue 3** + **Inertia.js** frontend, providing a seamless modern user experience with a blazing fast backend.

## 🛠 Tech Stack

- **Backend**: Go 1.23+, Fiber v2 (Web Framework), GORM (ORM), SQLite (Database).
- **Frontend**: Vue 3, Inertia.js, TailwindCSS, Vite.
- **Architecture**: Monolithic (Go serving Inertia assets).

## 🚀 Getting Started

### Prerequisites

- **Go**: v1.23 or higher.
- **Node.js**: v18+ (for building frontend assets).
- **GCC**: Required for CGO (SQLite driver).

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/tarakreasi/taraNoteGo.git
    cd taraNoteGo
    ```

2.  **Setup Environment**
    Copy the example configuration:
    ```bash
    cp .env.example .env
    ```

3.  **Install Dependencies**
    ```bash
    # Backend
    go mod download

    # Frontend
    npm install
    ```

4.  **Database Migration**
    Initialize the SQLite database and tables:
    ```bash
    make migrate
    # OR
    go run cmd/migrate/main.go
    ```

### 🏃‍♂️ Running Locally

1.  **Build Frontend** (Required for production, or run `npm run dev` for hot reload if configured):
    ```bash
    npm run build
    ```

2.  **Start Backend Server**:
    ```bash
    make run
    # OR
    go run cmd/server/main.go
    ```
    The server will start at `http://127.0.0.1:3000`.

## 📦 Building for Production

You can build optimized binaries using the `Makefile`:

```bash
make build
```
This produces:
- `bin/server`: Main application binary.
- `bin/migrate`: Database migration utility.

## 🐳 Docker Deployment

A multi-stage `Dockerfile` is included for lightweight deployment (Alpine Linux).

```bash
# Build Image
docker build -t taranote-go .

# Run Container
docker run -p 3000:3000 -v $(pwd)/database:/app/database taranote-go
```

## 📂 Project Structure

```
├── cmd/
│   ├── server/       # Main web server entry point
│   └── migrate/      # Database migration utility
├── internal/
│   ├── config/       # Configuration (Session, Env)
│   ├── database/     # DB Connection logic
│   ├── handlers/     # HTTP Controllers (Auth, Notes, etc.)
│   ├── middleware/   # Fiber Middleware (Auth Guard)
│   ├── models/       # GORM Data Models
│   └── utils/        # Helpers (Inertia JSON generator)
├── resources/        # Frontend source (Vue, CSS)
├── public/           # Static assets & compiled build
└── views/            # HTML Templates (Inertia root)
```

## 📝 License

This project is open-source and available under the [MIT License](LICENSE).
