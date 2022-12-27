package mongostorage

import (
	"testing"
	"time"

	"github.com/sandrolain/identity/src/storage"
)

func checkStorageInterface(s storage.PersistentStorage) {
	return
}

func TestStorageInterface(t *testing.T) {
	s, _ := CreateMongoDBStorage("", "", time.Second)
	checkStorageInterface(s)
}
