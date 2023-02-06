FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine
WORKDIR /root/

COPY --from=builder /app/main .
COPY ./certificates ./certificates
COPY ./data/policy.csv ./data/policy.csv
COPY ./data/rbac_model.conf ./data/rbac_model.conf
EXPOSE 8081
CMD ["./main"]