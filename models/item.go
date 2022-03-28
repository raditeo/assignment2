package models

type Item struct {
	ItemID      uint `gorm:"primaryKey"`
	ItemCode    string
	Description string
	Quantity    uint
	OrderID     uint
}
