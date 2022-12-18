package storage

import (
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

type PersistentStorage interface {
	GetEntity(entityId string) (entities.Entity, error)
	SaveEntity(entity entities.Entity) error
	DeleteEntity(entityId string) error

	GetSession(sessionId string) (sessions.Session, error)
	SaveSession(session sessions.Session) error
	DeleteSession(sessionId string) error
	GetAllSessions(entityId string) ([]sessions.Session, error)
	DeleteAllSessions(entityId string) error
}

type VolatileStorage interface {
	GetSession(sessionId string) (sessions.Session, error)
	SaveSession(session sessions.Session) error
	DeleteSession(sessionId string) error
	GetAllSessions(entityId string) ([]sessions.Session, error)
	DeleteAllSessions(entityId string) error

	GetExpiringKeys(scope string) (keys.ExpiringKeyList, error)
	SaveExpiringKeys(scope string, keysList keys.ExpiringKeyList) error
}
