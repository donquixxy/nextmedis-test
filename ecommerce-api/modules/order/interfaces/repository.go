package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/order/payload"
)

type OrderRepository interface {
	Create(ctx context.Context, data payload.OrderCreate) (*model.Order, error)
	CreateItem(ctx context.Context, data payload.OrderItemCreate) (*model.OrderItem, error)
}
