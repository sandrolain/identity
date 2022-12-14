package admingrpc

import (
	context "context"

	"github.com/sandrolain/identity/src/entities"
)

func (s admingrpcServer) Login(ctx context.Context, req *LoginRequest) (res *LoginResponse, err error) {
	r, err := s.Api.Login(entities.TypeAdmin, req.Email, req.Password)
	if err == nil {
		res = &LoginResponse{
			TotpToken: r.TotpToken,
			TotpUri:   r.TotpUri,
		}
	}
	return
}

func (s admingrpcServer) LoginConfirm(ctx context.Context, req *LoginConfirmRequest) (res *LoginConfirmResponse, err error) {
	r, err := s.Api.LoginTotp(req.TotpToken, req.TotpCode)
	if err == nil {
		res = &LoginConfirmResponse{
			SessionToken: r.SessionToken,
		}
	}
	return
}

func (s admingrpcServer) Logout(ctx context.Context, req *LogoutRequest) (res *LogoutResponse, err error) {
	r, err := s.Api.Logout(req.SessionToken)
	if err == nil {
		res = &LogoutResponse{
			Id:        r.EntityId,
			SessionId: r.SessionId,
		}
	}
	return
}

func (s admingrpcServer) CreateMachine(ctx context.Context, req *CreateMachineRequest) (res *CreateMachineResponse, err error) {
	r, err := s.Api.CreateMachine(req.SessionToken, req.Email, req.Roles)
	if err == nil {
		res = &CreateMachineResponse{
			MachineId: r.MachineId,
		}
	}
	return
}

func (s admingrpcServer) InitMachineSession(ctx context.Context, req *InitMachineSessionRequest) (res *InitMachineSessionResponse, err error) {
	r, err := s.Api.InitMachineSession(req.SessionToken, req.MachineId, req.AllowedIps)
	if err == nil {
		res = &InitMachineSessionResponse{
			MachineId: r.MachineId,
			SessionId: r.SessionId,
			Subject:   r.Subject,
			Secret:    r.Secret,
		}
	}
	return
}
