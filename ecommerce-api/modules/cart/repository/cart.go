package repository

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/cart/interfaces"
	"ecommerce-api/modules/cart/payload"
	"errors"
	"time"
)

type cartRepository struct {
	opt commons.Options
}

func (s *cartRepository) CartItemGetAll(ctx context.Context, filter payload.CartItemFilter) ([]*model.CartItem, int64, error) {
	var (
		result []*model.CartItem
		query  = s.opt.Database.WithContext(ctx)
		count  int64
	)

	query = query.Model(&model.CartItem{}).Where("deleted_at IS NULL")

	if filter.CartID != nil {
		query = query.Where("cart_id = ?", *filter.CartID)
	}

	if filter.ID != nil {
		query = query.Where("cart_id = ?", *filter.ID)
	}

	if filter.ProductID != nil {
		query = query.Where("product_id = ?", *filter.ProductID)
	}

	if filter.NotInID != nil && len(filter.NotInID) > 0 {
		query = query.Where("id NOT IN (?)", filter.NotInID)
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

	err = query.Find(&result).Error

	if err != nil {
		return nil, count, err
	}

	return result, count, nil
}

func (s *cartRepository) CartItemDelete(ctx context.Context, data payload.CartItemDelete) error {
	var (
		query  = s.opt.Database.WithContext(ctx)
		result *model.CartItem
	)

	if data.DbTx != nil {
		query = data.DbTx.WithContext(ctx)
	}

	if err := query.Where("id = ?", data.ID).First(&result).Error; err != nil {
		return errors.New("cart item not found")
	}

	now := time.Now()
	result.DeletedAt = &now

	if err := query.Updates(&result).Error; err != nil {
		return err
	}

	return nil
}

func NewCartRepository(opt commons.Options) interfaces.CartRepository {
	return &cartRepository{opt: opt}
}

func (s *cartRepository) CartGet(ctx context.Context, filter payload.CartFilter) (*model.Cart, error) {
	var (
		result *model.Cart
		query  = s.opt.Database.WithContext(ctx)
	)

	query = query.Where("deleted_at IS NULL")

	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID)
	}

	if filter.UserID != nil {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.WithItems {
		query = query.Preload("Items", "deleted_at is null")
		query = query.Preload("Items.Product")
	}

	if err := query.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cartRepository) CartCreate(ctx context.Context, data payload.CartCreate) (*model.Cart, error) {
	var (
		query = s.opt.Database.WithContext(ctx)
		now   = time.Now()
	)

	if data.DbTx != nil {
		query = data.DbTx.WithContext(ctx)
	}

	result := &model.Cart{
		ID:        data.ID,
		UserID:    data.UserID,
		Total:     data.Total,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := query.Create(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cartRepository) CartUpdate(ctx context.Context, data payload.CartUpdate) (*model.Cart, error) {
	var (
		query  = s.opt.Database.WithContext(ctx)
		now    = time.Now()
		result *model.Cart
	)

	if data.DbTx != nil {
		query = data.DbTx.WithContext(ctx)
	}

	if err := query.Where("id = ?", data.ID).First(&result).Error; err != nil {
		return nil, errors.New("cart not found")
	}

	if data.Total != nil {
		result.Total = *data.Total
	}

	result.UpdatedAt = now

	if err := query.Updates(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cartRepository) CartDelete(ctx context.Context, data payload.CartDelete) error {
	var (
		query  = s.opt.Database.WithContext(ctx)
		now    = time.Now()
		result *model.Cart
	)

	if err := query.Where("id = ?", data.ID).First(&result).Error; err != nil {
		return errors.New("cart not found")
	}

	result.DeletedAt = &now

	if err := query.Updates(&result).Error; err != nil {
		return err
	}

	return nil
}

func (s *cartRepository) CartItemCreate(ctx context.Context, data payload.CartItemCreate) (*model.CartItem, error) {
	var (
		now   = time.Now()
		query = s.opt.Database.WithContext(ctx)
	)

	if data.DbTx != nil {
		query = data.DbTx.WithContext(ctx)
	}

	result := &model.CartItem{
		ID:        data.ID,
		CartID:    data.CartID,
		ProductID: data.ProductID,
		Price:     data.Price,
		Quantity:  data.Quantity,
		SubTotal:  data.SubTotal,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := query.Create(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cartRepository) CartItemUpdate(ctx context.Context, data payload.CartItemUpdate) (*model.CartItem, error) {
	var (
		now    = time.Now()
		result *model.CartItem
		query  = s.opt.Database.WithContext(ctx)
	)

	if data.DbTx != nil {
		query = data.DbTx.WithContext(ctx)
	}

	if err := query.Where("id = ?", data.ID).First(&result).Error; err != nil {
		return nil, errors.New("cart item not found")
	}

	if data.ProductID != nil {
		result.ProductID = *data.ProductID
	}

	if data.SubTotal != nil {
		result.SubTotal = *data.SubTotal
	}

	if data.Quantity != nil {
		result.Quantity = *data.Quantity
	}

	if data.Price != nil {
		result.Price = *data.Price
	}

	result.UpdatedAt = now

	if err := query.Updates(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cartRepository) CartItemGet(ctx context.Context, filter payload.CartItemFilter) (*model.CartItem, error) {
	var (
		result *model.CartItem
		query  = s.opt.Database.WithContext(ctx)
	)

	query = query.Model(&model.CartItem{}).Where("deleted_at is null")

	if filter.ID != nil {
		query = query.Where("id = ?", *filter.ID)
	}

	if filter.CartID != nil {
		query = query.Where("cart_id = ?", *filter.CartID)
	}

	if filter.ProductID != nil {
		query = query.Where("product_id = ?", *filter.ProductID)
	}

	if err := query.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
