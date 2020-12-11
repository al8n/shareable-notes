FROM golang:latest as builder

ENV GO111MODULE=on

RUN mkdir /app

WORKDIR /app

COPY ./gateway ./gateway

COPY ./config.yml ./config.yml

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/gateway .

EXPOSE 8080

CMD ["./gateway"]