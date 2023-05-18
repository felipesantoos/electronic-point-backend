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

func (upwd *updatePassword) CurrentPassword() string {
	return upwd.currentPassword
}

func (upwd *updatePassword) NewPassword() string {
	return upwd.newPassword
}
