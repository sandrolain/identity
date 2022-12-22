package admingrpc

import (
	context "context"

	"github.com/sandrolain/identity/src/entities"
)

func (s admingrpcServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	res, err := s.Api.Login(entities.TypeAdmin, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		TotpToken: res.TotpToken,
		TotpUri:   res.TotpUri,
	}, nil
}

func (s admingrpcServer) LoginConfirm(ctx context.Context, req *LoginConfirmRequest) (*LoginConfirmResponse, error) {
	return &LoginConfirmResponse{}, nil
}
