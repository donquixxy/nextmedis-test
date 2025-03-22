package interfaces

import (
	"context"
	"ecommerce-api/model"
	"ecommerce-api/modules/user/payload"
)

type UserService interface {
	Create(ctx context.Context, request payload.UserCreate) (*model.User, error)
	Login(ctx context.Context, request payload.Login) (*payload.LoginResponse, string, error)
	Get(ctx context.Context, request payload.UserGet) (*model.User, error)
}
