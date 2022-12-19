package api

import (
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/storage"
)

type API struct {
	Config            config.Config
	VolatileStorage   storage.VolatileStorage
	PersistentStorage storage.PersistentStorage
}
