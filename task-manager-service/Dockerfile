FROM golang:1.25.5-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o task-manager-service ./cmd
FROM gcr.io/distroless/static

COPY --from=builder /app/task-manager-service  /task-manager-service 
ENTRYPOINT ["/task-manager-service"]