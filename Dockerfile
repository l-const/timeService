# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /solution

RUN mkdir -p ./cmd
RUN mkdir -p ./config
RUN mkdir -p ./lib

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY config ./config
COPY lib ./lib


RUN go build -o ./server ./cmd/server/main.go 

ENTRYPOINT [ "./server" ]
