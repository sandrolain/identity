package api

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/roles"
	"github.com/sandrolain/identity/src/sessions"
)

type AuthenticateMachineResult struct {
	MachineId string
	Type      entities.EntityType
	Roles     roles.Roles
}

func (a *API) AuthenticateMachine(token string, email string, ip string) (res AuthenticateMachineResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeMachine, token)
	if err != nil {
		return
	}
	if u.Id != email {
		err = crudutils.NotAuthorized(email)
		return
	}
	ipOk, err := s.HasAllowedIP(sessions.SessionIP(ip))
	if err != nil {
		err = crudutils.InvalidValue(ip)
		return
	}
	if !ipOk {
		err = crudutils.NotAuthorized(ip)
		return
	}

	res.MachineId = u.Id
	res.Type = u.Type
	res.Roles = u.Roles

	return
}
