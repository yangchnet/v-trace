FROM ubuntu:20.04

ENV TZ Asia/Shanghai

RUN apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends ca-certificates curl

ARG SERVICE_NAME=

ARG SERVICE_PORT=

WORKDIR .

COPY --from=vtrace/builder /vtrace/_output/platforms/linux/amd64/${SERVICE_NAME}  ./service

COPY app/${SERVICE_NAME}/config/ app/${SERVICE_NAME}/config/

COPY build/chainmaker build/chainmaker

ENTRYPOINT ["./service"]

EXPOSE ${SERVICE_PORT}