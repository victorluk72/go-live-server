FROM golang:1.18-alpine

WORKDIR /opt/code/

ADD ./ /opt/code/

RUN apk update && apk upgrade && apk add --no-cache git

RUN go mod download

RUN go build -o bin/webserver cmd/web/main.go

ENTRYPOINT [ "bin/webserver" ]
