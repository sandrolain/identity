package admingrpc

import (
	"fmt"
	"log"
	"net"

	"github.com/sandrolain/identity/src/api"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type admingrpcServer struct {
	UnimplementedAdminServiceServer
	Api *api.API
}

func StartServer(a *api.API) error {
	cfg := a.Config.AdminGrpc
	address := fmt.Sprintf("localhost:%d", cfg.Port)
	log.Printf("start gRPC server on %v", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer lis.Close()
	creds, err := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return fmt.Errorf("failed to create gRPC TLS: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	RegisterAdminServiceServer(grpcServer, admingrpcServer{Api: a})
	return grpcServer.Serve(lis)
}
