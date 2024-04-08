FROM golang:1.21-bullseye

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy && go mod verify

COPY . . 