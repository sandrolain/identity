package storage

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sandrolain/identity/src/authnweb"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/sessions"
)

type PersistentStorage interface {
	GetEntity(entityId string) (entities.Entity, error)
	SaveEntity(entity entities.Entity) error
	DeleteEntity(entityId string) error
	GetSession(sessionId string) (sessions.Session, error)
	SaveSession(session sessions.Session) error
	DeleteSession(sessionId string) error
	GetEntitySessions(entityId string) ([]sessions.Session, error)
	SaveWebauthnCredential(credential authnweb.EntityCredential) error
	GetWebauthnCredentials(entityId string) ([]authnweb.EntityCredential, error)
}

type VolatileStorage interface {
	GetSession(sessionId string) (sessions.Session, error)
	SaveSession(session sessions.Session) error
	DeleteSession(sessionId string) error
	GetEntitySessions(entityId string) ([]sessions.Session, error)
	SaveWebauthnSessionData(sessionId string, data webauthn.SessionData) error
	GetWebauthnSessionData(sessionId string) (webauthn.SessionData, error)
	DeleteWebauthnSessionData(sessionId string) error
}
