package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/order/payload"
)

type OrderService interface {
	SubmitOrder(ctx context.Context, data payload.SubmitOrder) (*model.Order, string, error)
}
