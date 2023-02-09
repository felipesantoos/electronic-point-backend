package utils

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hashedPassword), nil
}

func PasswordIsValid(originalPassword, givenPassword string) bool {
	decodedPassword, _ := hex.DecodeString(originalPassword)
	err := bcrypt.CompareHashAndPassword(decodedPassword, []byte(givenPassword))
	return err == nil
}
