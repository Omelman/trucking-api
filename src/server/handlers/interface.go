package handlers

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

type AuthService interface {
	Login(ctx context.Context, loginReq models.LoginRequest) (resp models.LoginResponse, err error)
	Logout(ctx context.Context, accessToken string) (err error)
	RefreshToken(ctx context.Context, oldTokens *models.TokenPair) (*models.TokenPair, error)
	CreateUser(ctx context.Context, user *models.User) error
}
