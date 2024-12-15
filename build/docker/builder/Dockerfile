FROM golang:1.18.3-bullseye as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

ENV TZ Asia/Shanghai

WORKDIR /vtrace

COPY . .

RUN go mod tidy && make go.build

