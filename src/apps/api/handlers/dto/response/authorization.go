package response

import "eletronic_point/src/core/domain/authorization"

type Authorization struct {
	Token string `json:"access_token"`
}

type authorizationBuilder struct{}

func NewAuthorizationBuilder() *authorizationBuilder {
	return &authorizationBuilder{}
}

func (*authorizationBuilder) BuildFromDomain(data authorization.Authorization) *Authorization {
	return &Authorization{
		data.Token(),
	}
}
