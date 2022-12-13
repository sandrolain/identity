# Notes


https://github.com/protocolbuffers/protobuf/releases


sudo mv ./include/google /usr/local
sudo mv ./bin/protoc /usr/local/bin

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```sh
protoc --go_out=. --go-grpc_out=. ./clientgrpc.proto
```