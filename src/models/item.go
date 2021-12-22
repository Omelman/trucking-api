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
	tableName struct{} `pg:"item"`

	ID          int          `json:"id"`
	Description string       `json:"description"`
	Quantity    int          `json:"quantity"`
	Volume      int          `json:"volume"`
	Weight      int          `json:"weight"`
	Category    ItemCategory `json:"category"`
	UserID      int          `json:"user_id"`
	ShipmentID  int          `json:"shipment_id"`

	Date *time.Time `json:"date"`

	User     *User     `json:"-" fk:"user_id"`
	Shipment *Shipment `json:"-" fk:"shipment_id"`
}
