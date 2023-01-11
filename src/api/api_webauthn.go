package api

import (
	"encoding/json"

	"github.com/sandrolain/identity/src/authnweb"
	"github.com/sandrolain/identity/src/sessions"
)

type WebauthnRegisterBeginResult struct {
	CredentialCreation string
}

type WebauthnRegisterFinishnResult struct {
	Credential string
}

type WebauthnLoginBeginResult struct {
	WebauthnToken       string
	CredentialAssertion string
}

type WebauthnLoginFinishResult struct {
	SessionToken string
}

func (a *API) WebauthnRegisterBegin(token string) (res WebauthnRegisterBeginResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}

	cred, data, err := authnweb.RegisterBegin(u, a.Config.WebAuthn)
	if err != nil {
		return
	}

	err = a.VolatileStorage.SaveWebauthnSessionData(s.Id, data)
	if err != nil {
		return
	}

	jsonCred, err := json.Marshal(cred)
	if err != nil {
		return
	}

	res.CredentialCreation = string(jsonCred)

	return
}

func (a *API) WebauthnRegisterFinish(token string, request []byte) (res WebauthnRegisterFinishnResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeLogin, token)
	if err != nil {
		return
	}

	data, err := a.VolatileStorage.GetWebauthnSessionData(s.Id)
	if err != nil {
		return
	}

	cred, err := authnweb.RegisterFinish(u, data, request, a.Config.WebAuthn)
	if err != nil {
		return
	}

	err = a.PersistentStorage.SaveWebauthnCredential(cred)
	if err != nil {
		return
	}

	jsonCred, err := json.Marshal(cred)
	if err != nil {
		return
	}

	res.Credential = string(jsonCred)

	return
}

func (a *API) WebauthnLoginBegin(entityId string) (res WebauthnLoginBeginResult, err error) {
	u, err := a.GetEntityById(entityId)
	if err != nil {
		return
	}

	cred, data, err := authnweb.LoginBegin(u, a.Config.WebAuthn)
	if err != nil {
		return
	}

	token, s, err := a.CreateSessionAndJWT(sessions.ScopeWebauthn, u.Id)
	if err != nil {
		return
	}

	err = a.VolatileStorage.SaveWebauthnSessionData(s.Id, data)
	if err != nil {
		return
	}

	jsonCred, err := json.Marshal(cred)
	if err != nil {
		return
	}

	res.WebauthnToken = token
	res.CredentialAssertion = string(jsonCred)

	return
}

func (a *API) WebauthnLoginFinish(webauthToken string, request []byte) (res WebauthnLoginFinishResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeWebauthn, webauthToken)
	if err != nil {
		return
	}

	data, err := a.VolatileStorage.GetWebauthnSessionData(s.Id)
	if err != nil {
		return
	}

	creds, err := a.PersistentStorage.GetWebauthnCredentials(u.Id)
	if err != nil {
		return
	}

	// TODO: credential verification
	_, err = authnweb.LoginFinish(u, data, creds, request, a.Config.WebAuthn)
	if err != nil {
		return
	}

	sssToken, _, err := a.CreateSessionAndJWT(sessions.ScopeLogin, u.Id)
	if err != nil {
		return
	}

	res.SessionToken = sssToken

	return
}
