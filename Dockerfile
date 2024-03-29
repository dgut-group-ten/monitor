FROM golang:1.13-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache gcc musl-dev git

WORKDIR /go/src/monitor

RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn,direct

# 下载依赖
COPY go.mod go.mod
RUN go mod download

COPY main.go main.go
COPY core /go/src/monitor/core

EXPOSE 8004

ENTRYPOINT ["go", "run", "main.go"]
