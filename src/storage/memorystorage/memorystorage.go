package memorystorage

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

type MemoryStorage struct {
	entities         map[string]entities.Entity
	sessions         map[string]sessions.Session
	entitiesSessions map[string]map[string]bool
	expKeys          map[string]keys.ExpiringKeyList
}

func (s *MemoryStorage) GetEntity(entityId string) (u entities.Entity, err error) {
	u, ok := s.entities[entityId]
	if !ok {
		err = crudutils.NotFound(entityId)
	}
	return
}
func (s *MemoryStorage) SaveEntity(u entities.Entity) (err error) {
	s.entities[u.Id] = u
	return
}
func (s *MemoryStorage) DeleteEntity(entityId string) (err error) {
	delete(s.entities, entityId)
	return
}

func (s *MemoryStorage) GetSession(sessionId string) (sess sessions.Session, err error) {
	sess, ok := s.sessions[sessionId]
	if !ok {
		err = crudutils.NotFound(sessionId)
	}
	if sess.IsExpired() {
		delete(s.sessions, sessionId)
		ok = false
	}
	return
}
func (s *MemoryStorage) GetEntitySessions(entityId string) (res []sessions.Session, err error) {
	for _, sess := range s.sessions {
		if sess.EntityId == entityId {
			res = append(res, sess)
		}
	}
	return
}
func (s *MemoryStorage) SaveSession(sess sessions.Session) error {
	s.sessions[sess.Id] = sess
	return nil
}
func (s *MemoryStorage) DeleteSession(sessionId string) error {
	delete(s.sessions, sessionId)
	return nil
}
func (s *MemoryStorage) DeleteEntitySessions(entityId string) error {
	ss, found := s.entitiesSessions[entityId]
	if !found {
		return nil
	}
	for sid, _ := range ss {
		err := s.DeleteSession(sid)
		if err != nil {
			return err
		}
	}
	delete(s.entitiesSessions, entityId)
	return nil
}

func (s *MemoryStorage) GetExpiringKeys(scope string) (keys.ExpiringKeyList, error) {
	if u, ok := s.expKeys[scope]; ok {
		return u, nil
	}
	return keys.ExpiringKeyList{}, crudutils.NotFound(scope)
}
func (s *MemoryStorage) SaveExpiringKeys(scope string, keysList keys.ExpiringKeyList) error {
	s.expKeys[scope] = keysList
	return nil
}

func CreateMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		entities:         make(map[string]entities.Entity),
		sessions:         make(map[string]sessions.Session),
		entitiesSessions: make(map[string]map[string]bool),
		expKeys:          make(map[string]keys.ExpiringKeyList),
	}
}
