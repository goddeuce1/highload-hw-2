FROM golang:1.12-stretch AS lang

ADD . /go/highload-hw-2
WORKDIR /go/highload-hw-2

RUN go build

FROM ubuntu:18.04

WORKDIR /usr/src/source
COPY --from=lang /go/highload-hw-2 .

EXPOSE 8080

CMD ./highload-hw-2