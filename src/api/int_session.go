package api

import (
	"time"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/logutils"
	"github.com/sandrolain/identity/src/sessions"
)

func (a *API) CreateSessionAndJWT(scope sessions.SessionScope, username string) (string, sessions.Session, error) {
	sess, err := a.CreateSession(scope, username, []string{})
	if err != nil {
		return "", sess, err
	}
	mk := a.Config.Keys.MasterKey
	token, err := sess.CreateSessionJWT(a.Config.Jwt.Issuer, mk)
	if err != nil {
		return "", sess, err
	}
	return token, sess, nil
}

func (a *API) GetSessionScopeDuration(scope sessions.SessionScope) (dur time.Duration) {
	var minutes int
	switch scope {
	case sessions.ScopeTotp:
		minutes = a.Config.Session.TotpRequestMinutes
	case sessions.ScopeValidation:
		minutes = a.Config.Session.ValidationMinutes
	case sessions.ScopeLogin:
		minutes = a.Config.Session.LoginSessionMinutes
	case sessions.ScopeMachine:
		minutes = a.Config.Session.MachineKeyMinutes
	case sessions.ScopePassword:
		minutes = a.Config.Session.ChangePasswordMinutes
	case sessions.ScopeWebauthn:
		minutes = a.Config.Session.WebauthLoginMinutes
	}
	dur = time.Minute * time.Duration(minutes)
	return
}

func (a *API) CreateSession(scope sessions.SessionScope, username string, allowedIps []string) (s sessions.Session, err error) {
	duration := a.GetSessionScopeDuration(scope)
	if s, err = sessions.NewSession(scope, username, duration, allowedIps, a.Config.Keys.MasterKey); err != nil {
		return
	}
	err = a.VolatileStorage.SaveSession(s)
	if scope == sessions.ScopeMachine {
		err = a.PersistentStorage.SaveSession(s)
	}
	return
}

func (a *API) GetSession(scope sessions.SessionScope, sessionId string) (s sessions.Session, err error) {
	s, err = a.VolatileStorage.GetSession(sessionId)
	if err == nil {
		return
	}
	if scope == sessions.ScopeMachine {
		s, err = a.PersistentStorage.GetSession(sessionId)
		if err != nil {
			nbErr := a.VolatileStorage.SaveSession(s)
			if nbErr != nil {
				logutils.Error(nbErr, "cannot save session to volatile storage", s.Id)
			}
		}
	}
	return
}

func (a *API) ExtendSession(s sessions.Session) (res sessions.Session, err error) {
	if s.IsExpired() {
		err = crudutils.ExpiredResource(s.Id)
		return
	}
	s.Extend(a.GetSessionScopeDuration(s.Scope))
	if err = a.VolatileStorage.SaveSession(s); err != nil {
		return
	}
	res = s
	return
}

func (a *API) DeleteSession(sessionId string) (err error) {
	err = a.VolatileStorage.DeleteSession(sessionId)
	if err == nil {
		err = a.PersistentStorage.DeleteSession(sessionId)
	}
	return
}

func (a *API) GetEntitySessions(entityId string) ([]sessions.Session, error) {
	return a.VolatileStorage.GetEntitySessions(entityId)
}
