package memorystorage

import (
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

type MemoryStorage struct {
	sessions         map[string]sessions.Session
	entitiesSessions map[string]map[string]bool
	expKeys          map[string]keys.ExpiringKeyList
}

func (s *MemoryStorage) GetSession(sessionId string) (sessions.Session, error) {
	if u, ok := s.sessions[sessionId]; ok {
		return u, nil
	}
	return sessions.Session{}, crudutils.NotFound(sessionId)
}
func (s *MemoryStorage) GetEntitySessions(entityId string) ([]sessions.Session, error) {
	ss, ok := s.entitiesSessions[entityId]
	if !ok {
		return make([]sessions.Session, 0), nil
	}
	res := make([]sessions.Session, len(ss))
	i := 0
	for sid, _ := range ss {
		sess, err := s.GetSession(sid)
		if err != nil {
			return nil, err
		}
		res[i] = sess
		i++
	}
	return res, nil
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
		sessions:         make(map[string]sessions.Session),
		entitiesSessions: make(map[string]map[string]bool),
		expKeys:          make(map[string]keys.ExpiringKeyList),
	}
}
