package sessions

import (
	"testing"
	"time"

	"github.com/sandrolain/identity/src/keys"
)

func getKeyParams() keys.SecureKeyParams {
	return keys.SecureKeyParams{
		Length:    32,
		MasterKey: keys.MasterKeyFromBytes(keys.GenerateKeyValue(32)),
	}
}

func TestSessionCreate(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	sess, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	if sess.EntityId != username {
		t.Fatalf(`Expect "%s" to be equal to "%s"`, sess.EntityId, username)
	}

	limit := time.Now().Add(duration)
	if sess.Expire.UnixMilli() > limit.UnixMilli() {
		t.Fatalf(`Expect "%s" to be less or equal to to "%s"`, sess.Expire, limit)
	}
}

func TestSessionIsValid(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	sess, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	if !sess.Valid() {
		t.Fatalf(`Expect %+v to be valid`, sess)
	}

	sess = Session{}

	if sess.Valid() {
		t.Fatalf(`Expect %+v to be not valid`, sess)
	}
}

func TestSessionExtend(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	sess, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	limit := time.Now().Add(duration)
	if sess.Expire.Unix() != limit.Unix() {
		t.Fatalf(`Expect "%s" to be less or equal to to "%s"`, sess.Expire, limit)
	}

	duration, _ = time.ParseDuration("45m")
	sess.Extend(duration)
	limit = time.Now().Add(duration)
	if sess.Expire.Unix() != limit.Unix() {
		t.Fatalf(`Expect "%s" to be less or equal to to "%s"`, sess.Expire, limit)
	}
}

func TestSessionIsExpired(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	sess, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	if sess.IsExpired() {
		t.Fatalf(`Expect "%s" to be not expired`, sess.Expire)
	}

	duration, _ = time.ParseDuration("-45m")
	sess.Extend(duration)
	if !sess.IsExpired() {
		t.Fatalf(`Expect "%s" to be expired`, sess.Expire)
	}
}

func TestSessionGetID(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	sess, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	if sess.GetID() == "" {
		t.Fatalf(`Expect "%s" to be not empty`, sess.GetID())
	}

	sess2, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	if sess.GetID() == sess2.GetID() {
		t.Fatalf(`Expect "%s" to be not equal to "%s"`, sess.GetID(), sess2.GetID())
	}
}

func TestSessionGetEntityname(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	sess, err := NewSession(scope, username, duration, getKeyParams())

	if err != nil {
		t.Fatal(err)
	}

	if sess.GetEntityname() != username {
		t.Fatalf(`Expect "%s" to be equal to "%s"`, sess.GetEntityname(), username)
	}
}

func TestSessionJWT(t *testing.T) {
	username := "john"
	scope := ScopeLogin
	duration, _ := time.ParseDuration("30m")
	kp := getKeyParams()
	sess, err := NewSession(scope, username, duration, kp)

	if err != nil {
		t.Fatal(err)
	}

	jwt, err := sess.CreateSessionJWT("issuer.com", kp.MasterKey)
	if err != nil {
		t.Fatal(err)
	}

	err = sess.VerifySessionJWT(jwt, kp)
	if err != nil {
		t.Fatal(err)
	}

}
