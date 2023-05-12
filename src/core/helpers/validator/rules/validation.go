package rules

import "net/mail"

func ValidateEmailRule() *Rule {
	return &Rule{
		validator:   validateEmailFunc,
		Type:        EMAIL_VALIDATION,
		Description: "verify if a value is a valid email",
	}
}

func validateEmailFunc(value interface{}) bool {
	if email, ok := value.(string); !ok {
		return false
	} else if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}
