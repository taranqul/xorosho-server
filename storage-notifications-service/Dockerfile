FROM golang:1.25.3-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o storage-notifications-service ./cmd
FROM gcr.io/distroless/static

COPY --from=builder /app/storage-notifications-service  /storage-notifications-service
ENTRYPOINT ["/storage-notifications-service"]