package api

import (
	"time"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

func (a *API) CreateSessionAndJWT(scope sessions.SessionScope, username string) (string, sessions.Session, error) {
	sess, err := a.CreateSession(scope, username)
	if err != nil {
		return "", sess, err
	}
	mk := a.Config.SecureKey.MasterKey
	token, err := sess.CreateSessionJWT(a.Config.Jwt.Issuer, mk)
	if err != nil {
		return "", sess, err
	}
	return token, sess, nil
}

func (a *API) GetSessionScopeDuration(scope sessions.SessionScope) (dur time.Duration) {
	switch scope {
	case sessions.ScopeTotp:
		dur = time.Minute * time.Duration(a.Config.Session.TotpRequestMinutes)
	case sessions.ScopeLogin:
		dur = time.Minute * time.Duration(a.Config.Session.LoginSessionMinutes)
	case sessions.ScopeMachine:
		dur = time.Minute * time.Duration(a.Config.Session.MachineKeyMinutes)
	}
	return
}

func (a *API) CreateSession(scope sessions.SessionScope, username string) (s sessions.Session, err error) {
	kp := keys.SecureKeyParams{
		Length:    a.Config.SecureKey.Length,
		MasterKey: a.Config.SecureKey.MasterKey,
	}
	duration := a.GetSessionScopeDuration(scope)
	if s, err = sessions.NewSession(scope, username, duration, kp); err != nil {
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
			a.VolatileStorage.SaveSession(s)
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

func (a *API) DeleteSession(sessionId string) error {
	return a.VolatileStorage.DeleteSession(sessionId)
}

func (a *API) GetEntitySessions(sessionId string) ([]sessions.Session, error) {
	return a.VolatileStorage.GetEntitySessions(sessionId)
}