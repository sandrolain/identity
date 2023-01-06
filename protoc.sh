#!/bin/sh

cd ./src/grpc
protoc --go_out=. --go-grpc_out=. ./clientgrpc.proto
protoc --go_out=. --go-grpc_out=. ./admingrpc.proto

cp ./clientgrpc.proto ../../src-svc/waweb/clientgrpc.proto
cd ../../src-svc/waweb
protoc --go_out=. --go-grpc_out=. ./clientgrpc.proto