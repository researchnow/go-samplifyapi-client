package samplify

import "time"

// TokenRequest contains parameters to obtain an access token
type TokenRequest struct {
	ClientID string `json:"clientId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// TokenResponse stores auth tokens
type TokenResponse struct {
	AccessToken      string `json:"accessToken"`
	ExpiresIn        uint   `json:"expiresIn"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn uint   `json:"refreshExpiresIn"`
	Acquired         *time.Time
}

// AccessTokenExpired ...
func (t *TokenResponse) AccessTokenExpired() bool {
	if t.Acquired == nil ||
		time.Since(*t.Acquired).Seconds() > float64(t.ExpiresIn) {
		return true
	}
	return false
}
