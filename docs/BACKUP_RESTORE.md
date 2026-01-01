# Backup & Restore Guide

Your data is the most important part of TaraNote. This guide explains how to safely backup and restore your "Digital Garden".

> **What needs backing up?**
> 1. **Database** (MySQL/SQLite): Contains all your notes, users, and settings.
> 2. **Storage** (`storage/app/public`): Contains uploaded images (avatars, covers).



## 1. Automated Backup Script (Linux)

Create a simple shell script `backup.sh` in your home directory:

```bash
#!/bin/bash

# Configuration
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_DIR="/home/youruser/backups"
APP_DIR="/var/www/taranote"
DB_USER="taranote_user"
DB_PASS="your_password"
DB_NAME="taranote_db"

# Create backup directory
mkdir -p $BACKUP_DIR

# 1. Backup Database
echo "Backing up Database..."
mysqldump -u $DB_USER -p$DB_PASS $DB_NAME > "$BACKUP_DIR/db_$TIMESTAMP.sql"

# 2. Backup Storage Images
echo "Backing up Storage..."
tar -czf "$BACKUP_DIR/storage_$TIMESTAMP.tar.gz" -C $APP_DIR storage/app/public

# 3. Compress Everything
tar -czf "$BACKUP_DIR/full_backup_$TIMESTAMP.tar.gz" "$BACKUP_DIR/db_$TIMESTAMP.sql" "$BACKUP_DIR/storage_$TIMESTAMP.tar.gz"

# Cleanup Temp Files
rm "$BACKUP_DIR/db_$TIMESTAMP.sql"
rm "$BACKUP_DIR/storage_$TIMESTAMP.tar.gz"

echo "Backup Complete: $BACKUP_DIR/full_backup_$TIMESTAMP.tar.gz"
```

**Automate with Cron (Daily at 2 AM):**
```bash
0 2 * * * /bin/bash /home/youruser/backup.sh
```



## 2. Restore Process

Disaster recovery steps.

### A. Restore Database
```bash
# If using MySQL
mysql -u taranote_user -p taranote_db < db_20251221.sql

# If using SQLite
cp database.sqlite.backup database/database.sqlite
```

### B. Restore Images
```bash
# Extract storage archive
tar -xzf storage_20251221.tar.gz -C /var/www/taranote/

# Ensure permissions are correct
chown -R www-data:www-data /var/www/taranote/storage
```

### C. Link Storage (Crucial)
After restoring, always ensure the public link exists:
```bash
php artisan storage:link
```



## 3. Manual Export (For Migrating Servers)

If you are moving to a new server:

1. **Zip everything**:
   ```bash
   zip -r taranote_migration.zip .
   ```
2. **Transfer**:
   ```bash
   scp taranote_migration.zip new_server_user@new_ip:~/
   ```
3. **Unzip & Setup**:
   ```bash
   unzip taranote_migration.zip
   composer install
   npm install
   php artisan migrate
   ```
