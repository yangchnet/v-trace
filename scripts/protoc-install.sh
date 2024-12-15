#!/bin/bash
PROTOC_ZIP=scripts/bin/protoc/protoc-3.19.4-linux-x86_64.zip
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP