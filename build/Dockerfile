FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go