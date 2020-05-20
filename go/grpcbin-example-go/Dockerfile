# build
FROM golang:1.14-alpine as builder
RUN             apk update && apk add git
COPY            . /go/src/github.com/moul/grpcbin-example/go
WORKDIR         /go/src/github.com/moul/grpcbin-example/go
#RUN             go get -v ./...
RUN             go build -o /go/bin/grpcbin-example-go -v .

# minimal runtime
FROM            alpine
COPY            --from=builder /go/bin/grpcbin-example-go /bin/grpcbin-example-go
ENTRYPOINT      ["/bin/grpcbin-example-go"]
