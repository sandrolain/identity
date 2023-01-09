package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/logutils"
	"github.com/sandrolain/identity/src/sessions"
)

type InitEntityValidationResult struct {
	ValidationToken string
}

type CompleteEntityValidationResult struct {
	SessionToken string
}

func (a *API) InitEntityValidation(token string) (res InitEntityValidationResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}

	if u.Validated {
		err = crudutils.InvalidValue("")
		return
	}

	vldToken, _, err := a.CreateSessionAndJWT(sessions.ScopeValidation, u.Id)
	if err != nil {
		return
	}

	res.ValidationToken = vldToken
	return
}

func (a *API) CompleteEntityValidation(token string) (res CompleteEntityValidationResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeValidation, token)
	if err != nil {
		return
	}

	if u.Validated {
		err = crudutils.InvalidValue("")
		return
	}

	u.Validate()

	err = a.PersistentStorage.SaveEntity(u)
	if err != nil {
		return
	}

	sssToken, _, err := a.CreateSessionAndJWT(sessions.ScopeLogin, u.Id)
	if err != nil {
		return
	}

	err = a.DeleteSession(s.Id)
	if err != nil {
		logutils.Error(err, "cannot delete validation session: %v", s.Id)
	}

	res.SessionToken = sssToken

	return
}
