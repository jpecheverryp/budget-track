# Build Stage
FROM golang:1.23 AS build

WORKDIR /src

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /bin/web ./cmd/web/

# Smaller Image for Deployment
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=build /bin/web /bin/web

CMD ["/bin/web"]

