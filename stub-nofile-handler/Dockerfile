FROM golang:1.25.3-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o stub-nofile-handler ./cmd
FROM gcr.io/distroless/static

COPY --from=builder /app/stub-nofile-handler  /stub-nofile-handler
ENTRYPOINT ["/stub-nofile-handler"]