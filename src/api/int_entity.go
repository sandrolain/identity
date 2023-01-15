package api

import (
	"fmt"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
)

func (a *API) CreateEntity(typ entities.EntityType, entityId string, password string, roles []string) (u entities.Entity, err error) {
	_, err = a.GetEntityById(entityId)
	if err == nil {
		err = fmt.Errorf(`entity "%s" already exist`, entityId)
		return
	} else if !crudutils.IsNotFound(err) {
		return
	}
	if u, err = entities.NewEntity(typ, entityId, password, a.Config.Totp); err != nil {
		return
	}
	if len(roles) > 0 {
		u.Roles.Add(roles...)
	}
	return u, a.PersistentStorage.SaveEntity(u)
}

func (a *API) GetEntityById(username string) (entities.Entity, error) {
	return a.PersistentStorage.GetEntity(username)
}

func (a *API) SetEntityPassword(username string, password string) (u entities.Entity, err error) {
	if u, err = a.PersistentStorage.GetEntity(username); err != nil {
		return u, err
	}
	if err = u.SetPassword(password); err != nil {
		return u, err
	}
	return u, a.PersistentStorage.SaveEntity(u)
}

func (a *API) ResetEntityTotp(username string) (u entities.Entity, err error) {
	if u, err = a.PersistentStorage.GetEntity(username); err != nil {
		return u, err
	}
	if err = u.ResetTotp(a.Config.Totp); err != nil {
		return u, err
	}
	return u, a.PersistentStorage.SaveEntity(u)
}

func (a *API) DeleteEntity(username string) error {
	_, err := a.PersistentStorage.GetEntity(username)
	if err != nil {
		return err
	}
	return a.PersistentStorage.DeleteEntity(username)
}

func (a *API) getTotpUri(u *entities.Entity) (uri string, err error) {
	if !u.TotpConfigured {
		err = u.ResetTotp(a.Config.Totp)
		if err != nil {
			return
		}
		err = a.PersistentStorage.SaveEntity(*u)
		if err != nil {
			return
		}
	}
	uri = u.TotpUri
	return
}

func (a *API) validateTotp(u *entities.Entity, totpCode string) (err error) {
	totpOk, err := u.ValidateTotp(totpCode)
	if err != nil {
		return
	}

	if !totpOk {
		err = crudutils.NotAuthorized("")
		return
	}

	if !u.TotpConfigured {
		u.SetTotpConfigured(true)
		err = a.PersistentStorage.SaveEntity(*u)
	}
	return
}
