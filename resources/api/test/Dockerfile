FROM golang:1.8-stretch

WORKDIR $GOPATH/src/github.com/restfulgopher/api/

ADD main.go .

ADD server/ server/

RUN go get -d -v ./...

ENTRYPOINT ["go", "test", "-bench=./...", "-benchmem", "-v", "-race", "./..."]
