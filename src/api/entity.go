package api

import (
	"fmt"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
)

func (a *API) CreateEntity(username string, password string) (u entities.Entity, err error) {
	_, err = a.GetEntityById(username)
	if err == nil {
		return u, fmt.Errorf(`Entity "%s" already exist`, username)
	} else if !crudutils.IsNotFound(err) {
		return u, err
	}
	if u, err = entities.CreateEntity(username, password, a.Config.Totp); err != nil {
		return u, err
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

func (a *API) SetEntityTotpConfigured(username string, configured bool) (u entities.Entity, err error) {
	if u, err = a.PersistentStorage.GetEntity(username); err != nil {
		return u, err
	}
	u.SetTotpConfigured(configured)
	return u, a.PersistentStorage.SaveEntity(u)
}

func (a *API) GetEntityTotpUri(username string) (string, error) {
	u, err := a.PersistentStorage.GetEntity(username)
	return u.TotpUri, err
}

func (a *API) ResetEntityOTP(username string) (u entities.Entity, err error) {
	if u, err = a.PersistentStorage.GetEntity(username); err != nil {
		return u, err
	}
	if err = u.ResetTotp(a.Config.Totp); err != nil {
		return u, err
	}
	return u, a.PersistentStorage.SaveEntity(u)
}

func (a *API) IsEntityTotpToConfigure(username string) (bool, error) {
	u, err := a.PersistentStorage.GetEntity(username)
	if err != nil {
		return false, err
	}
	return u.IsTotpToConfigure(), nil
}

func (a *API) DeleteEntity(username string) error {
	_, err := a.PersistentStorage.GetEntity(username)
	if err != nil {
		return err
	}
	return a.PersistentStorage.DeleteEntity(username)
}
