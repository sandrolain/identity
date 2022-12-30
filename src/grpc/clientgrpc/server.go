package clientgrpc

import (
	"fmt"
	"net"

	"github.com/sandrolain/go-utilities/pkg/logutils"
	"github.com/sandrolain/identity/src/api"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type clientgrpcServer struct {
	Api *api.API
	UnimplementedClientServiceServer
}

func StartServer(a *api.API) error {
	cfg := a.Config.ClientGrpc
	address := fmt.Sprintf("localhost:%d", cfg.Port)
	logutils.Infof("starting client gRPC server on %v", address)
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
	RegisterClientServiceServer(grpcServer, clientgrpcServer{Api: a})
	return grpcServer.Serve(lis)
}
