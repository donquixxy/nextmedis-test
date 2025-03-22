package model

import "time"

type User struct {
	ID        string     `json:"id" gorm:"column:id;primaryKey"`
	Name      string     `json:"name" gorm:"column:name"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"-" gorm:"column:password"`
	Token     *string    `json:"-" gorm:"column:token"`
	CreatedAt time.Time  `json:"created-at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated-at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
