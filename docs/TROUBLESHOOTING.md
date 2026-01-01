# Troubleshooting Guide

**TaraNote Technical Troubleshooting**

This guide helps developers and system administrators diagnose and fix common issues.



## 1. Installation & Setup Issues

### Problem: `composer install` fails

**Symptoms**:
```
Your requirements could not be resolved to an installable set of packages.
```

**Solutions**:
1. **Check PHP version**: TaraNote requires PHP 8.2+
   ```bash
   php -v
   ```
2. **Update Composer**:
   ```bash
   composer self-update
   ```
3. **Clear Composer cache**:
   ```bash
   composer clear-cache
   composer install
   ```
4. **Check memory limit**: Increase PHP memory limit
   ```bash
   php -d memory_limit=-1 /usr/local/bin/composer install
   ```



### Problem: `npm install` fails

**Symptoms**:
```
ERR! code ERESOLVE
ERR! ERESOLVE could not resolve
```

**Solutions**:
1. **Use correct Node version**: TaraNote requires Node 18+
   ```bash
   node -v
   nvm install 18
   nvm use 18
   ```
2. **Clear npm cache**:
   ```bash
   npm cache clean --force
   rm -rf node_modules package-lock.json
   npm install
   ```
3. **Use legacy peer deps**:
   ```bash
   npm install --legacy-peer-deps
   ```



### Problem: `.env` file errors

**Symptoms**:
```
RuntimeException: No application encryption key has been specified.
```

**Solutions**:
1. **Copy `.env.example`**:
   ```bash
   cp .env.example .env
   ```
2. **Generate app key**:
   ```bash
   php artisan key:generate
   ```
3. **Set database credentials** in `.env`:
   ```
   DB_CONNECTION=mysql
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_DATABASE=taranote
   DB_USERNAME=root
   DB_PASSWORD=yourpassword
   ```



### Problem: Database connection fails

**Symptoms**:
```
SQLSTATE[HY000] [2002] Connection refused
```

**Solutions**:
1. **Check if database service is running**:
   ```bash
   # MySQL
   sudo systemctl status mysql
   sudo systemctl start mysql

   # PostgreSQL
   sudo systemctl status postgresql
   sudo systemctl start postgresql
   ```
2. **Verify credentials** in `.env`
3. **Create database**:
   ```bash
   mysql -u root -p
   CREATE DATABASE taranote;
   exit;
   ```
4. **Check port**: Default is 3306 (MySQL) or 5432 (Postgres)



## 2. Migration & Database Issues

### Problem: `php artisan migrate` fails

**Symptoms**:
```
SQLSTATE[42S01]: Base table or view already exists
```

**Solutions**:
1. **Fresh migration** (deletes all data):
   ```bash
   php artisan migrate:fresh
   ```
2. **Rollback and re-migrate**:
   ```bash
   php artisan migrate:rollback
   php artisan migrate
   ```
3. **Specific migration**:
   ```bash
   php artisan migrate --path=/database/migrations/xxxx_create_users_table.php
   ```



### Problem: Foreign key constraint errors

**Symptoms**:
```
SQLSTATE[23000]: Integrity constraint violation
```

**Solutions**:
1. **Check migration order**: Foreign keys must reference existing tables
2. **Disable foreign key checks** (temporary):
   ```php
   Schema::disableForeignKeyConstraints();
   // ... drop tables
   Schema::enableForeignKeyConstraints();
   ```
3. **Use `onDelete('cascade')`**:
   ```php
   $table->foreignId('user_id')->constrained()->onDelete('cascade');
   ```



## 3. UI & Vue Component Issues

### Problem: Clicking notes doesn't display content (blank panel)

**Symptoms**:
- In TaraNote view (`/taraNote`): Clicking a note in the middle column doesn't show the article in the right panel
- In Dashboard (`/dashboard`): Clicking a note doesn't load the editor
- Browser console shows Vue warnings: `[Vue warn]: Unhandled error during execution of render function`

**Root Cause**: Vue render errors caused by references to archived/non-existent routes in component templates.

**Solutions**:
1. **Check for broken route references**:
   ```bash
   # Search for archived routes
   grep -r "route('portfolio.show'" resources/js/
   grep -r "route('articles.show'" resources/js/
   ```

2. **Fix TaraNote.vue** (if portfolio route is archived):
   ```vue
   <!-- BEFORE (broken) -->
   <Link :href="route('portfolio.show', user.username)">
       {{ user.name }}
   </Link>
   
   <!-- AFTER (fixed) -->
   <span class="text-sm font-bold">
       {{ user.name }}
   </span>
   ```

3. **Fix EditorSection.vue** (if articles route is archived):
   ```vue
   <!-- Comment out or remove the broken preview link -->
   <!--
   <a :href="route('articles.show', note.slug)">
       <span>visibility</span>
   </a>
   -->
   ```

4. **Clear browser cache and verify**:
   ```bash
   # In browser: Ctrl+Shift+R (hard refresh)
   # Or clear Vite cache
   rm -rf node_modules/.vite
   npm run dev
   ```

