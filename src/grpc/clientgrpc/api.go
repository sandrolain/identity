package clientgrpc

import context "context"

func (s clientgrpcServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{
		TotpToken: "test",
	}, nil
}

func (s clientgrpcServer) LoginConfirm(ctx context.Context, req *LoginConfirmRequest) (*LoginConfirmResponse, error) {
	return &LoginConfirmResponse{}, nil
}
