package models

//Query - model
type AddressOrder struct {
	ID      uint
	Address string  `json:"address"`
	Lat     float64 `gorm:"type:decimal(10,8)"`
	Lng     float64 `gorm:"type:decimal(11,8)"`
}
