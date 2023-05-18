package request

import (
	"backend_template/src/core/domain/credentials"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) ToDomain() credentials.Credentials {
	credentials := credentials.New()
	credentials.SetEmail(c.Email)
	credentials.SetPassword(c.Password)
	return credentials
}
