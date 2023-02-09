package permissions

import "dit_backend/src/core/domain/authorization"

type PermissionsHelper interface {
	AuthMatcherTemplate() string
	AuthPolicies() []Policy
	AuthCasbinPolicies() []map[string]string
}

type permissionsHelper struct{}

var casbinModelTemplate = `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && regexMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

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
	`\/api\/auth/reset-password`,
	`\/api\/auth/update-password[^\r\n]*`,
	`\/api\/res\/[^\r\n]*`,
	`\/api\/professional\/[^\r\n]*`,
}
var allowAll = "*"

func New() PermissionsHelper {
	return &permissionsHelper{}
}

func (*permissionsHelper) AuthMatcherTemplate() string {
	return casbinModelTemplate
}

func (*permissionsHelper) AuthPolicies() []Policy {
	policies := []Policy{}
	entries := []Entry{
		NewEntry(authorization.ANONYMOUS_ROLE_CODE, anonymousEntries),
		NewEntry(authorization.ADMIN_ROLE_CODE, adminEntries),
		NewEntry(authorization.PROFESSIONAL_ROLE_CODE, professionalEntries),
	}
	for _, entry := range entries {
		for _, obj := range entry.Objects() {
			policies = append(policies, NewPolicy(entry.Subject(), obj, allowAll))
		}
	}

	return policies
}

func (instance *permissionsHelper) AuthCasbinPolicies() []map[string]string {
	authPolicies := instance.AuthPolicies()
	policies := []map[string]string{}
	for _, policy := range authPolicies {
		policies = append(policies, map[string]string{
			"PType": "p",
			"V0":    policy.Subject(),
			"V1":    policy.Object(),
			"V2":    policy.Action(),
		})
	}
	return policies
}
