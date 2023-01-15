package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/sessions"
)

type InitEntityPasswordChangeResult struct {
	PasswordChangeToken string
}

type CompleteEntityPasswordChangeResult struct {
	SessionToken   string
	ValidatioToken string
}

func (a *API) EntityPasswordChange(token string, totpCode string, password string) (res CompleteEntityPasswordChangeResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeAuthChange, token)
	if err != nil {
		return
	}

	err = a.validateTotp(&u, totpCode)
	if err != nil {
		return
	}

	if password == "" {
		err = crudutils.InvalidValue("password")
		return
	}

	err = u.SetPassword(password)
	if err != nil {
		return
	}

	err = a.PersistentStorage.SaveEntity(u)
	if err != nil {
		return
	}

	res.SessionToken = token

	return
}
