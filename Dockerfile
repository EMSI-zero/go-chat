# syntax=docker/dockerfile:1

FROM golang:1.21 AS builder



WORKDIR $GOPATH/src/gitlab.com/funtory-webshop/webshop-api-service

COPY go.mod go.sum ./


RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g ./cmd/webshop/main.go -o ./docs

RUN go mod tidy


RUN CGO_ENABLED=0 GOOS=linux go build -o build ./cmd/webshop



FROM alpine:latest

ENV TZ Asia/Tehran

COPY --from=builder /go/src/gitlab.com/funtory-webshop/webshop-api-service/build /etc/webshop/server

WORKDIR /etc/webshop
CMD ["./server"]
