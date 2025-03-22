package model

import "time"

type Order struct {
	ID        string       `json:"id" gorm:"column:id;primaryKey"`
	UserID    string       `json:"user_id" gorm:"column:user_id"`
	Total     float64      `json:"total" gorm:"column:total"`
	Items     []*OrderItem `json:"items" gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time   `json:"deleted_at" gorm:"column:deleted_at"`
}

type OrderItem struct {
	ID        string    `json:"id" gorm:"column:id"`
	OrderID   string    `json:"order_id" gorm:"column:order_id"`
	ProductID string    `json:"product_id" gorm:"column:product_id"`
	Price     float64   `json:"price" gorm:"column:price"`
	Quantity  int       `json:"quantity" gorm:"column:quantity"`
	SubTotal  float64   `json:"sub_total" gorm:"column:sub_total"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
