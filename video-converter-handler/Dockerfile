FROM golang:1.25.5-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o video-converter-handler ./cmd

FROM alpine:3.21

RUN apk add --no-cache ffmpeg ca-certificates

COPY --from=builder /app/video-converter-handler /video-converter-handler
COPY --from=builder /app/resources /resources
WORKDIR /

ENTRYPOINT ["/video-converter-handler"]
