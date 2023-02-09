package authorization

import (
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	jwt.Claims `json:"c,emitempty"`
	AccountID  string `json:"sub"`
	RoleCode   string `json:"section"`
	Expiry     int64  `json:"exp"`
	Type       string `json:"typ"`
}

func newClaims(accountId string, roleCode string, typ string, exp int64) *AuthClaims {
	return &AuthClaims{
		AccountID: accountId,
		RoleCode:  roleCode,
		Type:      typ,
		Expiry:    exp,
	}
}
