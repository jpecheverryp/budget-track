## audit: run quality control checks
.PHONY: audit
audit: 
	go fmt ./...
	go mod tidy
	go mod verify
	go vet ./...

## dev: run application
dev:
	go run ./cmd/web/

migrate-up:
	migrate -path ./migrations -database 'postgres://budget-user:budget-password@localhost:8081/budget-track-db?sslmode=disable' up
