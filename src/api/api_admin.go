package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/sessions"
)

type InsertMachineResult struct {
	MachineId string
}

func (a *API) InsertMachine(token string, entityId string, roles []string) (res InsertMachineResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}
	if !u.IsAdmin() {
		err = crudutils.NotAuthorized("")
		return
	}
	e, err := a.CreateEntity(entities.TypeMachine, entityId, "", roles)
	if err != nil {
		return
	}
	res = InsertMachineResult{
		MachineId: e.Id,
	}
	return
}

type MachineSessionResult struct {
	MachineId string
	SessionId string
	Secret    []byte
}

func (a *API) InitMachineSession(token string, entityId string) (res MachineSessionResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}
	if !u.IsAdmin() {
		err = crudutils.NotAuthorized("")
		return
	}
	e, err := a.GetEntityById(entityId)
	if err != nil {
		return
	}
	if !e.IsMachine() {
		err = crudutils.InvalidValue(entityId)
		return
	}
	sess, err := a.CreateSession(sessions.ScopeLogin, u.Id)
	if err != nil {
		return
	}
	k, err := sess.Key.Unsecure(a.Config.SecureKey.MasterKey)
	if err != nil {
		return
	}

	res = MachineSessionResult{
		MachineId: sess.EntityId,
		SessionId: sess.Id,
		Secret:    k.Value,
	}
	return
}
