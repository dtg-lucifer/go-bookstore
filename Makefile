run: build
	@echo "\033[1;34m==> Running main binary...\033[0m"
	@./bin/main
.PHONY: run

build:
	@echo "\033[1;32m==> Building Go project...\033[0m"
	@go build -o bin/main cmd/**.go
.PHONY: build

db:
	@echo "\033[1;35m==> Starting DB in foreground...\033[0m"
	@sudo docker compose up db
.PHONY: db

db-d:
	@echo "\033[1;36m==> Starting DB in background...\033[0m"
	@sudo docker compose up -d
.PHONY: db-d

db-stop:
	@echo "\033[1;31m==> Stopping DB and cleaning up...\033[0m"
	@sudo docker compose down
.PHONY: db-stop
