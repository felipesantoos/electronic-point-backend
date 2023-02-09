package updatepassword

type UpdatePassword interface {
	CurrentPassword() string
	NewPassword() string
}

type updatePassword struct {
	currentPassword string
	newPassword     string
}

func New(currentPassword, newPassword string) UpdatePassword {
	return &updatePassword{currentPassword, newPassword}
}

func (instance *updatePassword) CurrentPassword() string {
	return instance.currentPassword
}

func (instance *updatePassword) NewPassword() string {
	return instance.newPassword
}
