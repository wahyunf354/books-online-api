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

#ENV SERVER_ADDRESS $GITHUB_SERVER_ADDRESS
#ENV HOST_DATABASE $GITHUB_HOST_DATABASE
#ENV PORT_DATABASE $GITHUB_PORT_DATABASE
#ENV USER_DATABASE $GITHUB_USER_DATABASE
#ENV PASSWORD_DATABASE $GITHUB_PASSWORD_DATABASE
#ENV JWT_SECRET $GITHUB_JWT_SECRET

RUN echo '{"debug":true,"server":{"address":":8080"},"context":{"timeout":2},"database":{"prod":{"host": "databases-books-online.cl8sopwirx70.us-east-2.rds.amazonaws.com","port": "3306","user": "books_online","pass": "books_online", "name": "books_online"}}, "jwt": {"secret": "KambingGUling","expired": 72}}' >> config.json

#COPY --from=builder /app/config.json .
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]