**Prevention**: After archiving routes, always search for references in Vue components:
```bash
# Find all route() calls
grep -rn "route('" resources/js/ | grep -v node_modules
```



## 4. Frontend & Build Issues

### Problem: `npm run build` fails

**Symptoms**:
```
ERROR in ./resources/js/app.js
Module not found
```

**Solutions**:
1. **Check imports**: Ensure all imports are correct
2. **Clear Vite cache**:
   ```bash
   rm -rf node_modules/.vite
   npm run build
   ```
3. **Increase Node memory**:
   ```bash
   NODE_OPTIONS=--max_old_space_size=4096 npm run build
   ```



### Problem: Vite dev server not hot-reloading

**Symptoms**: Changes to Vue files don't reflect in browser without manual refresh.

**Solutions**:
1. **Check `vite.config.js` HMR settings**:
   ```js
   server: {
       hmr: {
           host: 'localhost',
       },
   }
   ```
2. **Restart dev server**:
   ```bash
   npm run dev
   ```
3. **Check firewall**: Ensure port 5173 is open



### Problem: Assets (images) not loading

**Symptoms**: 404 errors for `/storage/uploads/image.jpg`

**Solutions**:
1. **Create storage symlink**:
   ```bash
   php artisan storage:link
   ```
2. **Check file permissions**:
   ```bash
   chmod -R 755 storage
   chown -R www-data:www-data storage
   ```
3. **Verify `.htaccess`** in `public/` allows symlinks:
   ```apache
   Options +FollowSymLinks
   ```



## 4. Authentication Issues

### Problem: Login redirects to login page (infinite loop)

**Symptoms**: After entering credentials, you're redirected back to login.

**Solutions**:
1. **Clear session**:
   ```bash
   php artisan session:clear
   ```
2. **Check `SESSION_DRIVER`** in `.env`:
   ```
   SESSION_DRIVER=file
   ```
3. **Check session permissions**:
   ```bash
   chmod -R 755 storage/framework/sessions
   ```



### Problem: "CSRF token mismatch" error

**Symptoms**:
```
419 | Page Expired
```

**Solutions**:
1. **Hard refresh**: `Ctrl+Shift+R` or `Cmd+Shift+R`
2. **Check session cookie domain** in `config/session.php`:
   ```php
   'domain' => env('SESSION_DOMAIN'),
   ```
3. **Clear browser cookies**
4. **Restart Laravel**:
   ```bash
   php artisan config:clear
   php artisan cache:clear
   ```



### Problem: Password reset emails not sending

**Symptoms**: "Password reset link sent" but no email received.

**Solutions**:
1. **Check mail configuration** in `.env`:
   ```
   MAIL_MAILER=smtp
   MAIL_HOST=smtp.mailtrap.io
   MAIL_PORT=2525
   MAIL_USERNAME=your_username
   MAIL_PASSWORD=your_password
   ```
2. **Test mail**:
   ```bash
   php artisan tinker
   Mail::raw('Test', function($msg) { $msg->to('test@example.com')->subject('Test'); });
   ```
3. **Check spam folder**
4. **Use queue** (if configured):
   ```bash
   php artisan queue:work
   ```



## 5. Permission & File Upload Issues

### Problem: File upload fails

**Symptoms**:
```
The file "image.jpg" was not uploaded due to an unknown error.
```

**Solutions**:
1. **Check file size**: Max upload is 5MB by default
2. **Check PHP upload limits** in `php.ini`:
   ```ini
   upload_max_filesize = 10M
   post_max_size = 12M
   ```
3. **Check storage permissions**:
   ```bash
   chmod -R 775 storage/app
   chown -R www-data:www-data storage
   ```
4. **Check disk space**:
   ```bash
   df -h
   ```



### Problem: "Permission denied" errors

**Symptoms**:
```
file_put_contents(/path/to/storage/logs/laravel.log): failed to open stream: Permission denied
```

**Solutions**:
1. **Fix ownership**:
   ```bash
   sudo chown -R www-data:www-data storage bootstrap/cache
   ```
2. **Fix permissions**:
   ```bash
   chmod -R 755 storage bootstrap/cache
   ```
3. **SELinux** (if enabled):
   ```bash
   sudo chcon -R -t httpd_sys_rw_content_t storage
   ```



## 6. Performance Issues

### Problem: Slow page loads

**Symptoms**: Pages take >5 seconds to load.

**Solutions**:
1. **Enable caching**:
   ```bash
   php artisan config:cache
   php artisan route:cache
   php artisan view:cache
   ```
2. **Optimize Composer**:
   ```bash
   composer install --optimize-autoloader --no-dev
   ```
3. **Use production build**:
   ```bash
   npm run build
   ```
4. **Enable OPcache** in `php.ini`:
   ```ini
   opcache.enable=1
   opcache.memory_consumption=128
   ```



### Problem: Database queries are slow

**Symptoms**: Laravel Debugbar shows queries taking >1 second.

**Solutions**:
1. **Add indexes**:
   ```php
   $table->index('user_id');
   $table->index('slug');
   ```
