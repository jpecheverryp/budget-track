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

## db/migrate-up Run migration file
db/migrate-up:
	migrate -path ./migrations -database 'postgres://budget-user:budget-password@localhost:8081/budget-track-db?sslmode=disable' up

## db/nuke
db/nuke:
	docker compose -f ./compose.dev.yml down -v

