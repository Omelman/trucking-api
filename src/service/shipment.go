package service

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

// CreateShipment creates new shipment.
func (s *Service) CreateShipment(
	ctx context.Context,
	shipment *models.Shipment,
) (*models.Shipment, error) {
	resp, err := s.shipmentRepo.CreateShipment(ctx, shipment)
	if err != nil {
		return &models.Shipment{}, err
	}

	return resp, nil
}

// UpdateShipment updates new shipment.
func (s *Service) UpdateShipment(
	ctx context.Context,
	shipment *models.Shipment,
) error {
	err := s.shipmentRepo.UpdateShipment(ctx, shipment)
	if err != nil {
		return err
	}

	return nil
}
