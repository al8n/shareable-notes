FROM golang:latest as builder

ENV GO111MODULE=on

RUN mkdir /app

WORKDIR /app

COPY ./shared ./shared

COPY ./config.yml ./config.yml

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/shared .

EXPOSE 8080 50051

CMD ["./shared"]

