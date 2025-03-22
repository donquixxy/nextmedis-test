package service

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/product/interfaces"
	"ecommerce-api/modules/product/payload"
	"github.com/google/uuid"
)

type productService struct {
	opt commons.Model
}

func (s productService) GetAll(ctx context.Context, filter payload.ProductFilter) ([]*model.Product, int64, error) {
	return s.opt.Repository.Product.GetAll(ctx, filter)
}

func (s productService) Create(ctx context.Context, data payload.ProductCreate) (*model.Product, error) {
	if data.ID == "" {
		data.ID = uuid.NewString()
	}

	return s.opt.Repository.Product.Create(ctx, data)
}

func NewProductService(opt commons.Model) interfaces.ProductService {
	return &productService{
		opt: opt,
	}
}
