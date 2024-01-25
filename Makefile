
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
