package service

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

// CreateItem creates new item.
func (s *Service) CreateItem(
	ctx context.Context,
	item *models.Item,
) error {
	_, err := s.itemRepo.CreateItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

// GetAllItems gets all items.
func (s *Service) GetAllItems(
	ctx context.Context,
) ([]models.Item, error) {
	res, err := s.itemRepo.GetAllItems(ctx)
	if err != nil {
		return []models.Item{}, err
	}

	return res, nil
}

// UpdateItem update new item.
func (s *Service) UpdateItem(
	ctx context.Context,
	item *models.Item,
) error {
	err := s.itemRepo.UpdateItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
