package authnweb

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/entities"
)

type EntityCredential struct {
	webauthn.Credential
	EntityId string `json:"entityId" bson:"entityId"`
}

func RegisterBegin(entity entities.Entity, cfg config.WebAuthnConfig) (credCreation protocol.CredentialCreation, sessionData webauthn.SessionData, err error) {
	u, err := url.Parse(cfg.Origin)
	if err != nil {
		err = fmt.Errorf("invalid webauthn origin: %v", err)
		return
	}

	wa, err := webauthn.New(&webauthn.Config{
		RPDisplayName: cfg.DisplayName,
		RPID:          u.Hostname(),
		RPOrigin:      cfg.Origin,
	})
	if err != nil {
		return
	}

	user := NewUser(entity.Id, []EntityCredential{})

	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = user.CredentialExcludeList()
	}

	cc, sd, err := wa.BeginRegistration(user, registerOptions)
	if err != nil {
		return
	}
	credCreation = *cc
	sessionData = *sd

	return
}

func RegisterFinish(entity entities.Entity, sessionData webauthn.SessionData, requestBody []byte, cfg config.WebAuthnConfig) (credential EntityCredential, err error) {
	u, err := url.Parse(cfg.Origin)
	if err != nil {
		err = fmt.Errorf("invalid webauthn origin: %v", err)
		return
	}

	wa, err := webauthn.New(&webauthn.Config{
		RPDisplayName: cfg.DisplayName,
		RPID:          u.Hostname(),
		RPOrigin:      cfg.Origin,
	})
	if err != nil {
		return
	}

	user := NewUser(entity.Id, []EntityCredential{})

	r := http.Request{}
	r.Body = io.NopCloser(bytes.NewReader(requestBody))

	cred, err := wa.FinishRegistration(user, sessionData, &r)
	if err != nil {
		return
	}

	credential = EntityCredential{
		*cred,
		entity.Id,
	}

	return
}

func LoginBegin(entity entities.Entity, creds []EntityCredential, cfg config.WebAuthnConfig) (credAssert protocol.CredentialAssertion, sessionData webauthn.SessionData, err error) {
	u, err := url.Parse(cfg.Origin)
	if err != nil {
		err = fmt.Errorf("invalid webauthn origin: %v", err)
		return
	}

	wa, err := webauthn.New(&webauthn.Config{
		RPDisplayName: cfg.DisplayName,
		RPID:          u.Hostname(),
		RPOrigin:      cfg.Origin,
	})
	if err != nil {
		return
	}

	user := NewUser(entity.Id, creds)

	ca, sd, err := wa.BeginLogin(user)
	if err != nil {
		return
	}
	credAssert = *ca
	sessionData = *sd

	return
}

func LoginFinish(entity entities.Entity, sessionData webauthn.SessionData, creds []EntityCredential, requestBody []byte, cfg config.WebAuthnConfig) (credential EntityCredential, err error) {
	u, err := url.Parse(cfg.Origin)
	if err != nil {
		err = fmt.Errorf("invalid webauthn origin: %v", err)
		return
	}

	wa, err := webauthn.New(&webauthn.Config{
		RPDisplayName: cfg.DisplayName,
		RPID:          u.Hostname(),
		RPOrigin:      cfg.Origin,
	})
	if err != nil {
		return
	}

	user := NewUser(entity.Id, creds)

	bodyReader := io.NopCloser(bytes.NewReader(requestBody))
	parsedResponse, err := protocol.ParseCredentialRequestResponseBody(bodyReader)
	if err != nil {
		return
	}

	cred, err := wa.ValidateLogin(user, sessionData, parsedResponse)
	if err != nil {
		return
	}

	credential = EntityCredential{
		*cred,
		entity.Id,
	}

	return
}
