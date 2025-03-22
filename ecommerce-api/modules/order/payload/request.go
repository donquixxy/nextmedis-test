package payload

import "gorm.io/gorm"

type (
	OrderCreate struct {
		ID     string
		UserID string
		Total  float64
		DbTx   *gorm.DB
	}

	OrderItemCreate struct {
		ID        string
		OrderID   string
		ProductID string
		Quantity  int
		Price     float64
		SubTotal  float64
		DbTx      *gorm.DB
	}

	SubmitOrder struct {
		CartID string `json:"cart_id" form:"cart_id" validate:"required"`
		UserID string
	}
)
