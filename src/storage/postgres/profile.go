package postgres

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

type ProfileRepo struct {
	*Postgres
}

func (p *Postgres) NewProfileRepo() *ProfileRepo {
	return &ProfileRepo{p}
}

func (p *ProfileRepo) GetProfilesByEmail(ctx context.Context, email string) (models.User, error) {
	var res models.User

	err := p.WithContext(ctx).
		Model(&res).
		Where("?TableAlias.email = ?", email).
		Select()

	return res, toServiceError(err)
}

func (p *ProfileRepo) AddNewProfile(ctx context.Context, newUser *models.User) (*models.User, error) {
	_, err := p.WithContext(ctx).
		Model(newUser).
		Insert()

	return newUser, toServiceError(err)
}
