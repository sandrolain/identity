package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/roles"
	"github.com/sandrolain/identity/src/sessions"
)

type CreateUserResult struct {
	UserId string
}

func (a *API) CreateUser(token string, entityId string, machineRoles roles.Roles) (res CreateUserResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}
	if !u.IsAdmin() || !u.Roles.Has(roles.RoleMachinesManager) {
		err = crudutils.NotAuthorized("")
		return
	}
	e, err := a.CreateEntity(entities.TypeUser, entityId, "", machineRoles)
	if err != nil {
		return
	}
	res.UserId = e.Id
	return
}

type CreateMachineResult struct {
	MachineId string
}

func (a *API) CreateMachine(token string, entityId string, machineRoles roles.Roles) (res CreateMachineResult, err error) {
	u, _, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}
	if !u.IsAdmin() || !u.Roles.Has(roles.RoleMachinesManager) {
		err = crudutils.NotAuthorized("")
		return
	}
	e, err := a.CreateEntity(entities.TypeMachine, entityId, "", machineRoles)
	if err != nil {
		return
	}
	res.MachineId = e.Id
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
	if !u.IsAdmin() || !u.Roles.Has(roles.RoleMachinesManager) {
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
	k, err := sess.Key.Unsecure(a.Config.Keys.MasterKey)
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
