package clientgrpc

import (
	context "context"

	"github.com/sandrolain/identity/src/entities"
)

func (s clientgrpcServer) Login(ctx context.Context, req *LoginRequest) (res *LoginResponse, err error) {
	r, err := s.Api.Login(entities.TypeUser, req.Email, req.Password)
	if err == nil {
		res = &LoginResponse{
			TotpToken: r.TotpToken,
			TotpUri:   r.TotpUri,
		}
	}
	return
}

func (s clientgrpcServer) LoginConfirm(ctx context.Context, req *LoginConfirmRequest) (res *LoginConfirmResponse, err error) {
	r, err := s.Api.LoginTotp(req.TotpToken, req.TotpCode)
	if err == nil {
		res = &LoginConfirmResponse{
			SessionToken: r.SessionToken,
		}
	}
	return
}

func (s clientgrpcServer) GetUserDetails(ctx context.Context, req *GetUserDetailsRequest) (res *GetUserDetailsResponse, err error) {
	r, err := s.Api.GetUserDetails(req.SessionToken)
	if err == nil {
		res = &GetUserDetailsResponse{
			Id:             r.EntityId,
			Type:           int32(r.Type),
			Roles:          res.Roles,
			TotpConfigured: r.TotpConfigured,
			TotpUri:        r.TotpUri,
		}
	}
	return
}

func (s clientgrpcServer) AuthenticateMachine(ctx context.Context, req *AuthenticateMachineRequest) (res *AuthenticateMachineResponse, err error) {
	r, err := s.Api.AuthenticateMachine(req.MachineToken, req.Email, req.Ip)
	if err == nil {
		res = &AuthenticateMachineResponse{
			Id:    r.MachineId,
			Type:  int32(r.Type),
			Roles: res.Roles,
		}
	}
	return
}
