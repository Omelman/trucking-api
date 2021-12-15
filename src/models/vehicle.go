package models

type VehicleType string

// Vehicle types enum.
const (
	Tent    VehicleType = "tent"
	Covered VehicleType = "covered"
)

type Vehicle struct {
	ID     int         `json:"id"`
	Type   VehicleType `json:"type"`
	UserID int         `json:"user_id"`

	User *User `json:"-"`
}
