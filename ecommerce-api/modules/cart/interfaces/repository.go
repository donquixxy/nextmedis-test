package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/cart/payload"
)

type CartRepository interface {
	CartGet(ctx context.Context, filter payload.CartFilter) (*model.Cart, error)
	CartCreate(ctx context.Context, data payload.CartCreate) (*model.Cart, error)
	CartUpdate(ctx context.Context, data payload.CartUpdate) (*model.Cart, error)
	CartDelete(ctx context.Context, data payload.CartDelete) error
	CartItemCreate(ctx context.Context, data payload.CartItemCreate) (*model.CartItem, error)
	CartItemUpdate(ctx context.Context, data payload.CartItemUpdate) (*model.CartItem, error)
	CartItemGet(ctx context.Context, filter payload.CartItemFilter) (*model.CartItem, error)
	CartItemDelete(ctx context.Context, data payload.CartItemDelete) error
	CartItemGetAll(ctx context.Context, filter payload.CartItemFilter) ([]*model.CartItem, int64, error)
}
