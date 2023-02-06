FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 9092
CMD ["./main"]