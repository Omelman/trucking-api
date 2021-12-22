package service

import (
	"context"
	"errors"
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

// UpdateVehicle updates vehicle.
func (s *Service) UpdateVehicle(
	ctx context.Context,
	vehicle *models.Vehicle,
) error {
	_, err := s.vehicleRepo.CreateVehicle(ctx, vehicle)
	if err != nil {
		return err
	}

	return nil
}

// GetAllVehicles get all vehicles.
func (s *Service) GetAllVehicles(
	ctx context.Context,
) ([]models.Vehicle, error) {
	resp, err := s.vehicleRepo.GetAllVehicles(ctx)
	if err != nil {
		return []models.Vehicle{}, err
	}

	return resp, nil
}

// GetUserVehicles get all vehicles for user.
func (s *Service) GetUserVehicles(
	ctx context.Context,
	userID int,
) ([]models.Vehicle, error) {
	resp, err := s.vehicleRepo.GetUserVehicles(ctx, userID)
	if err != nil {
		return []models.Vehicle{}, err
	}

	return resp, nil
}

// DeleteUserVehicles delete user vehicle.
func (s *Service) DeleteUserVehicles(
	ctx context.Context,
	vehicleID int,
	userID int,
) error {
	vehicles, err := s.vehicleRepo.GetUserVehicles(ctx, userID)
	if err != nil {
		return err
	}

	allowed := false
	for i := range vehicles {
		if vehicles[i].OwnerID == userID {
			allowed = true
			break
		}
	}

	if !allowed {
		return errors.New("not allowed to delete vehicle")
	}

	err = s.vehicleRepo.DeleteVehicle(ctx, vehicleID)
	if err != nil {
		return err
	}

	return nil
}
