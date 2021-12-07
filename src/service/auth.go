package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Omelman/trucking-api/src/models"
)

var (
	ErrSessionNotFound  = errors.New("session by token ID was not found")
	ErrTokensMismatched = errors.New("access and refresh tokens are mismatched")
	ErrSessionExpired   = errors.New("session expired")
)

func (s *Service) Login(ctx context.Context, loginReq models.LoginRequest) (models.LoginResponse, error) {
	user, err := s.checkUser(ctx, loginReq.Email)
	if err != nil {
		log.Error("login failed:", errors.Wrap(err, "user not found"))

		return models.LoginResponse{}, err
	}

	if !s.CompareHashAndPassword(loginReq.Password, user.Password) {
		log.Error("login failed:", errors.New("wrong password"))

		return models.LoginResponse{}, errors.New("wrong password")
	}

	tkn, err := s.createToken(ctx, user)
	if err != nil {
		log.Error("login failed:", errors.Wrap(err, "failed to create a token"))

		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		TokenPair: models.TokenPair{
			AccessToken:  tkn.Access,
			RefreshToken: tkn.Refresh,
		},
		User: user,
	}, err
}

func (s *Service) Logout(ctx context.Context, accessToken string) (err error) {
	claims, err := s.Revoke(accessToken)
	if err != nil {
		return err
	}

	err = s.authRepo.DisableSessionByID(ctx, claims.SessionID)
	if err != nil {
		return err
	}

	return nil
}

// RefreshToken refreshes access token.
func (s *Service) RefreshToken(ctx context.Context, oldTokens *models.TokenPair) (*models.TokenPair, error) {
	access, err := s.parseJWT(oldTokens.AccessToken)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	accessClaims := s.parseClaims(access)
	if accessClaims == nil {
		return nil, ErrTokenClaimsInvalid
	}

	session, err := s.authRepo.GetSessionByTokenID(ctx, accessClaims.TokenID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	if session.RefreshToken != oldTokens.RefreshToken {
		return nil, ErrTokensMismatched
	}

	now := time.Now().UTC()
	if now.After(*session.ExpiredAt) {
		return nil, ErrSessionExpired
	}

	claims := models.Claims{
		TokenID:   uuid.New(),
		SessionID: session.ID,
		UserID:    session.UserID,
	}

	accessNew, err := s.GenerateAccess(&claims)
	if err != nil {
		return nil, errors.Wrap(err, "token generate ")
	}

	session.TokenID = claims.TokenID
	session.UpdatedAt = &now

	err = s.authRepo.UpdateSession(ctx, session)
	if err != nil {
		return nil, errors.Wrap(err, "update user session ")
	}

	if _, err = s.Revoke(oldTokens.AccessToken); err != nil {
		return nil, err
	}

	return &models.TokenPair{AccessToken: accessNew, RefreshToken: oldTokens.RefreshToken}, nil
}

func (s *Service) checkUser(ctx context.Context, email string) (user models.User, err error) {
	userCandidate, err := s.profileRepo.GetProfilesByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	if userCandidate.ID != 0 {
		return userCandidate, nil
	}

	return models.User{}, errors.New("no user in system")
}

// GenerateAccess generates token with claims.
func (s *Service) GenerateAccess(claims *models.Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Unix() + int64(s.cfg.AccessTokenTTL)
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenWithClaims.SignedString([]byte(s.cfg.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

// GenerateRefresh generates refresh token.
func (s *Service) GenerateRefresh() (string, error) {
	return generateRandomString(s.cfg.RefreshTokenLen)
}

func (s *Service) parseClaims(token *jwt.Token) *models.Claims {
	if claims, ok := token.Claims.(*models.Claims); ok {
		return claims
	}

	return nil
}

// getExpiredAt returns session expiration time till the end of the current day.
func getExpiredAt(userSessionTTLSeconds int) time.Time {
	return time.Now().UTC().Add(time.Second * time.Duration(userSessionTTLSeconds))
}

func (s *Service) createSession(ctx context.Context, userID int) (session *models.UserSession, err error) {
	refreshToken, err := s.GenerateRefresh()
	if err != nil {
		return session, errors.Wrap(err, "generate refresh tokens")
	}

	now := time.Now().UTC()
	expiredAt := getExpiredAt(s.cfg.UserSessionTTL)

	session = &models.UserSession{
		ID:           uuid.New(),
		TokenID:      uuid.New(),
		UserID:       userID,
		RefreshToken: refreshToken,
		CreatedAt:    &now,
		UpdatedAt:    &now,
		ExpiredAt:    &expiredAt,
	}

	return s.authRepo.CreateSession(ctx, session)
}

// Revoke revokes access token.
func (s *Service) Revoke(accessToken string) (*models.Claims, error) {
	token, err := s.parseJWT(accessToken)
	if err != nil {
		return nil, err
	}

	claims := s.parseClaims(token)
	if claims == nil {
		return nil, ErrTokenClaimsInvalid
	}

	return claims, nil
}

// Generate generates token with claims
func (s *Service) Generate(claims *models.Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Unix() + getExpiredAt(s.cfg.AccessTokenTTL).Unix()

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenWithClaims.SignedString([]byte(s.cfg.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) createToken(
	ctx context.Context,
	user models.User,
) (_ *models.Token, err error) {
	session, err := s.createSession(ctx, user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "create session")
	}

	claims := models.Claims{
		SessionID: session.ID,
		TokenID:   session.TokenID,
		UserID:    user.ID,
		UserRole:  user.RoleID,
	}

	jwtToken, err := s.Generate(&claims)
	if err != nil {
		return nil, errors.Wrap(err, "generate jwt token")
	}

	return &models.Token{
		Access:  jwtToken,
		Refresh: session.RefreshToken,
	}, nil
}

// generateRandomString returns a URL-safe, base64 encoded securely generated random string.
func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}
