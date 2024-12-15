#!/bin/bash
GOPATH=`go env GOPATH` && \
BIN="/bin" && \
cp scripts/bin/buf/buf $GOPATH$BIN