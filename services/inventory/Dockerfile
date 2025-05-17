# Build stage with explicit platform
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

# Explicitly set build for Linux AMD64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server/

# Use distroless or alpine for runtime to have basic utilities
FROM alpine:latest

WORKDIR /
COPY --from=builder /app/main .

COPY ./config /config
COPY ./.env .
COPY ./locales /locales

# Make sure the binary is executable
RUN chmod +x /main

ENTRYPOINT ["/main"]