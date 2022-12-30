#!/bin/sh

cd ./src/grpc
protoc --go_out=. --go-grpc_out=. ./clientgrpc.proto
protoc --go_out=. --go-grpc_out=. ./admingrpc.proto