FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /app
RUN apk update && apk upgrade && \
    apk add bash git openssh gcc libc-dev
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build  /app/cmd/server

FROM alpine:3.14


ENV OTEL_SERVICE_NAME=crypto-watch
COPY --from=builder /app/server /app/server

WORKDIR /app
CMD ["/app/server"]
