ENV := local
COMPOSE_FILE := -f docker-compose.$(ENV).yaml
PROJECT_NAME := genshinacademycore-$(ENV)

docker:
	docker compose $(COMPOSE_FILE) -p $(PROJECT_NAME)

up:
	docker compose $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d

down:
	docker compose $(COMPOSE_FILE) -p $(PROJECT_NAME) down

rebuild:
	docker compose $(COMPOSE_FILE) -p $(PROJECT_NAME) down -v --remove-orphans
	docker compose $(COMPOSE_FILE) -p $(PROJECT_NAME) rm -vsf
	docker compose $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d --build

db_dev:
	docker compose -f docker-compose.dev.yaml -p genshinacademycore up -d postgres_db
d_purge:
	docker compose down -v --remove-orphans
	docker compose rm -vsf