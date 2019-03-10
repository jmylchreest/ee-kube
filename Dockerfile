FROM golang AS builder
ADD ./go-hello /go/src/github.com/jmylchreest/go-hello
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go install github.com/jmylchreest/go-hello

FROM scratch
USER nobody
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/go-hello /go/bin/go-hello
ENTRYPOINT [ "/go/bin/go-hello" ]
EXPOSE 8081