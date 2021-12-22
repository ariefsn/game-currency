FROM golang:alpine

RUN apk add build-base

RUN go version

RUN mkdir /app

WORKDIR /app

COPY go.mod go.mod

COPY go.sum go.sum

COPY global global

COPY services/currency-service services/currency-service

COPY services/currency-service/.env.example .env

RUN go mod tidy

RUN go build -o binary /app/services/currency-service

CMD [ "/app/binary" ]
