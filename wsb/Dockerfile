FROM golang:alpine

RUN apk add --update \
   git

COPY . /go/src/github.com/qneyrat/wsb
WORKDIR /go/src/github.com/qneyrat/wsb

RUN go get ./
RUN go build

ENTRYPOINT /go/bin/wsb

EXPOSE 8007 4000
