package models

import "time"

type Order struct {
	OrderID      uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}
