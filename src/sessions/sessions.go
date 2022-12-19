package sessions

import (
	"fmt"
	"time"

	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/segmentio/ksuid"
)

const (
	SCOPE_SESSION = "sid"
	SCOPE_OTP     = "otp"
	SCOPE_MACHINE = "mac"
)

type SessionExpiredError struct{}

func (m *SessionExpiredError) Error() string {
	return "Session Expired"
}

type SessionIP string
type SessionIPs []SessionIP
type Session struct {
	Id         string          `json:"_id" bson:"_id"`
	Scope      string          `json:"scope" bson:"scope"`
	EntityId   string          `json:"entityId" bson:"entityId"`
	Expire     time.Time       `json:"expire" bson:"expire"`
	Machine    bool            `json:"machine" bson:"machine"`
	AllowedIPs SessionIPs      `json:"allowedIps" bson:"allowedIps"`
	Key        keys.SecuredKey `json:"key" bson:"key"`
}

func ValidScope(scope string) bool {
	return scope != ""
}

// NewSession create a new Session object with the username and duration specified
func NewSession(scope string, username string, duration time.Duration, kp keys.SecureKeyParams) (s Session, err error) {
	if !ValidScope(scope) {
		return s, fmt.Errorf(`Invalid Session Scope "%s"`, scope)
	}
	if !entities.ValidEntityId(username) {
		return s, fmt.Errorf(`Invalid Session Entityname "%s"`, username)
	}
	id := ksuid.New().String()
	key, err := keys.NewSecureKey(id, kp)
	if err != nil {
		return s, err
	}
	expire := time.Now().Add(duration)
	return Session{
		Id:       id,
		Scope:    scope,
		EntityId: username,
		Expire:   expire,
		Key:      *key,
	}, nil
}

// Valid returns a boolean that indicate if the session data is not empty
func (s *Session) Valid() bool {
	return s.Id != "" && s.EntityId != ""
}

// Extend set the time limit for the session from now
func (s *Session) Extend(duration time.Duration) {
	s.Expire = time.Now().Add(duration)
}

// IsExpired returns a boolean that indicate if the time limit of the session is elapsed
func (s *Session) IsExpired() bool {
	now := time.Now()
	return s.Expire.Before(now)
}

// getID returns the identifier of the session
func (s *Session) GetID() string {
	return s.Id
}

// GetEntityname returns the username of the user associated to the session
func (s *Session) GetEntityname() string {
	return s.EntityId
}

func (s *Session) CreateSessionJWT(issuer string, kp keys.SecureKeyParams) (string, error) {
	key, err := s.Key.Unsecure(kp.MasterKey)
	if err != nil {
		return "", err
	}
	return jwtutils.CreateJWT(jwtutils.JWTParams{
		ExpiresAt: s.Expire,
		Issuer:    issuer,
		Secret:    key.Value,
		Scope:     s.Scope,
		Subject:   s.Id,
	})
}

func (s *Session) VerifySessionJWT(jwtString string, kp keys.SecureKeyParams) error {
	key, err := s.Key.Unsecure(kp.MasterKey)
	if err != nil {
		return err
	}
	sessionId, err := jwtutils.ParseJWT(jwtString, jwtutils.JWTParams{
		Secret: key.Value,
		Scope:  s.Scope,
	})
	if err != nil {
		return err
	}
	if sessionId != s.Id {
		return fmt.Errorf("JWT Session ID not match")
	}
	return nil
}
