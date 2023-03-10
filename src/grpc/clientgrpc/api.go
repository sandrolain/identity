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
			SessionToken:    r.SessionToken,
			ValidationToken: r.ValidationToken,
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

func (s clientgrpcServer) Logout(ctx context.Context, req *LogoutRequest) (res *LogoutResponse, err error) {
	r, err := s.Api.Logout(req.SessionToken)
	if err == nil {
		res = &LogoutResponse{
			Id:        r.EntityId,
			SessionId: r.SessionId,
		}
	}
	return
}

func (s clientgrpcServer) LogoutAllSessions(ctx context.Context, req *LogoutAllSessionsRequest) (res *LogoutAllSessionsResponse, err error) {
	r, err := s.Api.LogoutAllSessions(req.SessionToken)
	if err == nil {
		res = &LogoutAllSessionsResponse{
			Id:    r.EntityId,
			Count: r.SessionsCount,
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

func (s clientgrpcServer) InitValidation(ctx context.Context, req *InitValidationRequest) (res *InitValidationResponse, err error) {
	r, err := s.Api.InitEntityValidation(entities.TypeUser, req.Email)
	if err == nil {
		res = &InitValidationResponse{
			ValidationToken: r.ValidationToken,
		}
	}
	return
}

func (s clientgrpcServer) VerifyValidation(ctx context.Context, req *VerifyValidationRequest) (res *VerifyValidationResponse, err error) {
	r, err := s.Api.VerifyEntityValidation(entities.TypeUser, req.ValidationToken)
	if err == nil {
		res = &VerifyValidationResponse{
			TotpToken: r.TotpToken,
			TotpUri:   r.TotpUri,
		}
	}
	return
}

func (s clientgrpcServer) CompleteValidation(ctx context.Context, req *CompleteValidationRequest) (res *CompleteValidationResponse, err error) {
	r, err := s.Api.CompleteEntityValidation(entities.TypeUser, req.TotpToken, req.TotpCode)
	if err == nil {
		res = &CompleteValidationResponse{
			SessionToken: r.SessionToken,
		}
	}
	return
}

func (s clientgrpcServer) PasswordChange(ctx context.Context, req *PasswordChangeRequest) (res *PasswordChangeResponse, err error) {
	r, err := s.Api.EntityPasswordChange(req.SessionToken, req.TotpCode, req.Password)
	if err == nil {
		res = &PasswordChangeResponse{
			SessionToken: r.SessionToken,
		}
	}
	return
}

func (s clientgrpcServer) BeginWebauthnRegister(ctx context.Context, req *BeginWebauthnRegisterRequest) (res *BeginWebauthnRegisterResponse, err error) {
	r, err := s.Api.WebauthnRegisterBegin(req.SessionToken)
	if err == nil {
		res = &BeginWebauthnRegisterResponse{
			CredentialCreation: r.CredentialCreation,
		}
	}
	return
}

func (s clientgrpcServer) FinishWebauthnRegister(ctx context.Context, req *FinishWebauthnRegisterRequest) (res *FinishWebauthnRegisterResponse, err error) {
	r, err := s.Api.WebauthnRegisterFinish(req.SessionToken, req.Request)
	if err == nil {
		res = &FinishWebauthnRegisterResponse{
			SessionToken: r.SessionToken,
		}
	}
	return
}

func (s clientgrpcServer) BeginWebauthnLogin(ctx context.Context, req *BeginWebauthnLoginRequest) (res *BeginWebauthnLoginResponse, err error) {
	r, err := s.Api.WebauthnLoginBegin(req.Email)
	if err == nil {
		res = &BeginWebauthnLoginResponse{
			WebauthnToken:       r.WebauthnToken,
			CredentialAssertion: r.CredentialAssertion,
		}
	}
	return
}

func (s clientgrpcServer) FinishWebauthnLogin(ctx context.Context, req *FinishWebauthnLoginRequest) (res *FinishWebauthnLoginResponse, err error) {
	r, err := s.Api.WebauthnLoginFinish(req.WebauthnToken, req.Request)
	if err == nil {
		res = &FinishWebauthnLoginResponse{
			SessionToken: r.SessionToken,
		}
	}
	return
}

func (s clientgrpcServer) ValidateEmail(ctx context.Context, req *ValidateEmailRequest) (res *ValidateEmailResponse, err error) {
	r, err := s.Api.ValidateEmail(req.Email)
	if err == nil {
		res = &ValidateEmailResponse{
			Valid: r.Valid,
		}
	}
	return
}

func (s clientgrpcServer) ValidatePassword(ctx context.Context, req *ValidatePasswordRequest) (res *ValidatePasswordResponse, err error) {
	r, err := s.Api.ValidatePassword(req.Password)
	if err == nil {
		res = &ValidatePasswordResponse{
			Valid: r.Valid,
		}
	}
	return
}
