package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

func (a *API) AuthenticateWithCredentials(entityId string, password string) (u entities.Entity, err error) {
	u, err = a.PersistentStorage.GetEntity(entityId)
	if err != nil {
		return
	}
	if !u.HasPassword(password) {
		err = crudutils.NotFound(entityId)
		return
	}
	return
}

func (a *API) AuthenticateWithSessionJWT(scope sessions.SessionScope, token string) (u entities.Entity, s sessions.Session, err error) {
	info, err := jwtutils.ExtractInfoFromJWT(token)
	if err != nil {
		return
	}
	if info.Scope != string(scope) {
		err = crudutils.NotFound(info.Subject)
		return
	}

	if s, err = a.GetSession(scope, info.Subject); err != nil {
		return
	}

	if s.IsExpired() {
		err = crudutils.ExpiredResource(s.Id)
		return
	}

	kp := keys.SecureKeyParams{
		Length:    a.Config.SecureKey.Length,
		MasterKey: a.Config.SecureKey.MasterKey,
	}

	if err = s.VerifySessionJWT(token, kp); err != nil {
		return
	}

	u, err = a.GetEntityById(s.EntityId)

	return
}
