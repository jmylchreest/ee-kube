FROM golang
ADD ./go-hello /go/src/github.com/jmylchreest/go-hello
RUN go install github.com/jmylchreest/go-hello
ENTRYPOINT [ "/go/bin/go-hello" ]
EXPOSE 8081