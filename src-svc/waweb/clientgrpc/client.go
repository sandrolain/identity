package clientgrpc

import (
	"fmt"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewClient(address string, caFile string, host string) (client ClientServiceClient, conn *grpc.ClientConn, err error) {
	creds, err := credentials.NewClientTLSFromFile(caFile, host)
	if err != nil {
		err = fmt.Errorf("failed to create TLS credentials %v", err)
		return
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err = grpc.Dial(address, opts...)
	if err != nil {
		err = fmt.Errorf("fail to dial: %v", err)
		return
	}
	client = NewClientServiceClient(conn)
	return
}
