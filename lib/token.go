package samplify

import "time"

// TokenRequest contains parameters to obtain an access token
type TokenRequest struct {
	ClientID string `json:"clientId" conform:"trim"`
	Username string `json:"username" conform:"trim"`
	Password string `json:"password" conform:"trim"`
}

// TokenResponse stores auth tokens
type TokenResponse struct {
	AccessToken      string `json:"accessToken" conform:"trim"`
	ExpiresIn        uint   `json:"expiresIn"`
	RefreshToken     string `json:"refreshToken" conform:"trim"`
	RefreshExpiresIn uint   `json:"refreshExpiresIn"`
	Acquired         *time.Time
}

// AccessTokenExpired ...
func (t *TokenResponse) AccessTokenExpired() bool {
	if len(t.AccessToken) == 0 || t.Acquired == nil ||
		time.Since(*t.Acquired).Seconds() > float64(t.ExpiresIn) {
		return true
	}
	return false
}

// RefreshTokenExpired ...
func (t *TokenResponse) RefreshTokenExpired() bool {
	if len(t.RefreshToken) == 0 || t.Acquired == nil ||
		time.Since(*t.Acquired).Seconds() > float64(t.RefreshExpiresIn) {
		return true
	}
	return false
}
