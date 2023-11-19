db_dev:
	docker compose -f docker-compose.dev.yaml -p genshinacademycore up -d postgres_db
d_purge:
	docker compose down -v --remove-orphans
	docker compose rm -vsf