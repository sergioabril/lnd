#!/bin/sh

# Generate APIs by passing the parsed protos to ./gen
protoc -I/usr/local/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --plugin=protoc-gen-custom=$GOPATH/bin/promobile \
       --custom_out=./build \
       --custom_opt='package_name=lndmobile,target_package=github.com/lightningnetwork/lnd/lnrpc' \
       --proto_path=../lnrpc \
       rpc.proto
