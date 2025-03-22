package repository

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/order/interfaces"
	"ecommerce-api/modules/order/payload"
	"time"
)

type orderRepository struct {
	opt commons.Options
}

func (s orderRepository) Create(ctx context.Context, data payload.OrderCreate) (*model.Order, error) {
	var (
		now   = time.Now()
		query = s.opt.Database.WithContext(ctx)
	)

	if data.DbTx != nil {
		query = data.DbTx.WithContext(ctx)
	}

	result := &model.Order{
		ID:        data.ID,
		UserID:    data.UserID,
		Total:     data.Total,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := query.Create(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s orderRepository) CreateItem(ctx context.Context, data payload.OrderItemCreate) (*model.OrderItem, error) {
	var (
		now   = time.Now()
		query = s.opt.Database.WithContext(ctx)
	)

	result := &model.OrderItem{
		ID:        data.ID,
		OrderID:   data.OrderID,
		ProductID: data.ProductID,
		Price:     data.Price,
		Quantity:  data.Quantity,
		SubTotal:  data.SubTotal,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := query.Create(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func NewOrderRepository(opt commons.Options) interfaces.OrderRepository {
	return &orderRepository{
		opt: opt,
	}
}
