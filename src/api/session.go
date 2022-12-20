package api

import (
	"fmt"
	"time"

	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

func (a *API) CreateSessionAndJWT(scope string, username string) (string, error) {
	kp := keys.SecureKeyParams{
		Length:    a.Config.SecureKey.Length,
		MasterKey: a.Config.SecureKey.MasterKey,
	}
	sess, err := a.CreateSession(scope, username, kp)
	fmt.Printf("sess: %v\n", sess)
	if err != nil {
		return "", err
	}
	token, err := sess.CreateSessionJWT(a.Config.Jwt.Issuer, kp)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *API) CreateSession(scope string, username string, kp keys.SecureKeyParams) (s sessions.Session, err error) {
	var duration time.Duration
	switch scope {
	case sessions.SCOPE_OTP:
		duration = time.Minute * time.Duration(a.Config.Session.TotpRequestMinutes)
	case sessions.SCOPE_SESSION:
		duration = time.Minute * time.Duration(a.Config.Session.LoginSessionMinutes)
	case sessions.SCOPE_MACHINE:
		duration = time.Minute * time.Duration(a.Config.Session.MachineKeyMinutes)
	}
	if s, err = sessions.NewSession(scope, username, duration, kp); err != nil {
		return
	}
	err = a.VolatileStorage.SaveSession(s, duration)
	fmt.Printf("err: %v\n", err)
	return
}

func (a *API) GetSession(sessionId string) (sessions.Session, error) {
	return a.VolatileStorage.GetSession(sessionId)
}

func (a *API) ExtendSession(sessionId string, duration time.Duration) (s sessions.Session, err error) {
	s, err = a.VolatileStorage.GetSession(sessionId)
	if err != nil {
		return s, err
	}
	if s.IsExpired() {
		return s, &sessions.SessionExpiredError{}
	}
	s.Extend(duration)
	if err = a.VolatileStorage.SaveSession(s, duration); err != nil {
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
