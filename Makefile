## audit: run quality control checks
.PHONY: audit
audit: 
	go fmt ./...
	go mod tidy
	go mod verify
	go vet ./...

## dev: run application
.PHONY: dev
dev:
	templ generate --watch --cmd="go run ./cmd/web/"

.PHONY: db/start
db/start:
	docker compose -f ./compose.dev.yml down && docker compose -f ./compose.dev.yml up

## db/migrate-up Run migration file
.PHONY: db/migrate-up
db/migrate-up:
	migrate -path ./migrations -database 'postgres://budget-user:budget-password@localhost:8081/budget-track-db?sslmode=disable' up

## db/nuke
.PHONY: db/nuke
db/nuke:
	docker compose -f ./compose.dev.yml down -v

## Build Executable
.PHONY: build
build: audit
	templ generate && \
	go build -o="./bin/web" ./cmd/web
