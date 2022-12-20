package clientgrpc

import (
	context "context"
	"fmt"
)

func (s clientgrpcServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	res, err := s.Api.Login(req.Email, req.Password)
	fmt.Printf("res: %v\n", res)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		TotpToken: res.TotpToken,
		TotpUri:   res.TotpUri,
	}, nil
}

func (s clientgrpcServer) LoginConfirm(ctx context.Context, req *LoginConfirmRequest) (*LoginConfirmResponse, error) {
	return &LoginConfirmResponse{}, nil
}
