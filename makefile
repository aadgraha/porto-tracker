include .env
export

MIGRATIONS_PATH := db/migrations
MIGRATE := migrate

DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: migrate-up migrate-down migrate-down-all migrate-force migrate-version

migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-down-all:
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate-version:
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

migrate-force:
	@read -p "Force version: " v; \
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $$v
