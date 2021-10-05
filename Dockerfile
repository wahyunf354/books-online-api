#stage 1
FROM golang:1.16-alpine AS builder
RUN mkdir /app
ADD . /app

WORKDIR /app
RUN go clean --modcache
RUN go build -o main

#stage 2
FROM alpine:3.14
WORKDIR /root/
RUN echo '{
  "debug":true,
  "server":{"address":"${{ secrets.SERVER_ADDRESS }}"},
  "context":{"timeout":2},
  "database":{
  "prod":{
  "host": "${{ secrets.HOST_DATABASE }}",
  "port": "${{ secrets.PORT_DATABASE }}",
  "user": "${{ secrets.USER_DATABASE }}",
  "pass": "${{ secrets.PASSWORD_DATABASE }}"
  }
  }
  "jwt": {
  "secret": "${{ secrets.JWT_SECRET }}",
  "expired": 72
  }
}' >> /app/config.json
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]

