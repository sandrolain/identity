package api

import (
	"fmt"
	"testing"

	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/pwdutils"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/storage/memorystorage"
)

func TestLogin(t *testing.T) {
	cfg := config.GetDefaultConfiguration()
	cfg.SecureKey.MasterKey = cryptoutils.RandomBytes(32)

	storage := memorystorage.CreateMemoryStorage()

	entityId := "user1@sandrolain.com"
	password, err := pwdutils.Generate(16)
	if err != nil {
		t.Fatal(err)
	}

	u, err := entities.NewEntity(entities.TypeUser, entityId, password, cfg.Totp)
	if err != nil {
		t.Fatal(err)
	}
	err = storage.SaveEntity(u)
	if err != nil {
		t.Fatal(err)
	}

	a := API{
		Config:            cfg,
		VolatileStorage:   storage,
		PersistentStorage: storage,
	}

	{
		res, err := a.Login(entities.TypeUser, entityId, password)
		if err != nil {
			t.Fatal(err)
		}

		code, err := u.GenerateTotp()
		if err != nil {
			t.Fatal(err)
		}

		res2, err := a.LoginTotp(res.TotpToken, code)
		if err != nil {
			t.Fatal(err)
		}

		res3, err := a.GetEntityDetails(res2.SessionToken)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("res3: %v\n", res3)
	}

	{
		_, err := a.Login(entities.TypeAdmin, entityId, password)
		if err == nil {
			t.Fatal("Normal user should not login as admin")
		}
	}
}
