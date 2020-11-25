package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string
	CompanyID string
}

type Price struct {
	Price     int64
	ProductID uint
	Product   Product
}
