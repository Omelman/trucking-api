package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserSession struct {
	ID           uuid.UUID  `json:"id"`
	TokenID      uuid.UUID  `json:"token_id"`
	UserID       int        `json:"user_id"`
	RefreshToken string     `json:"refresh_token"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	ExpiredAt    *time.Time `json:"expired_at"`

	User *User `json:"-"`
}

func (us UserSession) GetUserRole() string {
	if us.User == nil {
		return ""
	}

	return string(us.User.Role)
}

// Claims token's claims/payload.
type Claims struct {
	SessionID uuid.UUID
	TokenID   uuid.UUID
	UserID    int
	UserRole  UserRole

	jwt.StandardClaims
}

// TTL returns TTL in seconds.
func (c *Claims) TTL() int64 {
	return c.StandardClaims.ExpiresAt - time.Now().Unix()
}

// swagger:model
type RegistrationRequest struct {
	FirstName string   `json:"first_name" validate:"gte=2,lte=70,required"`
	LastName  string   `json:"last_name" validate:"gte=2,lte=70,required"`
	Email     string   `json:"email" validate:"email,required"`
	Password  string   `json:"password" validate:"required,gte=8,lte=32"`
	Role      UserRole `json:"role" validate:"required"`
}

type RegistrationResponse struct {
	UserID         int       `json:"user_id"`
	ConfirmationID uuid.UUID `json:"confirmation_id"`
}
