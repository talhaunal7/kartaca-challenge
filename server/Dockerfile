FROM golang:1.18.3

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

EXPOSE 8080

RUN go mod tidy