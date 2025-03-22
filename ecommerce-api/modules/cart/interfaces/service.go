package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/cart/payload"
)

type CartService interface {
	AddItemCart(ctx context.Context, data payload.CartUpsert) (*model.Cart, string, error)
	Get(ctx context.Context, filter payload.CartFilter) (*model.Cart, string, error)
}
