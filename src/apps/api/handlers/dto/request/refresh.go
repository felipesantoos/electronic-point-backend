package request

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

func (r *RefreshToken) ToDomain() string {
	return r.RefreshToken
}
