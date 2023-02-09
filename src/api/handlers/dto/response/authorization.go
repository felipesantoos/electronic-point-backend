package response

import "dit_backend/src/core/domain/authorization"

type Authorization struct {
	Token string `json:"access_token"`
}

type authorizationBuilder struct{}

func AuthorizationBuilder() *authorizationBuilder {
	return &authorizationBuilder{}
}

func (*authorizationBuilder) FromDomain(auth authorization.Authorization) *Authorization {
	return &Authorization{
		Token: auth.Token(),
	}
}
