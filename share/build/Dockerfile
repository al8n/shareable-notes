FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

ADD . .
RUN chmod +x /app/entry-point.sh

EXPOSE 8080 50051

ENTRYPOINT ["/app/entry-point.sh"]