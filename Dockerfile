FROM golang:1.22-alpine AS builder
WORKDIR /build

RUN apk add --no-cache git

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine:latest AS server

WORKDIR /balance_api

COPY --from=builder /build/app .

CMD ["./app"]
