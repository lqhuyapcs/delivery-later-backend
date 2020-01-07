package models

//Query - model
type GetDistance struct {
	Lat1 float64 `gorm:"type:decimal(10,8)"`
	Lng1 float64 `gorm:"type:decimal(11,8)"`
	Lat2 float64 `gorm:"type:decimal(10,8)"`
	Lng2 float64 `gorm:"type:decimal(11,8)"`
}
