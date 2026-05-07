# Docker commands for Eater Service

.PHONY: help build up down logs clean rebuild ps env-create migrate backup restore db-shell

help:
	@echo "Eater Service - Docker Commands"
	@echo "================================"
	@echo "make build          - Build Docker image"
	@echo "make up             - Start containers (DATA SAVED)"
	@echo "make down           - Stop containers (DATA SAVED)"
	@echo "make down-clean     - Stop containers and remove volumes (DATA LOST!)"
	@echo "make ps             - Show running containers"
	@echo "make logs           - View service logs"
	@echo "make logs-db        - View database logs"
	@echo "make shell          - Connect to service container"
	@echo "make db-shell       - Connect to database shell"
	@echo "make env-create     - Create .env from .env.example"
	@echo "make backup         - Backup database to SQL file"
	@echo "make restore        - Restore database from backup"
	@echo "make rebuild        - Rebuild and restart"

build:
	docker compose build

up:
	docker compose up -d
	@echo "Service started. Access at http://localhost:8080"

down:
	docker compose down

down-clean:
	docker compose down -v
	@echo "Containers stopped and volumes removed"

logs:
	docker compose logs -f eater-service

logs-db:
	docker compose logs -f postgres

ps:
	docker compose ps

rebuild: down build up
	@echo "Rebuilt and restarted"

shell:
	docker compose exec eater-service /bin/sh

db-shell:
	@. .env && docker compose exec postgres psql -U $$DB_USER -d $$DB_NAME

env-create:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo ".env file created from .env.example"; \
	else \
		echo ".env file already exists"; \
	fi

backup:
	@bash backup.sh

restore:
	@if [ -z "$(BACKUP)" ]; then \
		echo "Usage: make restore BACKUP=backups/eater_db_backup_YYYYMMDD_HHMMSS.sql"; \
		exit 1; \
	fi
	@echo "⚠️  WARNING: This will overwrite current database!"
	@echo "Press Ctrl+C to cancel, Enter to continue..."
	@read; \
	. .env && \
	docker compose exec -T postgres psql -U $$DB_USER -d $$DB_NAME < $(BACKUP) && \
	echo "✓ Database restored from $(BACKUP)"

migrate:
	docker-compose exec postgres psql -U $$(grep DB_USER .env | cut -d '=' -f2) -d $$(grep DB_NAME .env | cut -d '=' -f2) < ../src/infrastructure/migrations/000001_init.up.sql
