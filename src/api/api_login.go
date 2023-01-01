package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/logutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/roles"
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
	Roles          roles.Roles
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
		err = u.ResetTotp(a.Config.Totp)
		if err != nil {
			return
		}
		err = a.PersistentStorage.SaveEntity(u)
		if err != nil {
			return
		}
		res.TotpUri = u.TotpUri
	}

	return
}

func (a *API) LoginTotp(token string, otp string) (res LoginTotpResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeTotp, token)
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
		err = a.PersistentStorage.SaveEntity(u)
		if err != nil {
			return
		}
	}

	token, _, err = a.CreateSessionAndJWT(sessions.ScopeLogin, u.Id)
	if err != nil {
		return
	}

	err = a.DeleteSession(s.Id)
	if err != nil {
		logutils.Error("cannot delete totp session", err)
	}

	res.SessionToken = token

	return
}

func (a *API) GetUserDetails(token string) (res EntityDetailsResult, err error) {
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
	_, err = a.ExtendSession(s)
	if err != nil {
		logutils.Error("cannot extend Entity session", err)
	}
	return
}

func (a *API) Logout(token string) (res EntityDetailsResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}
	if u.Type != entities.TypeAdmin && u.Type != entities.TypeUser {
		err = crudutils.NotAuthorized("user type")
		return
	}
	err = a.VolatileStorage.DeleteSession(s.Id)
	if err != nil {
		logutils.Error("cannot logout", err)
	}
	return
}
