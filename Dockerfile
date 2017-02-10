FROM ruby:2.1.10

RUN apt-get update && apt-get install -y golang

RUN mkdir -p /go/src/todo
WORKDIR /go
ENV GOPATH=/go
ADD . /go/src/todo
RUN go get todo
RUN cp /go/bin/todo /usr/bin/todo-cli

CMD ["/bin/sh"]
