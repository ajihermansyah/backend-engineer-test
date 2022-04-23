package model

type TokenClaims struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
}
