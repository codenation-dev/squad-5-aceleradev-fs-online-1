FROM golang:latest

RUN apt -y update && apt -y install git

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD ./src /go/src/app

RUN go get -v 


