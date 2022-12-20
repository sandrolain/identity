package storage

import (
	"time"

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
	GetEntitySessions(entityId string) ([]sessions.Session, error)
	DeleteEntitySessions(entityId string) error
}

type VolatileStorage interface {
	GetSession(sessionId string) (sessions.Session, error)
	SaveSession(session sessions.Session, ttl time.Duration) error
	DeleteSession(sessionId string) error
	GetEntitySessions(entityId string) ([]sessions.Session, error)
	DeleteEntitySessions(entityId string) error

	GetExpiringKeys(scope string) (keys.ExpiringKeyList, error)
	SaveExpiringKeys(scope string, keysList keys.ExpiringKeyList) error
}
