package admingrpc

import context "context"

func (s admingrpcServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{
		TotpToken: "test",
	}, nil
}

func (s admingrpcServer) LoginConfirm(ctx context.Context, req *LoginConfirmRequest) (*LoginConfirmResponse, error) {
	return &LoginConfirmResponse{}, nil
}
