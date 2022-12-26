package api

import (
	"fmt"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/roles"
)

func (a *API) CreateEntity(typ entities.EntityType, entityId string, password string, roles roles.Roles) (u entities.Entity, err error) {
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
	if roles.Length() > 0 {
		u.Roles.Add(roles)
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
