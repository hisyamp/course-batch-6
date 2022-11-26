# syntax=docker/dockerfile:1

FROM golang:1.19.3-alpine

WORKDIR /cmd/api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/api/main.go

EXPOSE 1234

CMD [ "/cmd/api/main" ]