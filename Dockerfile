FROM golang:1.14.2-alpine

ADD . /echo
WORKDIR /echo

RUN go build

EXPOSE 8080

CMD ["./echo"]



