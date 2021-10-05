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

ENV SERVER_ADDRESS ${{ secrets.SERVER_ADDRESS }}
ENV HOST_DATABASE ${{ secrets.HOST_DATABASE }}
ENV PORT_DATABASE ${{ secrets.PORT_DATABASE }}
ENV USER_DATABASE ${{ secrets.USER_DATABASE }}
ENV PASSWORD_DATABASE ${{ secrets.PASSWORD_DATABASE }}
ENV JWT_SECRET ${{ secrets.JWT_SECRET }}

RUN echo '{"debug":true,"server":{"address":"$SERVER_ADDRESS"},"context":{"timeout":2},"database":{"prod":{"host": "$HOST_DATABASE","port": "$PORT_DATABASE","user": "$USER_DATABASE","pass": "$PASSWORD_DATABASE"}}, "jwt": {"secret": "$JWT_SECRET","expired": 72}}' >> /app/config.json

COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]

