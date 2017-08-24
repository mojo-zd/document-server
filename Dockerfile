FROM golang:1.8 as build-step
MAINTAINER excessivespeed@126.com

ADD . /go/src/github.com/document-server
WORKDIR /go/src/github.com/document-server

RUN GOOS=linux GOARCH=amd64 go build -v -o /go/bin/document-server


FROM alpine
COPY --from=build-step /go/bin/document-server /usr/bin/document-server
ENTRYPOINT ["/usr/bin/document-server"]
EXPOSE 8080