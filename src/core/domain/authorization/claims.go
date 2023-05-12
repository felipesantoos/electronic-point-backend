package authorization

import (
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/role"
	"encoding/base64"
	"encoding/json"

	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	jwt.Claims  `json:"c,omitempty"`
	AccountID   string `json:"sub"`
	ProfileID   string `json:"profile_id"`
	RoleEntryID string `json:"role_entry_id,omitempty"`
	Role        string `json:"role"`
	Expiry      int64  `json:"exp"`
	Type        string `json:"typ"`
}

func newClaims(account account.Account, typ string, exp int64) *AuthClaims {
	return &AuthClaims{
		AccountID:   account.ID().String(),
		ProfileID:   account.Person().ID().String(),
		Role:        getStringifiedRoleData(account.Role()),
		RoleEntryID: getRoleEntryIDByAccount(account),
		Type:        typ,
		Expiry:      exp,
	}
}

func getStringifiedRoleData(role role.Role) string {
	var roleData = map[string]interface{}{
		"name": role.Name(),
		"code": role.Code(),
	}
	stringifiedRoleData, _ := json.Marshal(roleData)
	return base64.StdEncoding.EncodeToString(stringifiedRoleData)
}

func getRoleEntryIDByAccount(account account.Account) string {
	var roleEntryID string
	if account.Role().IsProfessional() {
		roleEntryID = account.Professional().ID().String()
	}
	return roleEntryID
}
