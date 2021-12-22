package postgres

import (
	"context"
	"fmt"

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
	fmt.Println(err)

	return newItem, toServiceError(err)
}
