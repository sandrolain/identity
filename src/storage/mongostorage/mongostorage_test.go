package mongostorage

import (
	"testing"
	"time"

	"github.com/sandrolain/identity/src/storage"
)

func checkStorageInterface(s storage.PersistentStorage) {
	s.GetEntity("hello")
}

func TestStorageInterface(t *testing.T) {
	s, _ := CreateMongoDBStorage("", "", time.Second)
	checkStorageInterface(s)
}
