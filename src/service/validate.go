package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"

	"github.com/Omelman/trucking-api/src/models"
)

// Token error list.
var (
	ErrTokenInvalid       = errors.New("token is invalid")
	ErrTokenInBlackList   = errors.New("token in black list")
	ErrTokenClaimsInvalid = errors.New("token claims are invalid")
)

// Validate validates access token.
func (s *Service) Validate(accessToken string) (*models.UserSession, error) {
	token, err := s.parseJWT(accessToken)
	if err != nil {
		return nil, err
	}

	claims := s.parseClaims(token)
	if claims == nil || !token.Valid {
		return nil, ErrTokenInvalid
	}

	return &models.UserSession{
		UserID:  claims.UserID,
		TokenID: claims.TokenID,
		User: &models.User{
			Role: claims.UserRole,
		},
	}, nil
}

func (s *Service) ValidateExternalAPIToken(externalToken string) error {
	if externalToken != s.cfg.ExternalToken {
		return errors.New("External token is invalid")
	}

	return nil
}

func (s *Service) parseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.AccessTokenSecret), nil
	})

	if token == nil {
		return nil, ErrTokenInvalid
	}

	//nolint: errorlint
	switch ve := err.(type) {
	case *jwt.ValidationError:
		if ve.Errors|(jwt.ValidationErrorExpired) != jwt.ValidationErrorExpired {
			return nil, ErrTokenInvalid
		}
	case nil:
	default:
		return nil, err
	}

	return token, nil
}
