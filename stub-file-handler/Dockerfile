FROM golang:1.25.3-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o stub-file-handler ./cmd
FROM gcr.io/distroless/static

COPY --from=builder /app/stub-file-handler  /stub-file-handler
COPY --from=builder /app/resources /resources
ENTRYPOINT ["/stub-file-handler"]