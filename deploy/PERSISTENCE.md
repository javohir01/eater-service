# Eater Service - Docker Deployment Guide

## 📊 Database Persistence

### Data Storage 💾

Your database data is **safely stored in Docker volumes**:

```bash
# View volumes
docker volume ls | grep eater

# Inspect volume location
docker volume inspect deploy_postgres_data
```

### Important Commands ⚠️

```bash
# SAFE - Data is preserved
docker compose down
docker compose up -d

# DANGEROUS - Data is deleted!
docker compose down -v    # ❌ DO NOT USE unless you want to delete data!
make down-clean           # ❌ Same as above!
```

## 🔄 Backup & Restore

### Automatic Backup

```bash
# Create backup
make backup

# Backup is saved to: backups/eater_db_backup_YYYYMMDD_HHMMSS.sql
ls -lh backups/
```

### Restore from Backup

```bash
# Restore database (you'll see a confirmation prompt)
make restore BACKUP=backups/eater_db_backup_20260506_123456.sql
```

## 🐚 Database Access

Access the database shell:

```bash
make db-shell

# Inside PostgreSQL shell:
\dt          # List tables
\d table     # Describe table
SELECT COUNT(*) FROM users;
\q           # Exit
```

Or directly:

```bash
cd deploy/
docker compose exec postgres psql -U $DB_USER -d $DB_NAME
```

## 📁 Volume Locations

Production volumes are stored on your system at:

```bash
# Linux/macOS
/var/lib/docker/volumes/

# Example
/var/lib/docker/volumes/deploy_postgres_data/_data/
```

## 🚀 Production Deployment

For production, use `docker-compose.prod.yml`:

```bash
# Start production services
docker compose -f docker-compose.prod.yml up -d

# View logs
docker compose -f docker-compose.prod.yml logs -f

# Regular backup schedule
# Add to crontab:
# 0 2 * * * cd /path/to/deploy && ./backup.sh
```

## 📋 Quick Reference

| Command | Effect | Data |
|---------|--------|------|
| `make up` | Start containers | ✓ Preserved |
| `make down` | Stop containers | ✓ Preserved |
| `make down-clean` | Stop & remove volumes | ✗ Deleted |
| `make backup` | Create SQL backup | ✓ Backed up |
| `make restore` | Restore from backup | ✓ Restored |

## ⚙️ Configuration

Database settings in `.env`:

```bash
DB_USER=eater
DB_PASSWORD=eater123
DB_NAME=eater_db
```

**Important:** Change `DB_PASSWORD` in production!

## 🆘 Emergency Recovery

If volumes are somehow corrupted:

```bash
# View backup files
ls -lh backups/

# List recent backups by date
ls -lh backups/ | sort -k6,7

# Restore from latest backup
make restore BACKUP=$(ls -t backups/*.sql | head -1)
```

## 📝 Notes

- Backups are stored in `../backups/` (project root level)
- Backup files are NOT version controlled (see `.gitignore`)
- Keep backups in a secure location
- For production, use cloud storage (S3, backup service, etc.)
- Weekly+ backup schedule recommended

---

**Need help?** Run `make help` for all available commands.
