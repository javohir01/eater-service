# Docker commands for Eater Service

.PHONY: help build up down logs clean rebuild ps env-create

help:
	@echo "Eater Service - Docker Commands"
	@echo "================================"
	@echo "make build          - Build Docker image"
	@echo "make up             - Start containers"
	@echo "make down           - Stop containers"
	@echo "make down-clean     - Stop containers and remove volumes"
	@echo "make logs           - View service logs"
	@echo "make logs-db        - View database logs"
	@echo "make ps             - Show running containers"
	@echo "make rebuild        - Rebuild and restart"
	@echo "make shell          - Connect to service container"
	@echo "make db-shell       - Connect to database"
	@echo "make env-create     - Create .env from .env.example"
	@echo "make migrate        - Run database migrations"

build:
	docker-compose build

up:
	docker-compose up -d
	@echo "Service started. Access at http://localhost:8080"

down:
	docker-compose down

down-clean:
	docker-compose down -v
	@echo "Containers stopped and volumes removed"

logs:
	docker-compose logs -f eater-service

logs-db:
	docker-compose logs -f postgres

ps:
	docker-compose ps

rebuild: down build up
	@echo "Rebuilt and restarted"

shell:
	docker-compose exec eater-service /bin/sh

db-shell:
	docker-compose exec postgres psql -U $$(grep DB_USER .env | cut -d '=' -f2) -d $$(grep DB_NAME .env | cut -d '=' -f2)

env-create:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo ".env file created from .env.example"; \
	else \
		echo ".env file already exists"; \
	fi

migrate:
	docker-compose exec postgres psql -U $$(grep DB_USER .env | cut -d '=' -f2) -d $$(grep DB_NAME .env | cut -d '=' -f2) < ../src/infrastructure/migrations/000001_init.up.sql
