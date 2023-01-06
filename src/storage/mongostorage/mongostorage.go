package mongostorage

import (
	"fmt"
	"time"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/mongoutils"
	"github.com/sandrolain/identity/src/authnweb"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/sessions"
	"github.com/sandrolain/identity/src/storage"
)

const (
	EntitiesCollection      = "entities"
	SessionsCollection      = "sessions"
	WebauthnCredsCollection = "waCredentials"
)

type MongoDBStorage struct {
	client *mongoutils.Client
}

func (s *MongoDBStorage) Setup() (err error) {
	if _, err := s.client.AssertIndex(SessionsCollection, "entityId"); err != nil {
		return err
	}
	if _, err := s.client.AssertTtlIndex(SessionsCollection, "expire", 1); err != nil {
		return err
	}
	if _, err := s.client.AssertIndex(WebauthnCredsCollection, "entityId"); err != nil {
		return err
	}
	return nil
}

func (s *MongoDBStorage) GetEntity(entityId string) (entities.Entity, error) {
	var u entities.Entity
	found, err := s.client.FindOneByField(EntitiesCollection, "_id", entityId, &u)
	if err != nil {
		return u, err
	}
	if !found {
		return u, crudutils.NotFound(entityId)
	}
	return u, nil
}
func (s *MongoDBStorage) SaveEntity(u entities.Entity) error {
	_, err := s.client.UpsertOneById(EntitiesCollection, u.Id, u)
	return err
}
func (s *MongoDBStorage) DeleteEntity(entityId string) error {
	_, err := s.client.DeleteOneById(EntitiesCollection, entityId)
	return err
}

func (s *MongoDBStorage) GetSession(sessionId string) (sessions.Session, error) {
	var u sessions.Session
	found, err := s.client.FindOneById(SessionsCollection, sessionId, &u)
	if err != nil {
		return u, err
	}
	if !found {
		return u, crudutils.NotFound(sessionId)
	}
	return u, nil
}
func (s *MongoDBStorage) GetEntitySessions(entityId string) ([]sessions.Session, error) {
	var u []sessions.Session
	err := s.client.FindManyByField(SessionsCollection, "entityId", entityId, 0, &u)
	return u, err
}
func (s *MongoDBStorage) SaveSession(sess sessions.Session) error {
	_, err := s.client.UpsertOneById(SessionsCollection, sess.Id, sess)
	return err
}
func (s *MongoDBStorage) DeleteSession(sessionId string) error {
	_, err := s.client.DeleteOneById(SessionsCollection, sessionId)
	return err
}
func (s *MongoDBStorage) DeleteEntitySessions(entityId string) error {
	_, err := s.client.DeleteMany(SessionsCollection, map[string]string{"entityId": entityId})
	return err
}
func (s *MongoDBStorage) SaveWebauthnCredential(cred authnweb.EntityCredential) error {
	_, err := s.client.InsertOne(WebauthnCredsCollection, cred)
	return err
}
func (s *MongoDBStorage) GetWebauthnCredentials(entityId string) (res []authnweb.EntityCredential, err error) {
	err = s.client.FindManyByField(WebauthnCredsCollection, "entityId", entityId, 0, &res)
	return
}

func CreateMongoDBStorage(uri string, database string, timeout time.Duration) (storage.PersistentStorage, error) {
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
