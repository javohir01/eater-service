#!/bin/bash

# Backup PostgreSQL Database
# Usage: ./backup.sh

BACKUP_DIR="../backups"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="$BACKUP_DIR/eater_db_backup_$TIMESTAMP.sql"

# Create backups directory if not exists
mkdir -p "$BACKUP_DIR"

echo "📦 Backing up database..."

# Load environment variables
set -a
source .env
set +a

# Backup database
docker compose exec -T postgres pg_dump \
  -U "$DB_USER" \
  -d "$DB_NAME" \
  --no-password > "$BACKUP_FILE"

if [ $? -eq 0 ]; then
  echo "✓ Backup created: $BACKUP_FILE"
  ls -lh "$BACKUP_FILE"
else
  echo "✗ Backup failed!"
  exit 1
fi
