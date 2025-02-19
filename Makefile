## audit: run quality control checks
.PHONY: audit
audit: 
	go fmt ./...
	go mod tidy
	go mod verify
	go vet ./...

## dev: run application
dev:
	docker compose -f ./compose.dev.yml down && docker compose -f ./compose.dev.yml up --build

migrate-up:
	migrate -path ./migrations -database 'postgres://budget-user:budget-password@localhost:8081/budget-track-db?sslmode=disable' up
