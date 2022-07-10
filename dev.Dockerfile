FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD air