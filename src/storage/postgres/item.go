package postgres

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

type ItemRepo struct {
	*Postgres
}

func (p *Postgres) NewItemRepo() *ItemRepo {
	return &ItemRepo{p}
}

func (p *ItemRepo) CreateItem(ctx context.Context, newItem *models.Item) (*models.Item, error) {
	_, err := p.WithContext(ctx).
		Model(newItem).
		Returning("*").
		Insert()

	return newItem, toServiceError(err)
}

func (p *ItemRepo) GetAllItems(ctx context.Context) ([]models.Item, error) {
	var res []models.Item

	err := p.WithContext(ctx).
		Model(&res).
		Select()

	return res, toServiceError(err)
}

func (p *ItemRepo) UpdateItem(ctx context.Context, newItem *models.Item) error {
	_, err := p.WithContext(ctx).
		Model(newItem).
		ExcludeColumn("id").
		Update()

	return toServiceError(err)
}
