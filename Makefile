
.phony run:
run:
	@echo "Running..."
	go run ./cmd/main.go

.phony up:
up:
	@echo "Starting docker-compose..."
	docker compose up -d

.phony down:
down:
	@echo "Stopping docker-compose..."
	docker compose down

# this is for local development only
.phony migrate:
migrate:
	@echo "Migrating..."
	pgroll init --postgres-url "postgres://postgres:postgres@localhost:54321/porter_management_db?sslmode=disable"
	pgroll --postgres-url "postgres://postgres:postgres@localhost:54321/porter_management_db?sslmode=disable" start ./migration/01_initial_migration.json --complete
	pgroll --postgres-url "postgres://postgres:postgres@localhost:54321/porter_management_db?sslmode=disable" start ./migration/02_create_table_equipment.json --complete