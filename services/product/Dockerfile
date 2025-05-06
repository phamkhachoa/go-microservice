FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/server/

FROM scratch

WORKDIR /
COPY --from=builder /app/main .

COPY ./config /config
COPY ./.env .

ENTRYPOINT ["/main"]