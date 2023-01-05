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
	CredentialAssertion string
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

func (a *API) WebauthnRegisterFinish(token string, request []byte) (res WebauthnRegisterBeginResult, err error) {
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

	res.CredentialCreation = string(jsonCred)

	return
}

func (a *API) WebauthnLoginBegin(token string) (res WebauthnLoginBeginResult, err error) {
	u, s, err := a.AuthenticateWithSessionJWT(sessions.ScopeWebauthn, token)
	if err != nil {
		return
	}

	cred, data, err := authnweb.LoginBegin(u, a.Config.WebAuthn)
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

	res.CredentialAssertion = string(jsonCred)

	return
}
