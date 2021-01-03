FROM golang
MAINTAINER "xiongchuanhong"
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release\
    EXPOSE=8080
WORKDIR /go/src/goimg
ADD . /go/src/goimg
RUN go build
ENTRYPOINT ["./goimg"]