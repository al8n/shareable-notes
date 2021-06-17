FROM golang:buster AS builder
ENV GO111MODULE=on
WORKDIR /app
ADD . .
RUN go build  /app/apigateway/gateway.go \
    && go build /app/share-svc/share.go

FROM debian:buster-slim
COPY --from=builder /app/gateway /usr/bin/
COPY --from=builder /app/share /usr/bin/
COPY --from=builder /app/conf /conf
WORKDIR /conf

