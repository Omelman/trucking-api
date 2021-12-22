package postgres

import (
	"context"

	"github.com/Omelman/trucking-api/src/models"
)

type ShipmentRepo struct {
	*Postgres
}

func (p *Postgres) NewShipmentRepo() *ShipmentRepo {
	return &ShipmentRepo{p}
}

func (p *ShipmentRepo) CreateShipment(ctx context.Context, newShipment *models.Shipment) (*models.Shipment, error) {
	_, err := p.WithContext(ctx).
		Model(newShipment).
		Returning("*").
		Insert()

	return newShipment, toServiceError(err)
}

func (p *ShipmentRepo) UpdateShipment(ctx context.Context, newShipment *models.Shipment) error {
	_, err := p.WithContext(ctx).
		Model(newShipment).
		Column("status, vehicle_id").
		Update()

	return toServiceError(err)
}
