FROM golang:latest
RUN mkdir -p $GOPATH/src/rmq
WORKDIR $GOPATH/src/rmq/rmq-consumer
RUN go get -u github.com/golang/dep/cmd/dep
COPY . $GOPATH/src/rmq/rmq-consumer
RUN dep ensure
RUN go build -o main
ENTRYPOINT ["./main"]
