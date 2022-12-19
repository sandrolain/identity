package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/sessions"
)

type LoginResult struct {
	TotpUri   string
	TotpToken string
}

type LoginTotpResult struct {
	SessionToken string
}

func (a *API) Login(entityId string, password string) (res LoginResult, err error) {
	u, err := a.AuthenticateWithCredentials(entityId, password)
	if err != nil {
		return
	}

	token, err := a.CreateSessionAndJWT(sessions.SCOPE_OTP, u.Id)
	if err != nil {
		return
	}

	res.TotpToken = token

	if !u.TotpConfigured {
		res.TotpUri = u.TotpUri
	}

	return
}

func (a *API) LoginTotp(token string, otp string) (res LoginTotpResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(token)
	if err != nil {
		return
	}

	totpOk, err := u.ValidateTotp(otp)
	if err != nil {
		return
	}
	if !totpOk {
		err = crudutils.NotAuthorized("")
		return
	}

	token, err = a.CreateSessionAndJWT(sessions.SCOPE_SESSION, u.Id)
	if err != nil {
		return
	}

	res.SessionToken = token

	return
}
