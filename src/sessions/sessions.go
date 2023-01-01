package sessions

import (
	"fmt"
	"strings"
	"time"

	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/identity/src/entities"
	"github.com/sandrolain/identity/src/keys"
	"github.com/segmentio/ksuid"
)

type SessionScope string

const (
	ScopeTotp     SessionScope = "otp"
	ScopeLogin    SessionScope = "sid"
	ScopePassword SessionScope = "pwd"
	ScopeWebauthn SessionScope = "wau"
	ScopeMachine  SessionScope = "mac"
)

type SessionExpiredError struct{}

func (m *SessionExpiredError) Error() string {
	return "Session Expired"
}

type SessionIP = string
type SessionIPs []SessionIP
type Session struct {
	Id         string          `json:"_id" bson:"_id"`
	Scope      SessionScope    `json:"scope" bson:"scope"`
	EntityId   string          `json:"entityId" bson:"entityId"`
	Expire     time.Time       `json:"expire" bson:"expire"`
	AllowedIPs SessionIPs      `json:"allowedIps" bson:"allowedIps"`
	Key        keys.SecuredKey `json:"key" bson:"key"`
}

func ValidScope(scope SessionScope) bool {
	return scope != ""
}

// NewSession create a new Session object with the username and duration specified
func NewSession(scope SessionScope, username string, duration time.Duration, allowedIps []string, mk keys.MasterKey) (s Session, err error) {
	if !ValidScope(scope) {
		return s, fmt.Errorf(`Invalid Session Scope "%s"`, scope)
	}
	if !entities.ValidEntityId(username) {
		return s, fmt.Errorf(`Invalid Session Entityname "%s"`, username)
	}
	id := ksuid.New().String()
	key, err := keys.NewSecureKey(32, mk)
	if err != nil {
		return s, err
	}
	expire := time.Now().Add(duration)
	return Session{
		Id:         id,
		Scope:      scope,
		EntityId:   username,
		Expire:     expire,
		AllowedIPs: SessionIPs(allowedIps),
		Key:        *key,
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
	return s.Expire.Before(time.Now())
}

func (s *Session) GetTtl() time.Duration {
	return s.Expire.Sub(time.Now())
}

// getID returns the identifier of the session
func (s *Session) GetID() string {
	return s.Id
}

// GetEntityname returns the username of the user associated to the session
func (s *Session) GetEntityname() string {
	return s.EntityId
}

func (s *Session) GetJwtSubject() string {
	return fmt.Sprintf("%s:%s", s.Scope, s.Id)
}

func (s *Session) CreateSessionJWT(issuer string, mk keys.MasterKey) (string, error) {
	key, err := s.Key.Unsecure(mk)
	if err != nil {
		return "", err
	}
	return jwtutils.CreateJWT(jwtutils.JWTParams{
		ExpiresAt: s.Expire,
		Issuer:    issuer,
		Secret:    key.Value,
		Subject:   s.GetJwtSubject(),
	})
}

func (s *Session) VerifySessionJWT(jwtString string, mk keys.MasterKey) error {
	key, err := s.Key.Unsecure(mk)
	if err != nil {
		return err
	}
	subject, err := jwtutils.ParseJWT(jwtString, jwtutils.JWTParams{
		Secret: key.Value,
	})
	if err != nil {
		return err
	}
	subj, err := ParseScopeSubject(subject)
	if err != nil {
		return err
	}
	if string(s.Scope) != subj.Scope {
		return fmt.Errorf(`JWT scope "%s" not as expected "%s"`, subj.Scope, s.Scope)
	}
	if s.Id != subj.SessionId {
		return fmt.Errorf(`JWT Session ID "%s" not match "%s"`, subj.SessionId, s.Id)
	}
	return nil
}

type ScopeSubject struct {
	Scope     string
	SessionId string
}

func ParseScopeSubject(subject string) (res ScopeSubject, err error) {
	i := strings.Index(subject, ":")
	if i < 0 {
		err = fmt.Errorf(`invalid JWT subject "%s"`, subject)
		return
	}
	res.Scope = subject[:i]
	res.SessionId = subject[i+1:]
	return
}
