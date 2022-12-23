package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/sessions"
)

type LoginResult struct {
	TotpUri   string
	TotpToken string
}

type LoginTotpResult struct {
	SessionToken string
}

type EntityDetailsResult struct {
	EntityId       string
	Type           entities.EntityType
	Roles          entities.EntityRoles
	TotpConfigured bool
	TotpUri        string
}

func (a *API) Login(entityType entities.EntityType, entityId string, password string) (res LoginResult, err error) {
	if !entities.ValidEntityId(entityId) {
		err = crudutils.InvalidValue(entityId)
		return
	}

	u, err := a.AuthenticateWithCredentials(entityId, password)
	if err != nil {
		return
	}

	if u.Type != entityType || !u.IsEnabled() {
		err = crudutils.NotFound(entityId)
		return
	}

	token, _, err := a.CreateSessionAndJWT(sessions.ScopeTotp, u.Id)
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
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeTotp, token)
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

	if !u.TotpConfigured {
		u.SetTotpConfigured(true)
		a.PersistentStorage.SaveEntity(u)
	}

	token, _, err = a.CreateSessionAndJWT(sessions.ScopeLogin, u.Id)
	if err != nil {
		return
	}

	res.SessionToken = token

	return
}

func (a *API) GetEntityDetails(token string) (res EntityDetailsResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}
	res.EntityId = u.Id
	res.Type = u.Type
	res.Roles = u.Roles
	res.TotpConfigured = u.TotpConfigured
	if !u.TotpConfigured {
		res.TotpUri = u.TotpUri
	}
	a.ExtendSession(s)
	return
}
