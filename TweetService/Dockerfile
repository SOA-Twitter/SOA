FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
COPY --from=build_container /app/main .
EXPOSE 9092
CMD ["./main"]