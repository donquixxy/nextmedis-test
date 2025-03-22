package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/product/payload"
)

type ProductService interface {
	Create(ctx context.Context, data payload.ProductCreate) (*model.Product, error)
	GetAll(ctx context.Context, filter payload.ProductFilter) ([]*model.Product, int64, error)
}
