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
