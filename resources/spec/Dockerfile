FROM golang:1.8-stretch

WORKDIR $GOPATH/src/github.com/betalotest/api/swagger/spec/

ADD swagger/spec/main.go .

ADD swagger/stub/api/swagger.yaml .

RUN go get -d -v ./...

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]
