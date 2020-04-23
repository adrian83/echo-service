FROM golang:1.13.7-buster

ADD . /echo
WORKDIR /echo

RUN go build

EXPOSE 8080

CMD ["./echo"]



