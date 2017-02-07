FROM golang:1.7.5

ADD . /go/src/github.com/tillkahlbrock/todo/

WORKDIR /go

RUN go get github.com/tillkahlbrock/todo

CMD ["/go/bin/todo"]
