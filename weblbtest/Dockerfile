# sysntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder

RUN apk --update add ca-certificates wget

RUN mkdir -p /build

WORKDIR /build

ADD go.mod /build
ADD server.go /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o serviced -ldflags="-w -s" ./server.go

FROM scratch AS runner

COPY --from=builder /build/serviced /serviced

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV GIN_MODE=release

ENTRYPOINT ["/serviced"]
