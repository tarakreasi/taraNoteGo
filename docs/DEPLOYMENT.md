# Production Deployment Protocol

**Version:** 1.0
**Target Environment:** Linux VPS (Ubuntu 22.04 LTS / Debian 12)
**Process Owner:** System Integrator

---

## 1. The Strategy: "Boring is Reliable"

TaraNote is designed to be deployed as a standard **Laravel Monolith**. We do not use Docker in production to minimize the abstraction layer complexity on small VPS instances (1GB RAM). We run on "Bare Metal" (LEMP Stack).

### Core Components
*   **Web Server:** Nginx (Reverse Proxy to PHP-FPM)
*   **Process Manager:** Systemd (Manages Queue Workers)
*   **Database:** MySQL 8.0 or SQLite (WAL Mode enabled)
*   **Cache/Session:** Redis (Optional, effectively handles high traffic)

---

## 2. Server Provisioning Scope

Before deploying the code, the environment (The "Hardware") must be ready.

**Minimum Specs:**
*   **CPU:** 1 vCPU
*   **RAM:** 1GB (Must enable 2GB Swap file for `npm run build` process)
*   **Storage:** 20GB NVMe

**Required Packages:**
```bash
sudo apt update
sudo apt install nginx php8.2-fpm php8.2-mysql php8.2-xml php8.2-curl git composer nodejs npm
```

---

## 3. Deployment Pipeline (Manual / CI)

### A. Clone & Install
```bash
cd /var/www
git clone https://github.com/tarakreasi/taraNote.git
cd taraNote
```

### B. Dependency Resolution
We optimize the autoloader for production speed.
```bash
composer install --optimize-autoloader --no-dev
npm ci
```

### C. Build Assets
Compiling the Vue 3 frontend into static assets.
```bash
npm run build
```

### D. Permission Hardening (Security)
Only specific folders should be writable by the web server.
```bash
chown -R www-data:www-data storage bootstrap/cache
chmod -R 775 storage bootstrap/cache
```

---

## 4. Nginx Configuration

A standard, secure server block configuration.

```nginx
server {
    listen 80;
    server_name taranote.com;
    root /var/www/taraNote/public;

    add_header X-Frame-Options "SAMEORIGIN";
    add_header X-Content-Type-Options "nosniff";

    index index.php;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location ~ \.php$ {
        include snippets/fastcgi-php.conf;
        fastcgi_pass unix:/var/run/php/php8.2-fpm.sock;
    }
}
```

---

## 5. Maintenance Procedures

### Updating the Application
We follow a standard update SOP to ensure zero downtime is not promised, but minimal downtime is achieved.

```bash
git pull origin main
composer install --no-dev
php artisan migrate --force
npm run build
php artisan cache:clear
php artisan config:cache
```

### Backups
**Strategy:** automated cron job.
*   **DB:** `mysqldump` to `/backups`, then synced to S3.
*   **Assets:** `/storage/app/public` synced to S3.

---

**Note:** This document assumes the operator has basic Linux proficiency. Simplicity is our primary reliability feature.
