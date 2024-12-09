package validator

import (
	"strconv"
	"strings"
)

func IsRegistrationValid(registration string) bool {
	registration = strings.TrimSpace(registration)
	if convReg, err := strconv.Atoi(registration); err != nil {
		return false
	} else if convReg <= 0 {
		return false
	}

	return len(registration) == 10
}
