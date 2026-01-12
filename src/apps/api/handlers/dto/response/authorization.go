package response

import "eletronic_point/src/core/domain/authorization"

type Authorization struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type authorizationBuilder struct{}

func NewAuthorizationBuilder() *authorizationBuilder {
	return &authorizationBuilder{}
}

func (*authorizationBuilder) BuildFromDomain(data authorization.Authorization) *Authorization {
	return &Authorization{
		Token: data.Token(),
	}
}

func (*authorizationBuilder) BuildFromTokens(accessToken, refreshToken authorization.Authorization) *Authorization {
	return &Authorization{
		Token:        accessToken.Token(),
		RefreshToken: refreshToken.Token(),
	}
}
