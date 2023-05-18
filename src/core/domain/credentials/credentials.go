package credentials

type Credentials interface {
	Email() string
	Password() string
	SetEmail(string)
	SetPassword(string)
}

type credentials struct {
	username string
	password string
}

func New() Credentials {
	return &credentials{}
}

func (c *credentials) Email() string {
	return c.username
}

func (c *credentials) Password() string {
	return c.password
}

func (c *credentials) SetEmail(username string) {
	c.username = username
}

func (c *credentials) SetPassword(password string) {
	c.password = password
}
