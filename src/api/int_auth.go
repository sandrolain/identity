package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/identity/src/entities"
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
	infoSubject, err := sessions.ParseScopeSubject(info.Subject)
	if err != nil {
		return
	}

	if infoSubject.Scope != string(scope) {
		err = crudutils.NotFound(info.Subject)
		return
	}

	if s, err = a.GetSession(scope, infoSubject.SessionId); err != nil {
		return
	}

	if s.IsExpired() {
		err = crudutils.ExpiredResource(s.Id)
		return
	}

	if err = s.VerifySessionJWT(token, a.Config.Keys.MasterKey); err != nil {
		return
	}

	u, err = a.GetEntityById(s.EntityId)

	return
}
