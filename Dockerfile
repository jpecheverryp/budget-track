# Build
FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/web ./cmd/web/

# Release
FROM gcr.io/distroless/base-debian12 AS release
WORKDIR /
COPY --from=build /bin/web /bin/web
CMD ["/bin/web"]

