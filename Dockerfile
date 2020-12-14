FROM golang:1.13.6

LABEL version = "1.0"

WORKDIR $GOPATH/src/github.com/ridwanakf/weebinar

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/weebinar cmd/weebinar/*.go

EXPOSE 11001

ENTRYPOINT ["/go/bin/weebinar"]