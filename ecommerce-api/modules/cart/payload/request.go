package payload

import (
	"ecommerce-api/payload"
	"gorm.io/gorm"
)

type (
	CartCreate struct {
		ID     string
		UserID string
		Total  float64
		DbTx   *gorm.DB
	}

	CartItemCreate struct {
		ID        string
		CartID    string
		ProductID string
		Quantity  int
		SubTotal  float64
		Price     float64
		DbTx      *gorm.DB
	}

	CartUpdate struct {
		ID    string
		Total *float64
		DbTx  *gorm.DB
	}

	CartItemUpdate struct {
		ID        string
		ProductID *string
		Quantity  *int
		SubTotal  *float64
		Price     *float64
		DbTx      *gorm.DB
	}

	CartFilter struct {
		ID     *string
		UserID *string
		payload.Params
		payload.Pagination
		WithItems bool
	}
	CartDelete struct {
		ID string
	}

	CartItemFilter struct {
		ID        *string
		CartID    *string
		ProductID *string
		NotInID   []string
		payload.Params
		payload.Pagination
	}

	CartUpsert struct {
		UserID    string
		ProductID string `json:"product_id" form:"product_id" validate:"required"`
		Quantity  *int   `json:"quantity" form:"quantity" validate:"required"`
	}

	CartItemDelete struct {
		ID   string
		DbTx *gorm.DB
	}
)
