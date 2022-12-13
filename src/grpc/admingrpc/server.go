package admingrpc

import (
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type admingrpcServer struct {
	UnimplementedAdminServiceServer
}

func StartServer(port int) error {
	address := fmt.Sprintf("localhost:%d", port)
	log.Printf("start admin gRPC server on %v", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// TODO: SSL
	// TODO: Auth
	grpcServer := grpc.NewServer(opts...)
	RegisterAdminServiceServer(grpcServer, admingrpcServer{})
	return grpcServer.Serve(lis)
}
