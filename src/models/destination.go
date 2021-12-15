package models

type Destination struct {
	ID     int     `json:"id"`
	Lat    float32 `json:"lat"`
	Lon    float32 `json:"lon"`
	ItemID int     `json:"item_id"`

	Item *Item `json:"-"`
}
