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

func (instance *credentials) Email() string {
	return instance.username
}

func (instance *credentials) Password() string {
	return instance.password
}

func (instance *credentials) SetEmail(username string) {
	instance.username = username
}

func (instance *credentials) SetPassword(password string) {
	instance.password = password
}
