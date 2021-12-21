package postgres

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

type VehicleRepo struct {
	*Postgres
}

func (p *Postgres) NewVehicleRepo() *VehicleRepo {
	return &VehicleRepo{p}
}

func (p *VehicleRepo) CreateVehicle(ctx context.Context, newVehicle *models.Vehicle) (*models.Vehicle, error) {
	_, err := p.WithContext(ctx).
		Model(newVehicle).
		Returning("*").
		Insert()

	return newVehicle, toServiceError(err)
}

func (p *VehicleRepo) UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {
	_, err := p.WithContext(ctx).
		Model(vehicle).
		ExcludeColumn("id", "created_at").
		Update()

	return vehicle, toServiceError(err)
}
