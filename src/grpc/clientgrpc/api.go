package clientgrpc

import (
	context "context"

	"github.com/sandrolain/identity/src/entities"
)

func (s clientgrpcServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	res, err := s.Api.Login(entities.TypeUser, req.Email, req.Password)
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
