package service

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

func (s *Service) GetStatistics(
	ctx context.Context,
) (models.Statistics, error) {
	vehicles, err := s.vehicleRepo.GetAllVehicles(ctx)
	if err != nil {
		return models.Statistics{}, err
	}

	items, err := s.itemRepo.GetAllItems(ctx)
	if err != nil {
		return models.Statistics{}, err
	}

	var (
		sumItemQuantity    int
		sumItemVolume      int
		sumVehicleVolume   int
		sumVehicleCapacity int
	)

	for i := range items {
		sumItemVolume += items[i].Volume
		sumItemQuantity += items[i].Quantity
	}

	for i := range vehicles {
		sumVehicleVolume += vehicles[i].UsefulVolume
		sumVehicleCapacity += vehicles[i].CarryingCapacity
	}

	resp := models.Statistics{
		AverageItemsQuantity:   float32(sumItemQuantity) / float32(len(items)),
		AverageItemsVolume:     float32(sumItemVolume) / float32(len(items)),
		AverageVehicleVolume:   float32(sumVehicleVolume) / float32(len(vehicles)),
		AverageVehicleCapacity: float32(sumVehicleCapacity) / float32(len(vehicles)),
	}

	return resp, nil
}
