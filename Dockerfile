FROM golang:apline AS dependencies
COPY go.mod go.sum ./
