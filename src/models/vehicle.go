package models

import (
	"time"
)

type VehicleType string

// Vehicle types enum.
const (
	Tent    VehicleType = "tent"
	Covered VehicleType = "covered"
)

type Vehicle struct {
	tableName struct{} `pg:"vehicle"`

	ID               int         `json:"id"`
	Type             VehicleType `json:"type"`
	ConnectionString string      `json:"connection_string"`
	CarryingCapacity int         `json:"carrying_capacity"`
	UsefulVolume     int         `json:"useful_volume"`
	Length           int         `json:"length"`
	Height           int         `json:"height"`
	Width            int         `json:"width"`
	OwnerID          int         `json:"owner_id"`

	CreatedAt time.Time `json:"-"`

	Owner *User `json:"-" pg:"fk:owner_id"`
}
