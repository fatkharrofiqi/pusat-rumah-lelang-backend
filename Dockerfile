FROM golang:1.21-alpine

ENV GIN_MODE=release

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
