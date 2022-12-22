package memorystorage

import (
	"testing"

	"github.com/sandrolain/identity/src/storage"
)

func checkVolatileStorageInterface(s storage.VolatileStorage) {
	s.GetSession("hello")
}

func checkPersistentStorageInterface(s storage.PersistentStorage) {
	s.GetSession("hello")
}

func TestStorageInterface(t *testing.T) {
	s := CreateMemoryStorage()
	checkVolatileStorageInterface(s)
	checkPersistentStorageInterface(s)
}
