FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git make

WORKDIR $GOPATH/src/go-cli-template
COPY . .
RUN make build


FROM scratch

COPY --from=builder /go/src/go-cli-template/bin/go-cli-template /bin/go-cli-template
ENTRYPOINT ["/bin/go-cli-template"]
