package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Omelman/trucking-api/src/models"
)

type AuthRepo struct {
	*Postgres
}

func (p *Postgres) NewAuthRepo() *AuthRepo {
	return &AuthRepo{p}
}

func (ar *AuthRepo) CreateSession(ctx context.Context, session *models.UserSession) (*models.UserSession, error) {
	_, err := ar.WithContext(ctx).
		Model(session).
		Returning("*").
		Insert()

	return session, err
}

func (ar *AuthRepo) DisableSessionByID(ctx context.Context, sessionID uuid.UUID) error {
	expiredAt := time.Now().UTC()
	session := models.UserSession{
		ID:        sessionID,
		ExpiredAt: &expiredAt,
	}
	_, err := ar.WithContext(ctx).
		Model(&session).
		WherePK().
		UpdateNotZero()

	return err
}

func (ar *AuthRepo) GetSessionByTokenID(ctx context.Context, tokenID uuid.UUID) (*models.UserSession, error) {
	session := &models.UserSession{}
	err := ar.WithContext(ctx).Model(session).
		Where("token_id = ?", tokenID).
		Limit(1).
		Select()

	return session, err
}

func (ar *AuthRepo) UpdateSession(ctx context.Context, userSession *models.UserSession) error {
	_, err := ar.WithContext(ctx).
		Model(userSession).
		WherePK().
		Update()

	return err
}
