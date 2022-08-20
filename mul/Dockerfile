# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app/app ./app
ENTRYPOINT ["./app"]