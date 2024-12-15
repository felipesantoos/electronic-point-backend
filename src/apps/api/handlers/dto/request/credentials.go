package request

import (
	"eletronic_point/src/core/domain/credentials"
)

type Credentials struct {
	Email    string `json:"email" example:"jose@gmail.com"`
	Password string `json:"password" example:"123456"`
}

func (c *Credentials) ToDomain() credentials.Credentials {
	return credentials.New(c.Email, c.Password)
}
