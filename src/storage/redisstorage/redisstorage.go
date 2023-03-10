package redisstorage

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/redisutils"
	"github.com/sandrolain/identity/src/sessions"
	"github.com/sandrolain/identity/src/storage"
)

type RedisStorage struct {
	client *redisutils.Client
}

func (s *RedisStorage) GetSession(sessionId string) (sessions.Session, error) {
	var sess sessions.Session
	found, err := s.client.Get(redisutils.Key{"sessions", sessionId}, &sess)
	if err != nil {
		return sess, err
	}
	if !found {
		return sess, crudutils.NotFound(sessionId)
	}
	return sess, err
}
func (s *RedisStorage) GetEntitySessions(entityId string) (res []sessions.Session, err error) {
	var id string
	all, err := s.client.GetAll(redisutils.Key{"entitySessions", entityId}, &id)
	if err != nil {
		return
	}

	ids, err := redisutils.AllAsType[string](all)
	if err != nil {
		return
	}

	var sess sessions.Session
	for _, id := range ids {
		sess, err = s.GetSession(id)
		if err != nil {
			return
		}
		res = append(res, sess)
	}

	return
}
func (s *RedisStorage) SaveSession(sess sessions.Session) error {
	err := s.client.Set(redisutils.Key{"entitySessions", sess.EntityId, sess.Id}, sess.Id, sess.GetTtl())
	if err != nil {
		return err
	}
	return s.client.Set(redisutils.Key{"sessions", sess.Id}, sess, sess.GetTtl())
}
func (s *RedisStorage) DeleteSession(sessionId string) error {
	return s.client.Delete(redisutils.Key{"sessions", sessionId})
}
func (s *RedisStorage) SaveWebauthnSessionData(sessionId string, data webauthn.SessionData) error {
	return s.client.Set(redisutils.Key{"webauthn", sessionId}, data, time.Minute*15) // TODO: duration from config
}
func (s *RedisStorage) GetWebauthnSessionData(sessionId string) (res webauthn.SessionData, err error) {
	found, err := s.client.Get(redisutils.Key{"webauthn", sessionId}, &res)
	if err != nil {
		return
	}
	if !found {
		err = crudutils.NotFound(sessionId)
		return
	}
	return
}
func (s *RedisStorage) DeleteWebauthnSessionData(sessionId string) error {
	return s.client.Delete(redisutils.Key{"webauthn", sessionId})
}

func CreateRedisStorage(host string, password string, tls *tls.Config, timeout time.Duration) (storage.VolatileStorage, error) {
	client, err := redisutils.NewClient(host, password, tls, timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot create new Redis client: %v", err)
	}
	store := &RedisStorage{client: client}
	return store, nil
}
