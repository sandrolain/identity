package api

import (
	"testing"
	"time"

	"github.com/sandrolain/go-utilities/pkg/crudutils"
	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/go-utilities/pkg/pwdutils"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/sandrolain/identity/src/roles"
	"github.com/sandrolain/identity/src/storage/memorystorage"
)

func TestAdminMachineManagement(t *testing.T) {
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
	u.Roles.Add(roles.RoleMachinesManager)
	err = storage.SaveEntity(u)
	if err != nil {
		t.Fatal(err)
	}

	a := API{
		Config:            cfg,
		VolatileStorage:   storage,
		PersistentStorage: storage,
	}

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

	code, err = u.GenerateTotp()
	if err != nil {
		t.Fatal(err)
	}

	res2a, err := a.VerifyEntityValidation(entities.TypeAdmin, res2.ValidationToken)
	if err != nil {
		t.Fatal(err)
	}

	res2b, err := a.CompleteEntityValidation(entities.TypeAdmin, res2a.TotpToken, code)
	if err != nil {
		t.Fatal(err)
	}

	res3, err := a.CreateMachine(res2b.SessionToken, "machine@sandrolain.com", []string{})
	if err != nil {
		t.Fatal(err)
	}

	res4, err := a.InitMachineSession(res2b.SessionToken, res3.MachineId, []string{"192.168.1.0/25", "127.0.0.1"})
	if err != nil {
		t.Fatal(err)
	}

	expire, err := time.Parse(time.RFC3339, res4.Expire)
	if err != nil {
		t.Fatal(err)
	}
	jwt, err := jwtutils.CreateJWT(jwtutils.JWTParams{
		Subject:   res4.Subject,
		Secret:    res4.Secret,
		Issuer:    "test.com",
		ExpiresAt: expire,
	})
	if err != nil {
		t.Fatal(err)
	}

	res5, err := a.AuthenticateMachine(jwt, res4.MachineId, "192.168.1.127")
	if err != nil {
		t.Fatal(err)
	}

	if res5.MachineId != res3.MachineId {
		t.Fatalf("machine id not match %v != %v", res5.MachineId, res3.MachineId)
	}

	res5, err = a.AuthenticateMachine(jwt, res4.MachineId, "127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	res5, err = a.AuthenticateMachine(jwt, res4.MachineId, "192.168.1.128")
	if err == nil || !crudutils.IsNotAuthorized(err) {
		t.Fatalf("expected ip not authorized, instead: %v", err)
	}
}
