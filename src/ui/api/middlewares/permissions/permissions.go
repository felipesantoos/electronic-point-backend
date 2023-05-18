package permissions

import (
	"backend_template/src/core/domain/role"
)

type PermissionsHelper interface {
	AuthPolicies() []Policy
}

type permissionsHelper struct{}

var anonymousEntries = []string{
	`\/api\/auth\/login`,
	`\/api\/auth\/reset-password`,
	`\/api\/auth\/reset-password\/[^\r\n]*`,
	`\/api\/docs\/[^\r\n]*`,
}
var adminEntries = []string{
	`\/[^\r\n]*`,
}
var professionalEntries = []string{
	`\/api\/accounts\/[^\r\n]*`,
	`\/api\/auth/logout`,
	`\/api\/auth/reset-password`,
	`\/api\/auth/update-password[^\r\n]*`,
	`\/api\/res\/[^\r\n]*`,
	`\/api\/professional\/[^\r\n]*`,
}
var allowAll = "*"

func New() PermissionsHelper {
	return &permissionsHelper{}
}

func (*permissionsHelper) AuthPolicies() []Policy {
	policies := []Policy{}
	entries := []Entry{
		NewEntry(role.ANONYMOUS_ROLE_CODE, anonymousEntries),
		NewEntry(role.ADMIN_ROLE_CODE, adminEntries),
		NewEntry(role.PROFESSIONAL_ROLE_CODE, professionalEntries),
	}
	for _, entry := range entries {
		for _, obj := range entry.Objects() {
			policies = append(policies, NewPolicy(entry.Subject(), obj, allowAll))
		}
	}

	return policies
}
