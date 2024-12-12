package messages

const (
	AccountIDClaim    = "sub"
	ExpirationClaim   = "exp"
	IssuedAtClaim     = "iat"
	TokenTypeClaim    = "typ"
	AccessAudience    = "aud"
	AccountRolesClaim = "roles"
	BearerTokenType   = "Bearer"

	// error messages
	NoAccessAudienceProvidedErr     = "Você precisa fornecer a sessão de acesso do token!"
	InvalidAccessAudienceErr        = "A sessão de acesso fornecida é inválida!"
	InvalidRefreshTokenErrorMessage = "O token de atualização fornecido é inválido!"
)
