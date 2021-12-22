package postgres

import (
	"context"

	"github.com/go-pg/pg/v10/orm"

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

func (p *VehicleRepo) GetAllVehicles(ctx context.Context) ([]models.Vehicle, error) {
	var res []models.Vehicle

	err := p.WithContext(ctx).
		Model(&res).
		Relation("Owner").
		Select()

	return res, toServiceError(err)
}

func (p *VehicleRepo) GetUserVehicles(ctx context.Context, userID int) ([]models.Vehicle, error) {
	var res []models.Vehicle

	err := p.WithContext(ctx).
		Model(&res).
		Relation("Owner", func(query *orm.Query) (*orm.Query, error) {
			return query.Where("id = ?", userID), nil
		}).
		Select()

	return res, toServiceError(err)
}

func (p *VehicleRepo) DeleteVehicle(ctx context.Context, vehicleID int) error {
	_, err := p.WithContext(ctx).
		Model((*models.Vehicle)(nil)).
		Where("id = ?", vehicleID).
		Update()

	return toServiceError(err)
}
