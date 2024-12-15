FROM mysql:8.0

LABEL version="v0.0.1" description="mysql-server" by="yangchnet"

COPY build/docker/mysql/ddl /docker-entrypoint-initdb.d/

COPY build/docker/mysql/my.cnf /etc/mysql/my.cnf

EXPOSE 13306

