FROM golang:1.19-alpine

RUN apk update &&\
    apk add git

RUN go install github.com/cosmtrek/air@latest