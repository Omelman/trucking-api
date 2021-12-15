package models

type ShipmentStatus string

// Shipment statuses enum.
const (
	Waiting    ShipmentStatus = "waiting"
	Done       ShipmentStatus = "done"
	InProgress ShipmentStatus = "in progress"
)

type Shipment struct {
	ID        int            `json:"id"`
	Status    ShipmentStatus `json:"status"`
	VehicleID int            `json:"vehicle_id"`
}
