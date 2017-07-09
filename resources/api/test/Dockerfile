FROM golang:1.8-stretch

WORKDIR $GOPATH/src/github.com/betalotest/api/

ADD main.go .

ADD server/ server/

RUN go get -d -v ./...

ENTRYPOINT ["go", "test", "-bench=./...", "-benchmem", "-v", "-race", "./..."]
