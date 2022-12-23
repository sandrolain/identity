package entities

import (
	"crypto/rand"
	"fmt"
	"net/mail"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/jwtutils"
	"github.com/sandrolain/go-utilities/pkg/pwdutils"
	"github.com/sandrolain/identity/src/config"
)

type EntityType int

const (
	TypeUndefined EntityType = iota
	TypeUser
	TypeMachine
	TypeAdmin
)

type EntityMetadata map[string]string
type EntityRoles []string

type Entity struct {
	Id             string      `json:"id" bson:"_id"`
	Type           EntityType  `json:"type" bson:"type"`
	PasswordHash   []byte      `json:"passwordHash" bson:"passwordHash"`
	Suspended      bool        `json:"suspended" bson:"suspended"`
	TotpConfigured bool        `json:"totpConfigured" bson:"totpConfigured"`
	TotpUri        string      `json:"totpUri" bson:"totpUri"`
	RecoveryTokens []string    `json:"recoveryTokens" bson:"recoveryTokens"`
	Roles          EntityRoles `json:"roles" bson:"roles"`
}

func ValidEntityId(entityId string) bool {
	_, err := mail.ParseAddress(entityId)
	return err == nil
}

type TotpParams struct {
	Enabled bool
	Issuer  string
}

func NewEntity(typ EntityType, entityId string, password string, totpConfig config.TotpConfig) (u Entity, err error) {
	u = Entity{
		Type:  typ,
		Id:    entityId,
		Roles: make(EntityRoles, 0),
	}

	if !ValidEntityId(u.Id) {
		err = fmt.Errorf(`invalid entity ID "%s"`, u.Id)
		return
	}

	if err = u.SetPassword(password); err != nil {
		return
	}

	if typ != TypeMachine {
		if err = u.ResetTotp(totpConfig); err != nil {
			return
		}
	}

	return
}

func (u *Entity) Suspend() {
	u.Suspended = true
}

func (u *Entity) Enable() {
	u.Suspended = false
}

func (u *Entity) IsEnabled() bool {
	return !u.Suspended
}

func (u *Entity) IsUser() bool {
	return u.Type == TypeUser
}

func (u *Entity) IsMachine() bool {
	return u.Type == TypeMachine
}

func (u *Entity) IsAdmin() bool {
	return u.Type == TypeAdmin
}

func (u *Entity) SetPassword(password string) (err error) {
	if password == "" {
		password, err = pwdutils.Generate(16)
	} else if err = pwdutils.Validate(password); err != nil {
		err = fmt.Errorf("invalid password: %v", err)
	}

	if err != nil {
		return
	}

	hashBytes, err := cryptoutils.BcryptHash([]byte(password))
	if err != nil {
		return
	}

	u.PasswordHash = hashBytes

	return
}

func (u *Entity) HasPassword(password string) bool {
	return cryptoutils.BcryptCompare([]byte(password), u.PasswordHash)
}

func (u *Entity) ResetTotp(config config.TotpConfig) error {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.Issuer,
		AccountName: u.Id,
	})
	if err != nil {
		return err
	}

	u.TotpConfigured = false
	u.TotpUri = key.URL()
	u.RecoveryTokens = generateRecoveryTokens(config.RecoveryTokens)

	return nil
}

func (u *Entity) ValidateTotp(code string) (bool, error) {
	key, err := otp.NewKeyFromURL(u.TotpUri)
	if err != nil {
		return false, err
	}
	return totp.Validate(code, key.Secret()), nil
}

func (u *Entity) GenerateTotp() (code string, err error) {
	key, err := otp.NewKeyFromURL(u.TotpUri)
	if err != nil {
		return
	}
	return totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{})
}

func (u *Entity) SetTotpConfigured(configured bool) {
	u.TotpConfigured = configured
}

func (u *Entity) IsTotpToConfigure() bool {
	return !u.TotpConfigured
}

func (u *Entity) CreateTotpJWT(dur time.Duration, issuer string, secret []byte) (string, error) {
	return jwtutils.CreateJWT(jwtutils.JWTParams{
		ExpiresAt: time.Now().Add(dur),
		Issuer:    issuer,
		Secret:    secret,
		Scope:     "otp",
		Subject:   u.Id,
	})
}

func ParseTotpJWT(jwtString string, secret []byte) (string, error) {
	return jwtutils.ParseJWT(jwtString, jwtutils.JWTParams{
		Secret: secret,
		Scope:  "otp",
	})
}

func generateRecoveryTokens(cfg config.RecoveryTokensConfig) []string {
	tokens := make([]string, cfg.Length)
	for i := 0; i < cfg.Length; i++ {
		b := make([]byte, cfg.Size)
		rand.Read(b)
		tokens[i] = fmt.Sprintf("%x", b)
	}
	return tokens
}
