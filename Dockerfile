FROM golang:1.13-alpine as build-env

ENV GO111MODULE=on

ENV BUILDPATH=demo

ENV GOPROXY=https://goproxy.cn

RUN mkdir -p /go/src/${BUILDPATH}

COPY ./ /go/src/${BUILDPATH}

WORKDIR /go/src/${BUILDPATH}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v

FROM alpine:3.10

COPY --from=build-env /go/src/demo/demo /usr/bin/godemo

CMD ["/usr/bin/godemo"]
