#stage 1
FROM golang:1.16-alpine AS builder
RUN mkdir /app
ADD . /app

WORKDIR /app
RUN ls
RUN go clean --modcache
RUN go build -o main

#stage 2
FROM alpine:3.14
WORKDIR /root/

RUN pwd
RUN ls
COPY --from=builder /app/config.json .
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]

