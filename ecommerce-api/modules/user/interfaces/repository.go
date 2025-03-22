package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/user/payload"
)

type UserRepository interface {
	Get(ctx context.Context, filter payload.UserGet) (*model.User, error)
	Create(ctx context.Context, data payload.UserCreate) (*model.User, error)
	Update(ctx context.Context, data payload.UserUpdate) (*model.User, error)
}
