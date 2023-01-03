# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./out/go-sample .

## Deploy
FROM alpine:3.16.2

COPY --from=build-base /app/out/go-sample  /app/go-sample

CMD  ["/app/go-sample"]