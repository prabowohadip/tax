package model

import "time"

type OrderItemSchema struct {
	BaseModel
	Name      string     `gorm:"not null"    json:"name"`
	TaxCode   string     `gorm:"not null"    json:"tax_code"`
	Price     float64    `gorm:"not null"    json:"price"`
}

type BaseModel struct {
	ID        uint 		 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time	 `json:"created_at,omitempty"`
}

func (OrderItemSchema) TableName() string {
	return "order_item"
}

type OrderRequest struct {
	Order     []OrderItemRequest  `json:"order" valid:"required"`
}

type OrderItemRequest struct {
	Name      string  `json:"name" valid:"required"`
	TaxCode   string  `json:"tax_code" valid:"required"`
	Price     float64 `json:"price" valid:"required"`
}

type OrderItemTax struct {
	Name       string  `json:"name"`
	TaxCode    string  `json:"tax_code"`
	Price      float64 `json:"price"`
	Type       string  `json:"type"`
	Refundable bool    `json:"refundable"`
	Tax        float64 `json:"tax"`
	Amount     float64 `json:"amount"`
}