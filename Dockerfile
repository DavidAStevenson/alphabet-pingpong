FROM golang as builder
ENV GO_SRC_DIR=/go/src
COPY . ${GO_SRC_DIR}
WORKDIR ${GO_SRC_DIR}
RUN go get github.com/nats-io/go-nats && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -a -tags netgo -ldflags '-w' -o $GOPATH/bin/alphabet-pingpong \
alphabet-pingpong.go

FROM scratch
COPY --from=builder /go/bin/alphabet-pingpong /alphabet-pingpong
ENTRYPOINT ["/alphabet-pingpong"]
CMD ["nats://192.168.99.100:4222", "A"]
