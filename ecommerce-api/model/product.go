package model

import "time"

type Product struct {
	ID        string     `gorm:"column:id;primaryKey" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Price     float64    `gorm:"column:price" json:"price"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
