# Nginx Configuration Guide

This guide provides the production Nginx configuration for hosting TaraNote on a Linux VPS (Ubuntu/Debian).



## 1. Site Configuration

Create a new file at `/etc/nginx/sites-available/taranote`:

```nginx
server {
    listen 80;
    server_name your-domain.com www.your-domain.com;
    root /var/www/taranote/public;

    add_header X-Frame-Options "SAMEORIGIN";
    add_header X-Content-Type-Options "nosniff";
    add_header X-XSS-Protection "1; mode=block";

    index index.php;

    charset utf-8;

    # Handle Frontend Routes (Inertia.js/Vue)
    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    # Handle Favicon/Robots
    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    # Handle PHP Scripts
    location ~ \.php$ {
        fastcgi_pass unix:/var/run/php/php8.2-fpm.sock; # Adjust PHP version if needed
        fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
        include fastcgi_params;
    }

    # Deny Access to Hidden Files (.env, .git)
    location ~ /\.(?!well-known).* {
        deny all;
    }
}
```

## 2. Enable Site

Link the configuration to `sites-enabled`:

```bash
sudo ln -s /etc/nginx/sites-available/taranote /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

## 3. SSL Setup (Certbot)

Secure your site with a free Let's Encrypt certificate:

```bash
# Install Certbot (if not installed)
sudo apt install certbot python3-certbot-nginx

# Auto-configure SSL
sudo certbot --nginx -d your-domain.com -d www.your-domain.com
```

Certbot will automatically update your Nginx config to force HTTPS.



## 4. Optimization Tips

### Gzip Compression
Ensure `/etc/nginx/nginx.conf` has gzip enabled for JSON/API responses:

```nginx
gzip on;
gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
```

### File Upload Limits
If you plan to upload large Hero images, increase the body size limit:

```nginx
client_max_body_size 10M;
```
