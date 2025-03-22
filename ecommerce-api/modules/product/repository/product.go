package repository

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/product/interfaces"
	"ecommerce-api/modules/product/payload"
	"fmt"
	"time"
)

type productRepository struct {
	opt commons.Options
}

func (s productRepository) Get(ctx context.Context, filter payload.ProductFilter) (*model.Product, error) {
	var (
		result *model.Product
		query  = s.opt.Database.WithContext(ctx)
	)

	if filter.ID != nil {
		query = query.Where("id = ?", *filter.ID)
	}

	if filter.Name != nil {
		query = query.Where("name LIKE ?", "%"+*filter.Name+"%")
	}

	if err := query.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s productRepository) Create(ctx context.Context, data payload.ProductCreate) (*model.Product, error) {
	var (
		query = s.opt.Database.WithContext(ctx)
		now   = time.Now()
	)

	result := &model.Product{
		ID:        data.ID,
		Name:      data.Name,
		Price:     data.Price,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := query.Create(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s productRepository) GetAll(ctx context.Context, filter payload.ProductFilter) ([]*model.Product, int64, error) {
	var (
		query  = s.opt.Database.WithContext(ctx)
		count  int64
		result []*model.Product
	)

	query = query.Model(&model.Product{}).Where("deleted_at is null")

	if filter.ID != nil {
		query = query.Where("id = ?", *filter.ID)
	}

	if filter.Name != nil {
		query = query.Where("name like ?", "%"+*filter.Name+"%")
	}

	err := query.Count(&count).Error

	if err != nil {
		return nil, 0, err
	}

	if !filter.Pagination.All {
		if filter.Pagination.Page >= 1 {
			offset := (filter.Pagination.Page - 1) * filter.Pagination.Limit
			query = query.Offset(offset)
		}

		if filter.Pagination.Limit > 0 {
			query = query.Limit(filter.Pagination.Limit)
		}
	}

	if filter.OrderBy != "" && filter.SortedBy != "" {
		// Fieldname + asc/desc
		query = query.Order(fmt.Sprintf("products.%v %v", filter.SortedBy, filter.OrderBy))
	} else {
		query = query.Order("products.created_at desc")
	}

	err = query.Find(&result).Error

	if err != nil {
		return nil, count, err
	}

	return result, count, nil
}

func NewProductRepository(opt commons.Options) interfaces.ProductRepository {
	return &productRepository{
		opt: opt,
	}
}
