FROM golang:1.18 as builder
WORKDIR /go/src/github.com/lei006/gmqtt-broker
COPY . .
RUN CGO_ENABLED=0 go build -o hmq -a -ldflags '-extldflags "-static"' .


FROM alpine:3.17.3
WORKDIR /
COPY --from=builder /go/src/github.com/lei006/gmqtt-broker/hmq .
EXPOSE 1883

ENTRYPOINT ["/hmq"]
