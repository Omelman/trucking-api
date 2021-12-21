package service

import (
	"context"
	"time"

	"github.com/Omelman/trucking-api/src/models"
)

// CreateVehicle creates new vehicle.
func (s *Service) CreateVehicle(
	ctx context.Context,
	vehicle *models.Vehicle,
) error {
	vehicle.CreatedAt = time.Now()

	_, err := s.vehicleRepo.CreateVehicle(ctx, vehicle)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateVehicle(
	ctx context.Context,
	vehicle *models.Vehicle,
) error {
	vehicle.CreatedAt = time.Now()

	_, err := s.vehicleRepo.CreateVehicle(ctx, vehicle)
	if err != nil {
		return err
	}

	return nil
}
