package redisstorage

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/sandrolain/go-utilities/pkg/testredisutils"
	"github.com/sandrolain/identity/src/storage"
)

const (
	testPassword = "development.password"
)

func checkStorageInterface(s storage.VolatileStorage) {
	return
}

func TestStorageInterface(t *testing.T) {
	redisMock := testredisutils.NewMockServer(t, testPassword)

	s, _ := CreateRedisStorage(redisMock.Addr(), testPassword, &tls.Config{}, time.Second)

	checkStorageInterface(s)
}
