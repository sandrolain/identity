package redisstorage

import (
	"crypto/tls"
	"fmt"
	"time"

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
func (s *RedisStorage) GetEntitySessions(entityId string) ([]sessions.Session, error) {
	var sess sessions.Session
	var res []sessions.Session
	all, err := s.client.GetAll(redisutils.Key{"sessions", entityId}, &sess)
	if err != nil {
		return res, err
	}
	return redisutils.AllAsType[sessions.Session](all)
}
func (s *RedisStorage) SaveSession(sess sessions.Session) error {
	return s.client.Set(redisutils.Key{"sessions", sess.Id}, sess, sess.GetTtl())
}
func (s *RedisStorage) DeleteSession(sessionId string) error {
	return s.client.Delete(redisutils.Key{"sessions", sessionId})
}
func (s *RedisStorage) DeleteEntitySessions(entityId string) error {
	// TODO:
	return nil
}

func CreateRedisStorage(host string, password string, tls *tls.Config, timeout time.Duration) (storage.VolatileStorage, error) {
	client, err := redisutils.NewClient(host, password, tls, timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot create new Redis client: %v", err)
	}
	store := &RedisStorage{client: client}
	return store, nil
}
