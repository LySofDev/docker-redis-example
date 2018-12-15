FROM golang:1.11.3-alpine3.8
RUN apk add git
RUN go get -u github.com/go-redis/redis
RUN mkdir /go/src/go-app
COPY . /go/src/go-app/
WORKDIR /go/src/go-app
