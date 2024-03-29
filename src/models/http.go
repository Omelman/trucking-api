package models

// swagger:model
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// swagger:model
type LoginResponse struct {
	TokenPair TokenPair `json:"tokens_pair"`
	User      User      `json:"user"`
}

// swagger:model
type UserRegistrationRequest struct {
	RegistrationRequest
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

// swagger:model
type Statistics struct {
	AverageItemsQuantity   float32 `json:"average_items"`
	AverageItemsVolume     float32 `json:"average_items_volume"`
	AverageVehicleVolume   float32 `json:"average_vehicle_volume"`
	AverageVehicleCapacity float32 `json:"average_vehicle_capacity"`
}
