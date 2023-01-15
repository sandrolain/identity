package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/logutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/sessions"
)

type InitEntityValidationResult struct {
	ValidationToken string
}

type VerifyEntityValidationResult struct {
	TotpToken string
	TotpUri   string
}

type CompleteEntityValidationResult struct {
	SessionToken string
}

func (a *API) InitEntityValidation(entityType entities.EntityType, email string) (res InitEntityValidationResult, err error) {
	u, err := a.GetEntityById(email)
	if err != nil {
		return
	}

	if u.Type != entityType || !u.IsEnabled() {
		err = crudutils.NotFound(u.Id)
		return
	}

	vldToken, _, err := a.CreateSessionAndJWT(sessions.ScopeValidation, u.Id)
	if err != nil {
		return
	}

	res.ValidationToken = vldToken
	return
}

func (a *API) VerifyEntityValidation(entityType entities.EntityType, token string) (res VerifyEntityValidationResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeValidation, token)
	if err != nil {
		return
	}

	if u.Type != entityType || !u.IsEnabled() {
		err = crudutils.NotFound(u.Id)
		return
	}

	totpToken, _, err := a.CreateSessionAndJWT(sessions.ScopeTotp, u.Id)
	if err != nil {
		return
	}

	res.TotpToken = totpToken

	totpUri, err := a.getTotpUri(&u)
	if err != nil {
		return
	}

	err = a.DeleteSession(s.Id)
	if err != nil {
		logutils.Error(err, "cannot delete validation session: %v", s.Id)
	}

	res.TotpUri = totpUri

	return
}

func (a *API) CompleteEntityValidation(entityType entities.EntityType, token string, totpCode string) (res CompleteEntityValidationResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeTotp, token)
	if err != nil {
		return
	}

	if u.Type != entityType || !u.IsEnabled() {
		err = crudutils.NotFound(u.Id)
		return
	}

	err = a.validateTotp(&u, totpCode)
	if err != nil {
		return
	}

	if !u.Validated {
		u.Validate()

		err = a.PersistentStorage.SaveEntity(u)
		if err != nil {
			return
		}
	}

	sssToken, _, err := a.CreateSessionAndJWT(sessions.ScopeLogin, u.Id)
	if err != nil {
		return
	}

	err = a.DeleteSession(s.Id)
	if err != nil {
		logutils.Error(err, "cannot delete validation totp session: %v", s.Id)
	}

	res.SessionToken = sssToken

	return
}
