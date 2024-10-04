package jaquar

import (
	"time"
)

// JaquarProduct represents the jaquar_products table
type Product struct {
	ID          uint64    `json:"id"`
	Series      string    `json:"series"`
	ColorCode   string    `json:"color_code"`
	CodeNumber  string    `json:"code_number"`
	Description string    `json:"description"`
	NRP         uint64    `json:"nrp"`
	MRP         uint64    `json:"mrp" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" `
}

func (p Product) TableName() string {
	return "jaquar_products"
}
