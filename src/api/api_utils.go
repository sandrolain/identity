package api

import "github.com/sandrolain/identity/src/entities"

type ValidateEmailResponse struct {
	Valid bool
}
type ValidatePasswordResponse struct {
	Valid bool
}

func (a *API) ValidateEmail(email string) (res ValidateEmailResponse, err error) {
	res.Valid = entities.ValidEntityId(email)
	return
}

func (a *API) ValidatePassword(password string) (res ValidatePasswordResponse, err error) {
	res.Valid = entities.ValidPassword(password)
	return
}
