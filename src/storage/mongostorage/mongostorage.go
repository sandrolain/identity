package mongostorage

import (
	"fmt"
	"time"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/mongoutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/sessions"
)

const (
	ENTITIES_COLLECTION      = "entities"
	SESSIONS_COLLECTION      = "sessions"
	KEYS_COLLECTION          = "keys"
	EXPIRING_KEYS_COLLECTION = "expiringKeys"
)

type MongoDBStorage struct {
	client *mongoutils.Client
}

func (s *MongoDBStorage) Setup() error {
	_, err := s.client.AssertIndex(SESSIONS_COLLECTION, "entityId")
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoDBStorage) GetEntity(entityId string) (entities.Entity, error) {
	var u entities.Entity
	found, err := s.client.FindOneByField(ENTITIES_COLLECTION, "_id", entityId, &u)
	if err != nil {
		return u, err
	}
	if !found {
		return u, crudutils.NotFound(entityId)
	}
	return u, nil
}
func (s *MongoDBStorage) SaveEntity(u entities.Entity) error {
	_, err := s.client.UpsertOneById(ENTITIES_COLLECTION, u.Id, u)
	return err
}
func (s *MongoDBStorage) DeleteEntity(entityId string) error {
	_, err := s.client.DeleteOneById(ENTITIES_COLLECTION, entityId)
	return err
}

func (s *MongoDBStorage) GetSession(sessionId string) (sessions.Session, error) {
	var u sessions.Session
	found, err := s.client.FindOneById(SESSIONS_COLLECTION, sessionId, &u)
	if err != nil {
		return u, err
	}
	if !found {
		return u, crudutils.NotFound(sessionId)
	}
	return u, nil
}
func (s *MongoDBStorage) GetAllSessions(entityId string) ([]sessions.Session, error) {
	var u []sessions.Session
	err := s.client.FindManyByField(SESSIONS_COLLECTION, "entityId", entityId, 1, -1, &u)
	return u, err
}
func (s *MongoDBStorage) SaveSession(sess sessions.Session) error {
	_, err := s.client.UpsertOneById(SESSIONS_COLLECTION, sess.Id, sess)
	return err
}
func (s *MongoDBStorage) DeleteSession(sessionId string) error {
	_, err := s.client.DeleteOneById(SESSIONS_COLLECTION, sessionId)
	return err
}
func (s *MongoDBStorage) DeleteAllSessions(entityId string) error {
	_, err := s.client.DeleteMany(SESSIONS_COLLECTION, map[string]string{"entityId": entityId})
	return err
}

func (s *MongoDBStorage) GetKey(name string) (keys.SecuredKey, error) {
	var u keys.SecuredKey
	found, err := s.client.FindOneByField(KEYS_COLLECTION, "_id", name, &u)
	if err != nil {
		return u, err
	}
	if !found {
		return u, crudutils.NotFound(name)
	}
	return u, nil
}
func (s *MongoDBStorage) UpdateKey(name string, k *keys.SecuredKey) (*keys.SecuredKey, error) {
	_, err := s.client.UpdateOneById(KEYS_COLLECTION, name, k)
	if err != nil {
		return nil, err
	}
	return k, nil
}
func (s *MongoDBStorage) DeleteKey(name string) error {
	_, err := s.client.DeleteOneById(KEYS_COLLECTION, name)
	return err
}

func CreateMongoDBStorage(uri string, database string, timeout time.Duration) (*MongoDBStorage, error) {
	client, err := mongoutils.NewClient(uri, database, timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot create new MongoDB client: %v", err)
	}
	store := &MongoDBStorage{client: client}
	err = store.Setup()
	if err != nil {
		return nil, fmt.Errorf("cannot setup new MongoDB client: %v", err)
	}
	return store, nil
}
