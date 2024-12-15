FROM ubuntu:20.04

ENV TZ=Asia/Shanghai

COPY build/chainmaker/bin/chainmaker /chainmaker-go/bin/chainmaker
COPY build/chainmaker/lib/libwasmer.so /usr/lib/libwasmer.so
COPY build/chainmaker/config/ /chainmaker-go/config
COPY build/chainmaker/bin/wxdec /usr/bin/wxdec

ENV LD_LIBRARY_PATH=/usr/lib:$LD_LIBRARY_PATH

ENV WASMER_BACKTRACE=1

WORKDIR /chainmaker-go/bin

ENTRYPOINT [ "nohup", \
    "./chainmaker", "start", "-c", "../config/org1/chainmaker.yml" \
    ,">panic.log",  "2>&1", "&" ]