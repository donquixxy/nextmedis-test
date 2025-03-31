FROM golang:1.24.1-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -ldflags '-s -w' -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main ./

COPY .env ./

EXPOSE 7899

CMD ["./main"]