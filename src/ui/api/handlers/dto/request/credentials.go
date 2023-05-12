package request

import (
	"backend_template/src/core/domain/credentials"
	"backend_template/src/core/domain/errors"
	"fmt"
	"net/mail"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type credentialsBuilder struct{}

func CredentialsBuilder() *credentialsBuilder {
	return &credentialsBuilder{}
}

func (*credentialsBuilder) FromBody(data map[string]interface{}) (*Credentials, errors.Error) {
	email := fmt.Sprint(data["email"])
	if addr, _ := mail.ParseAddress(email); addr == nil {
		return nil, errors.NewFromString("you must provide a valid email!")
	}
	return &Credentials{
		Email:    fmt.Sprint(data["email"]),
		Password: fmt.Sprint(data["password"]),
	}, nil
}

func (instance *Credentials) ToDomain() credentials.Credentials {
	credentials := credentials.New()
	credentials.SetEmail(instance.Email)
	credentials.SetPassword(instance.Password)
	return credentials
}
