package models

// Token describes access and refresh tokens.
type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}
