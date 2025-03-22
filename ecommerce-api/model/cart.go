package model

import "time"

type Cart struct {
	ID        string      `json:"id" gorm:"column:id;primaryKey"`
	UserID    string      `json:"user_id" gorm:"column:user_id"`
	Total     float64     `json:"sub_total" gorm:"column:sub_total"`
	Items     []*CartItem `json:"items" gorm:"foreignKey:CartID;references:ID"`
	CreatedAt time.Time   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time  `json:"-" gorm:"column:deleted_at"`
}

type CartItem struct {
	ID        string     `json:"id" gorm:"column:id;primaryKey"`
	CartID    string     `json:"cart_id" gorm:"column:cart_id"`
	ProductID string     `json:"product_id" gorm:"column:product_id"`
	Product   *Product   `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	Price     float64    `json:"price" gorm:"column:price"`
	Quantity  int        `json:"quantity" gorm:"column:quantity"`
	SubTotal  float64    `json:"sub_total" gorm:"column:sub_total"` // Price * Quantity
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
