FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o bitcoin-service .

EXPOSE 8080

CMD ["./bitcoin-service"]
