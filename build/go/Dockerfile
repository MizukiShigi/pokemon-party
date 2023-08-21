FROM golang:1.20
RUN apt-get update && apt-get install -y git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app