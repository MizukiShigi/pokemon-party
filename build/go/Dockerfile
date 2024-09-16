FROM golang:1.22
RUN apt-get update && apt-get install -y git
RUN mkdir /go/src/app
WORKDIR /go/src/app
RUN go install github.com/air-verse/air@latest
CMD ["air"]