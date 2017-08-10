FROM golang:1.8-stretch

WORKDIR $GOPATH/src/github.com/restfulgopher/api/

ADD main.go .

ADD server/ server/

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8000

ENTRYPOINT ["api"]
