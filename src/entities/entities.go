package entities

import (
	"crypto/rand"
	"fmt"
	"net/mail"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/sandrolain/go-utilities/pkg/cryptoutils"
	"github.com/sandrolain/go-utilities/pkg/pwdutils"
	"github.com/sandrolain/identity/src/config"
	"github.com/sandrolain/identity/src/roles"
)

type EntityType int

const (
	TypeUndefined EntityType = iota
	TypeUser
	TypeMachine
	TypeAdmin
)

type EntityMetadata map[string]string

type Entity struct {
	Id             string      `json:"id" bson:"_id"`
	Type           EntityType  `json:"type" bson:"type"`
	PasswordHash   []byte      `json:"passwordHash" bson:"passwordHash"`
	Validated      bool        `json:"validated" bson:"validated"`
	Suspended      bool        `json:"suspended" bson:"suspended"`
	TotpConfigured bool        `json:"totpConfigured" bson:"totpConfigured"`
	TotpUri        string      `json:"totpUri" bson:"totpUri"`
	RecoveryTokens []string    `json:"recoveryTokens" bson:"recoveryTokens"`
	Roles          roles.Roles `json:"roles" bson:"roles"`
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
		Roles: roles.Roles{},
	}

	if !ValidEntityId(u.Id) {
		err = fmt.Errorf(`invalid entity ID "%s"`, u.Id)
		return
	}

	if err = u.SetPassword(password); err != nil {
		return
	}

	if typ == TypeMachine {
		u.Validated = true
	} else {
		if err = u.ResetTotp(totpConfig); err != nil {
			return
		}
	}

	return
}

func (u *Entity) Validate() {
	u.Validated = true
}

func (u *Entity) Invalidate() {
	u.Validated = false
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

	recTokens, err := generateRecoveryTokens(config.RecoveryTokens)
	if err != nil {
		return err
	}
	u.RecoveryTokens = recTokens

	return nil
}

func (u *Entity) ValidateTotp(code string) (bool, error) {
	key, err := otp.NewKeyFromURL(u.TotpUri)
	if err != nil {
		return false, err
	}
	return totp.ValidateCustom(code, key.Secret(), time.Now(), totp.ValidateOpts{
		Period:    uint(key.Period()),
		Digits:    key.Digits(),
		Algorithm: key.Algorithm(),
	})
}

func (u *Entity) GenerateTotp() (code string, err error) {
	key, err := otp.NewKeyFromURL(u.TotpUri)
	if err != nil {
		return
	}
	return totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{
		Period:    uint(key.Period()),
		Digits:    key.Digits(),
		Algorithm: key.Algorithm(),
	})
}

func (u *Entity) SetTotpConfigured(configured bool) {
	u.TotpConfigured = configured
}

func (u *Entity) IsTotpToConfigure() bool {
	return !u.TotpConfigured
}

func generateRecoveryTokens(cfg config.RecoveryTokensConfig) ([]string, error) {
	tokens := make([]string, cfg.Length)
	for i := 0; i < cfg.Length; i++ {
		b := make([]byte, cfg.Size)
		_, err := rand.Read(b)
		if err != nil {
			return tokens, err
		}
		tokens[i] = fmt.Sprintf("%x", b)
	}
	return tokens, nil
}
