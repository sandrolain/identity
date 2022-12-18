package memorystorage

import (
	"testing"

	"github.com/sandrolain/identity/src/storage"
)

func checkStorageInterface(s storage.VolatileStorage) {
	s.GetSession("hello")
}

func TestStorageInterface(t *testing.T) {
	s := CreateMemoryStorage()
	checkStorageInterface(s)
}
