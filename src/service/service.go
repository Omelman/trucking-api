package service

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"github.com/Omelman/trucking-api/src/config"
	"github.com/Omelman/trucking-api/src/models"
)

var (
	service *Service
	once    sync.Once
)

type Service struct {
	cfg          *config.Config
	authRepo     AuthRepo
	profileRepo  ProfileRepo
	vehicleRepo  VehicleRepo
	itemRepo     ItemRepo
	shipmentRepo ShipmentRepo
}

func New(
	cfg *config.Config,
	aur AuthRepo,
	pr ProfileRepo,
	vr VehicleRepo,
	it ItemRepo,
	sh ShipmentRepo,
) *Service {
	once.Do(func() {
		service = &Service{
			cfg:          cfg,
			authRepo:     aur,
			profileRepo:  pr,
			vehicleRepo:  vr,
			itemRepo:     it,
			shipmentRepo: sh,
		}
	})

	return service
}

func Get() *Service {
	return service
}

type AuthRepo interface {
	CreateSession(ctx context.Context, session *models.UserSession) (*models.UserSession, error)
	DisableSessionByID(ctx context.Context, sessionID uuid.UUID) error
	GetSessionByTokenID(ctx context.Context, tokenID uuid.UUID) (*models.UserSession, error)
	UpdateSession(ctx context.Context, userSession *models.UserSession) error
}

type ProfileRepo interface {
	GetProfilesByEmail(ctx context.Context, email string) (models.User, error)
	AddNewProfile(ctx context.Context, newUser *models.User) (*models.User, error)
}

type Encryptor interface {
	EncryptPassword(password string) (string, error)
	CompareHashAndPassword(password, hash string) bool
}

type VehicleRepo interface {
	CreateVehicle(ctx context.Context, newVehicle *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error)
	GetAllVehicles(ctx context.Context) ([]models.Vehicle, error)
	GetUserVehicles(ctx context.Context, userID int) ([]models.Vehicle, error)
	DeleteVehicle(ctx context.Context, vehicleID int) error
}

type ItemRepo interface {
	CreateItem(ctx context.Context, newItem *models.Item) (*models.Item, error)
	GetAllItems(ctx context.Context) ([]models.Item, error)
	UpdateItem(ctx context.Context, newItem *models.Item) error
}

type ShipmentRepo interface {
	CreateShipment(ctx context.Context, newShipment *models.Shipment) (*models.Shipment, error)
	UpdateShipment(ctx context.Context, newShipment *models.Shipment) error
}