2. **Eager load relationships**:
   ```php
   Note::with('user', 'notebook')->get();
   ```
3. **Use query logging** to find N+1 queries:
   ```php
   DB::enableQueryLog();
   // ... your code
   dd(DB::getQueryLog());
   ```



### Problem: High memory usage

**Symptoms**: PHP crashes with "Allowed memory size exhausted".

**Solutions**:
1. **Increase PHP memory limit** in `php.ini`:
   ```ini
   memory_limit = 512M
   ```
2. **Use chunking** for large datasets:
   ```php
   Note::chunk(200, function ($notes) {
       // Process notes
   });
   ```
3. **Clear unnecessary data**:
   ```bash
   php artisan cache:clear
   php artisan view:clear
   ```



## 7. Deployment Issues

### Problem: 500 Internal Server Error

**Symptoms**: Blank page or generic 500 error.

**Solutions**:
1. **Check Laravel logs**:
   ```bash
   tail -f storage/logs/laravel.log
   ```
2. **Enable debug mode** (temporarily):
   ```
   APP_DEBUG=true
   ```
3. **Check Apache/Nginx error logs**:
   ```bash
   # Apache
   sudo tail -f /var/log/apache2/error.log

   # Nginx
   sudo tail -f /var/log/nginx/error.log
   ```
4. **Check file permissions** (see Section 5)



### Problem: CSS/JS not loading in production

**Symptoms**: Site looks broken, no styles.

**Solutions**:
1. **Run production build**:
   ```bash
   npm run build
   ```
2. **Check `APP_URL` in `.env`**:
   ```
   APP_URL=https://yourdomain.com
   ```
3. **Verify asset paths** in `vite.config.js`:
   ```js
   build: {
       outDir: 'public/build',
   }
   ```



### Problem: Artisan commands hanging

**Symptoms**: `php artisan` commands freeze indefinitely.

**Solutions**:
1. **Check for infinite loops** in ServiceProviders or middleware
2. **Increase PHP timeout** in `php.ini`:
   ```ini
   max_execution_time = 60
   ```
3. **Run in verbose mode**:
   ```bash
   php artisan migrate -vvv
   ```



## 8. Testing Issues

### Problem: Tests fail with database errors

**Symptoms**:
```
SQLSTATE[HY000]: General error: 1 no such table: users
```

**Solutions**:
1. **Use in-memory SQLite** for tests:
   ```php
   // phpunit.xml
   <env name="DB_CONNECTION" value="sqlite"/>
   <env name="DB_DATABASE" value=":memory:"/>
   ```
2. **Run migrations in tests**:
   ```php
   use RefreshDatabase;
   ```
3. **Clear test cache**:
   ```bash
   php artisan test --clear-cache
   ```



### Problem: Feature tests fail with "Session store not set"

**Solutions**:
1. **Set session driver** in `phpunit.xml`:
   ```xml
   <env name="SESSION_DRIVER" value="array"/>
   ```
2. **Use `actingAs()`** for authenticated tests:
   ```php
   $this->actingAs($user)->get('/taranote');
   ```



## 9. Common Error Messages

### "Class not found"

**Solutions**:
1. **Dump autoload**:
   ```bash
   composer dump-autoload
   ```
2. **Check namespace** and file location
3. **Clear config cache**:
   ```bash
   php artisan config:clear
   ```



### "Route not found"

**Solutions**:
1. **Clear route cache**:
   ```bash
   php artisan route:clear
   ```
2. **Check route name** in `routes/web.php`
3. **Run route list**:
   ```bash
   php artisan route:list
   ```



### "Too many redirects"

**Solutions**:
1. **Check middleware** for infinite redirect loops
2. **Clear browser cookies**
3. **Check `.htaccess` for HTTPS redirects**



## 10. Getting Help

### Before Asking for Help

1. **Check logs**: `storage/logs/laravel.log`
2. **Search documentation**: [tarakreasi.com/docs](https://tarakreasi.com/docs)
3. **Search GitHub Issues**: [github.com/tarakreasi/taranote/issues](#)

### How to Report an Issue

Include:
1. **Environment**: OS, PHP version, Laravel version
2. **Steps to reproduce**: Exact commands run
3. **Error message**: Full error (not just "it doesn't work")
4. **Logs**: Relevant lines from `storage/logs/laravel.log`
5. **What you tried**: List troubleshooting steps attempted

### Contact Support

- **GitHub Issues**: [github.com/tarakreasi/taranote/issues](#)
- **Email**: ajarsinau@gmail.com
- **Discord** (coming soon)



## 11. Useful Commands

### Clear All Caches
```bash
php artisan optimize:clear
```

### Re-seed Database
```bash
php artisan migrate:fresh --seed
```

### Check Laravel Version
```bash
php artisan --version
```

### List All Routes
```bash
php artisan route:list
```

### Run Queue Worker
```bash
php artisan queue:work --tries=3
```

### Check Disk Space
```bash
df -h
```

### Monitor Logs (Live)
```bash
tail -f storage/logs/laravel.log
```



**Last Updated:** December 27, 2025 (Beta 0.9.80 (UI Design V6.5))
