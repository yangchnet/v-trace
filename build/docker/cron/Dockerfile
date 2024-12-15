FROM ubuntu:20.04

ARG CRON_NAME=

ARG DEBIAN_FRONTEND=noninteractive

ENV TZ=Asia/Shanghai

RUN apt-get update && apt-get install -y apt-utils tzdata

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN dpkg-reconfigure --frontend noninteractive tzdata

WORKDIR .

COPY --from=vtrace/builder /vtrace/_output/platforms/linux/amd64/${CRON_NAME}  ./cron

COPY app/${CRON_NAME}/config/ app/${CRON_NAME}/config/

COPY build/chainmaker build/chainmaker

ENTRYPOINT ["./cron"]
