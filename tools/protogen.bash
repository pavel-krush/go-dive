#!/usr/bin/env bash

SRC_DIR=proto/dive
GO_PACKAGE="github.com/pavel-krush/dive"

rm -rf api/go/*

echo " ++ Go"
find $SRC_DIR -type f -name *.proto | xargs \
protoc -I proto \
  --go_out=. --go_opt=module=$GO_PACKAGE \
  --go-grpc_out=. --go-grpc_opt=module=$GO_PACKAGE \
  --grpc-gateway_out=logtostderr=true,module=$GO_PACKAGE,allow_delete_body=true:.
