package models

//Query - model
type Query struct {
	Name string  `json:"name"`
	Lat  float64 `gorm:"type:decimal(10,8)"`
	Lng  float64 `gorm:"type:decimal(11,8)"`
}
