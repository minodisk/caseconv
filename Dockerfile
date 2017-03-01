FROM golang:1.8

WORKDIR /go/src/github.com/minodisk/caseconv
RUN go get -u \
      github.com/mattn/goveralls

COPY . .

CMD go test -v ./...
