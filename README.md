# TaraNote Go

![Go Badge](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go&logoColor=white)
![Fiber Badge](https://img.shields.io/badge/Fiber-v2-000000?style=flat&logo=gofiber&logoColor=white)
![Vue Badge](https://img.shields.io/badge/Vue.js-3-4FC08D?style=flat&logo=vue.js&logoColor=white)

> **"Built with the logic of a technician, the stability of an integrator, and the empathy of customer service."**

---

## 1. The Story Behind the Code

Hi, I'm **Tri Wantoro**.

This repository **TaraNote Go** is a small reflection of a long learning process; one that spans different roles, environments, and layers of technology.

### A Journey Through the Stack
My understanding of software didn’t start with frameworks or libraries, but grew through different perspectives:

* **Listening to Users (Customer Service)**
  Supporting users taught me how frustrating technology can be when it doesn’t behave as expected.
  → *It made me more careful about UX decisions, error handling, and communication through the interface.*

* **Learning Structure from Hardware (Electronics Technician)**
  Working with electronic components reinforced basic principles: clear inputs, predictable processes, and measurable outputs.
  → *In software, I try to apply the same discipline through modular and testable code.*

* **Respecting Production Environments (System Integrator — Ongoing)**
  Working with Linux servers and integrated systems (Milestone VMS) taught me that small mistakes can have real consequences.
  → *Because of this, I tend to be cautious about security, logging, and resource usage.*

* **Continuing to Build (Software Engineering Focus)**
  In this project, I chose **Go (Fiber)** for its raw performance and strict typing, and **Vue 3 (Inertia)** for a modern, reactive frontend.
  → *These tools let me focus on the product, not the plumbing—allowing me to write code that is easier to understand, maintain, and operate.*

---

## 2. Project Overview

**TaraNote Go** is a high-performance port of the original Laravel-based `taraNote` application. It acts as a bridge between the dynamic utility of PHP and the strict performance of Go.

*   **Goal:** Provide a lightweight, fast, and binary-distributable sanctuary for thoughts.
*   **Philosophy:** "Zen Mode" — software that disappears when you need to focus.

### Technical Highlight: The "RCA" Approach
**Challenge:** The original PHP application was flexible but resource-heavy for highly concurrent read operations.

**Solution:** Applying a **Root Cause Analysis (RCA)** mindset from my hardware days:
1.  **Isolate:** Identified that dynamic typing and per-request bootstrapping were the bottlenecks.
2.  **Trace:** Mapped every Laravel feature (Sessions, Policies, Eloquent) to a compiled equivalent.
3.  **Resolve:** Rebuilt the core in Go to achieve sub-millisecond response times.

---

## 3. Technical Architecture

| Component | Technology | Engineering Context |
| :--- | :--- | :--- |
| **Language** | Go 1.23+ | Chosen for blazing speed and low memory footprint. |
| **Framework** | [Fiber v2](https://gofiber.io/) | Express-style simplicity with Fasthttp performance. |
| **Database** | SQLite (via GORM) | Relational integrity in a zero-config, single-file format. |
| **Frontend** | Vue 3 + Inertia.js | Monolithic simplicity with SPA reactivity ("Integrated Circuit" approach). |
| **Protocol** | JSON/HTML Hybrid | Uses `X-Inertia` headers to serve data or documents dynamically. |

> **Note**: Production builds use `go-sqlite3` via CGO for maximum reliability.

---

## 4. Completed Sprints (Migration Phase)

| Sprint | Description |
| :--- | :--- |
| **Sprint 1** | Setup & Foundation (Fiber, GORM) |
| **Sprint 2** | Frontend integration (Asset copying, Vite config) |
| **Sprint 3** | Authentication (Sessions, Bcrypt, Middleware) |
| **Sprint 4** | Notebooks Domain (CRUD) |
| **Sprint 5** | Notes Domain (CRUD, Image Uploads) |
| **Sprint 6** | Public Views (Home, Article, Browser) |
| **Sprint 7** | Settings & Release Verification |
| **Sprint 8** | Unit Testing (Testify, In-Memory DB) |
| **Sprint 9** | Documentation & Narrative Alignment |

---

## 5. Development Workflow

Since I am accustomed to Linux CLI environments, here is the standard setup (SOP) to get this running locally.

### Prerequisites
*   **Go**: v1.23 or higher
*   **Node.js**: v18+ (for building frontend assets)
*   **GCC**: Required for CGO (SQLite driver)

### Quick Start Commands

```bash
# 1. Setup Environment
git clone https://github.com/tarakreasi/taraNoteGo.git
cd taraNoteGo
cp .env.example .env

# 2. Install Dependencies
go mod download
npm ci

# 3. Database Hydration (Migrate + Seed)
make migrate

# 4. Build & Run
npm run build
make run
```

---

## 6. Directory Structure Reference

```
taraNote_go/
├── cmd/
│   ├── server/       # main.go entry point
│   └── migrate/      # Database schema update utility
├── internal/
│   ├── config/       # Session, Environment configuration
│   ├── database/     # SQLite connection logic
│   ├── handlers/     # HTTP Controllers (Auth, Notes, etc.)
│   ├── models/       # GORM Data Models (structs)
│   └── utils/        # Inertia helper functions
├── resources/        # Vue/JS frontend source
├── public/           # Static assets & compiled build
└── views/            # HTML templates (Inertia root)
```

---

## 7. Docker Deployment

A multi-stage `Dockerfile` is included for lightweight deployment (Alpine Linux), reflecting my preference for clean production environments.

```bash
# Build Image
docker build -t taranote-go .

# Run Container
docker run -p 3000:3000 -v $(pwd)/database:/app/database taranote-go
```

---

## 8. Connect with Me

I am currently a System Integrator actively pivoting to a professional Fullstack Engineering role. I am ready to bring the reliability of a senior technician and the creativity of a developer to your team.

*   **LinkedIn:** [linkedin.com/in/twantoro](https://www.linkedin.com/in/twantoro)
*   **GitHub:** [github.com/tarakreasi](https://github.com/tarakreasi)
*   **Email:** ajarsinau@gmail.com

*"Ajarsinau" means "Learning to Learn". It represents my commitment to continuous evolution—from hardware to software, from technician to engineer.*

---

