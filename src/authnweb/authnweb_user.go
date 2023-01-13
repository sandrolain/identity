package authnweb

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

// User represents the user model
type User struct {
	id          string
	name        string
	displayName string
	credentials []webauthn.Credential
}

// NewUser creates and returns a new User
func NewUser(email string, creds []EntityCredential) *User {
	user := &User{}
	user.id = email
	user.name = email
	user.displayName = email

	for _, cred := range creds {
		user.AddCredential(cred.Credential)
	}

	return user
}

// WebAuthnID returns the user's ID
func (u User) WebAuthnID() []byte {
	return []byte(u.id)
}

func (u User) WebAuthnName() string {
	return u.name
}

func (u User) WebAuthnDisplayName() string {
	return u.displayName
}

func (u User) WebAuthnIcon() string {
	return ""
}

// AddCredential associates the credential to the user
func (u *User) AddCredential(cred webauthn.Credential) {
	u.credentials = append(u.credentials, cred)
}

// WebAuthnCredentials returns credentials owned by the user
func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

// CredentialExcludeList returns a CredentialDescriptor array filled with all the user's credentials
func (u User) CredentialExcludeList() []protocol.CredentialDescriptor {
	credentialExcludeList := []protocol.CredentialDescriptor{}
	for _, cred := range u.credentials {
		descriptor := protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		}
		credentialExcludeList = append(credentialExcludeList, descriptor)
	}
	return credentialExcludeList
}
