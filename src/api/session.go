package api

import (
	"time"

	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

func (a *API) CreateSessionAndJWT(scope sessions.SessionScope, username string) (string, error) {
	kp := keys.SecureKeyParams{
		Length:    a.Config.SecureKey.Length,
		MasterKey: a.Config.SecureKey.MasterKey,
	}
	sess, err := a.CreateSession(scope, username, kp)
	if err != nil {
		return "", err
	}
	token, err := sess.CreateSessionJWT(a.Config.Jwt.Issuer, kp)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *API) CreateSession(scope sessions.SessionScope, username string, kp keys.SecureKeyParams) (s sessions.Session, err error) {
	var duration time.Duration
	switch scope {
	case sessions.ScopeTotp:
		duration = time.Minute * time.Duration(a.Config.Session.TotpRequestMinutes)
	case sessions.ScopeLogin:
		duration = time.Minute * time.Duration(a.Config.Session.LoginSessionMinutes)
	case sessions.ScopeMachine:
		duration = time.Minute * time.Duration(a.Config.Session.MachineKeyMinutes)
	}
	if s, err = sessions.NewSession(scope, username, duration, kp); err != nil {
		return
	}
	err = a.VolatileStorage.SaveSession(s)
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

func (a *API) ExtendSession(s sessions.Session, duration time.Duration) (sessions.Session, error) {
	if s.IsExpired() {
		return s, &sessions.SessionExpiredError{}
	}
	s.Extend(duration)
	if err := a.VolatileStorage.SaveSession(s); err != nil {
		return s, err
	}
	return s, nil
}

func (a *API) DeleteSession(sessionId string) error {
	return a.VolatileStorage.DeleteSession(sessionId)
}

func (a *API) GetEntitySessions(sessionId string) ([]sessions.Session, error) {
	return a.VolatileStorage.GetEntitySessions(sessionId)
}
