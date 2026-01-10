package request

import (
	"eletronic_point/src/core/domain/credentials"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) ToDomain() credentials.Credentials {
	return credentials.New(c.Email, c.Password)
}
