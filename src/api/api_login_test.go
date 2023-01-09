package api

import (
	"testing"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/pwdutils"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/storage/memorystorage"
)

func TestUserLogin(t *testing.T) {
	key, err := cryptoutils.RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.GetDefaultConfiguration()
	cfg.Keys.MasterKey = keys.MasterKeyFromBytes(key)

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

		u, err = storage.GetEntity(u.Id)
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

		if res2.SessionToken != "" {
			t.Fatal("Session token should be empty for not validate users")
		}

		if res2.ValidationToken == "" {
			t.Fatal("Validation token should NOT be empty for not validate users")
		}

		res2a, err := a.CompleteEntityValidation(res2.ValidationToken)
		if err != nil {
			t.Fatal(err)
		}

		res3, err := a.GetUserDetails(res2a.SessionToken)
		if err != nil {
			t.Fatal(err)
		}

		if res3.EntityId != u.Id {
			t.Fatalf("entity details not match: %v != %v", res3.EntityId, u.Id)
		}

		res4, err := a.Logout(res2a.SessionToken)
		if err != nil {
			t.Fatal(err)
		}

		_, err = storage.GetSession(res4.SessionId)
		if _, ok := err.(*crudutils.NotFoundError); !ok {
			t.Fatal("session should not be available")
		}

		_, err = a.GetUserDetails(res2a.SessionToken)
		if _, ok := err.(*crudutils.NotFoundError); !ok {
			t.Fatal("API call with expired session should not be possible")
		}
	}

	{
		_, err := a.Login(entities.TypeAdmin, entityId, password)
		if err == nil {
			t.Fatal("Normal user should not login as admin")
		}
	}
}

func TestAdminLogin(t *testing.T) {
	key, err := cryptoutils.RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.GetDefaultConfiguration()
	cfg.Keys.MasterKey = keys.MasterKeyFromBytes(key)

	storage := memorystorage.CreateMemoryStorage()

	entityId := "user1@sandrolain.com"
	password, err := pwdutils.Generate(16)
	if err != nil {
		t.Fatal(err)
	}

	u, err := entities.NewEntity(entities.TypeAdmin, entityId, password, cfg.Totp)
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
		res, err := a.Login(entities.TypeAdmin, entityId, password)
		if err != nil {
			t.Fatal(err)
		}

		u, err = storage.GetEntity(u.Id)
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

		if res2.SessionToken != "" {
			t.Fatal("Session token should be empty for not validate users")
		}

		if res2.ValidationToken == "" {
			t.Fatal("Validation token should NOT be empty for not validate users")
		}

		res2a, err := a.CompleteEntityValidation(res2.ValidationToken)
		if err != nil {
			t.Fatal(err)
		}

		res3, err := a.GetUserDetails(res2a.SessionToken)
		if err != nil {
			t.Fatal(err)
		}

		if res3.EntityId != u.Id {
			t.Fatalf("entity details not match: %v != %v", res3.EntityId, u.Id)
		}
	}

	{
		_, err := a.Login(entities.TypeUser, entityId, password)
		if err == nil {
			t.Fatal("Admin user should not login as normal user")
		}
	}
}
