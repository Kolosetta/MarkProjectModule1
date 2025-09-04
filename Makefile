# Загружаем переменные из .env
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: migrate-up migrate-down migrate-status

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL_DOCKER)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL_DOCKER)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL_DOCKER)" status