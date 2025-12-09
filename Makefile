.PHONY: help build up down restart logs clean db-shell api-shell test

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build all services
	docker-compose build

up: ## Start all services
	docker-compose up -d

down: ## Stop all services
	docker-compose down

restart: down up ## Restart all services

logs: ## View logs from all services
	docker-compose logs -f

logs-api: ## View API logs only
	docker-compose logs -f api

logs-db: ## View PostgreSQL logs only
	docker-compose logs -f postgres

clean: ## Stop services and remove volumes
	docker-compose down -v

rebuild: ## Rebuild and restart services
	docker-compose up -d --build

db-shell: ## Access PostgreSQL shell
	docker-compose exec postgres psql -U kotonilo -d kotonilo_db

api-shell: ## Access API container shell
	docker-compose exec api sh

ps: ## Show running containers
	docker-compose ps

health: ## Check health of services
	@echo "Checking API health..."
	@curl -f http://localhost:8080/health || echo "API is not responding"
	@echo "\nChecking database connection..."
	@docker-compose exec -T postgres pg_isready -U kotonilo -d kotonilo_db
