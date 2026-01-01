# TaraNote UI Design V6 - Quick Start Guide

**Goal:** Plant your Digital Garden in < 5 minutes.
 
 > "The best time to plant a tree was 20 years ago. The second best time is now."
 
 Welcome to the **TaraNote Quickstart Guide**. Whether you are a developer reviewing the code or a user setting up your own personal workspace, this guide will get you from zero to "Hello World" efficiently.



## 1. Prerequisites
- **PHP 8.2+**
- **Node.js 18+** & NPM
- **Composer**
- **SQLite** (Default) or MySQL



## The Fastest Way (Zero-to-Hero)

We provide an automated script that handles everything for you (dependencies to database).

```bash
# 1. Clone the repository
git clone https://github.com/tarakreasi/taraNote.git
cd taraNote

# 2. Run the setup script
chmod +x run.sh
./run.sh
```

The script will:
- Check for PHP, Composer, and Node.js
- Install all backend & frontend dependencies
- **Automatically handle NPM peer dependency conflicts** (Vite version compatibility)
- Configure `.env` file
- Setup and seed the database
- Build assets and start the server

---

## Manual Installation

If you prefer to run commands manually or `run.sh` doesn't work for your OS:

### Step A: Clone & Install
```bash
git clone https://github.com/tarakreasi/taraNote.git
cd taraNote

# Backend Dependencies
composer install

# Frontend Dependencies
npm install
```

### Step B: Environment Setup
```bash
cp .env.example .env
php artisan key:generate
```

### Step C: Database
1. Create a MySQL database (e.g., `taranote`).
2. Edit `.env` with your DB credentials.
3. Migrate and Seed:
```bash
php artisan migrate:fresh --seed
```

---

## Running the App

Start the development server (runs Laravel and Vite concurrently):
```bash
npm start
```

Access the app at: `http://127.0.0.1:8000`

---

## Default Credentials

**Admin Account:**
- **Email:** `ajarsinau@gmail.com`
- **Password:** `password`

---

## Common Issues

- **Port 8000 occupied**: Stop other services or change port: `php artisan serve --port=8080`
- **Database error**: Ensure MySQL server is running.
- **White screen**: Check `storage/logs/laravel.log`.
- **"Database locked"**
> Restart `php artisan serve`.
