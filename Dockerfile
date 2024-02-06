FROM golang:1.16.3-alpine3.13 as builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd/server

RUN go get -d -v ./...

RUN go build -o server main.go

EXPOSE 8000

CMD ["./cmd/server/server"]