package models

import (
	"time"
)

type ItemCategory string

// Item categories enum.
const (
	Raw      ItemCategory = "raw materials"
	Building ItemCategory = "building materials"
	Products ItemCategory = "products"
)

type Item struct {
	ID          int          `json:"id"`
	Description string       `json:"description"`
	Quantity    string       `json:"quantity"`
	Volume      int          `json:"volume"`
	Weight      int          `json:"weight"`
	Category    ItemCategory `json:"category"`
	UserID      int          `json:"user_id"`
	ShipmentID  int          `json:"shipment_id"`

	Date *time.Time `json:"date"`

	User     *User     `json:"-"`
	Shipment *Shipment `json:"-"`
}
