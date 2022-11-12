FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine
COPY --from=build_container /app/main .
EXPOSE 8081
CMD ["./main"]