package credentials

type Credentials interface {
	Email() string
	Password() string
}

type credentials struct {
	username string
	password string
}

func New(username, password string) Credentials {
	return &credentials{username, password}
}

func (c *credentials) Email() string {
	return c.username
}

func (c *credentials) Password() string {
	return c.password
}
