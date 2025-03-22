package payload

import "ecommerce-api/payload"

type (
	ProductCreate struct {
		ID    string
		Name  string  `json:"name" form:"name" validate:"required"`
		Price float64 `json:"price" form:"price" validate:"required"`
	}

	ProductFilter struct {
		ID   *string `query:"id"`
		Name *string `query:"name"`
		payload.Pagination
		payload.Params
	}
)
