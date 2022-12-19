package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

func (a *API) AuthenticateWithCredentials(entityId string, password string) (*entities.Entity, error) {
	u, err := a.PersistentStorage.GetEntity(entityId)
	if err != nil {
		return nil, err
	}
	if !u.HasPassword(password) {
		return nil, crudutils.NotFound(entityId)
	}
	return &u, nil
}

func (a *API) AuthenticateWithSessionJWT(token string) (u entities.Entity, s sessions.Session, err error) {
	info, err := jwtutils.ExtractInfoFromJWT(token)
	if err != nil {
		return u, s, err
	}
	if s, err = a.GetSession(info.Subject); err != nil {
		return u, s, err
	}
	kp := keys.SecureKeyParams{
		Length:    a.Config.SecureKey.Length,
		MasterKey: a.Config.SecureKey.MasterKey,
	}
	if err = s.VerifySessionJWT(token, kp); err != nil {
		return u, s, err
	}
	if u, err = a.GetEntityById(s.EntityId); err != nil {
		return u, s, err
	}

	return u, s, nil
}
