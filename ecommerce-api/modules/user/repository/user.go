package repository

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/user/interfaces"
	"ecommerce-api/modules/user/payload"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
)

type userRepository struct {
	opt commons.Options
}

func (s userRepository) Update(ctx context.Context, data payload.UserUpdate) (*model.User, error) {
	var (
		result *model.User
		query  = s.opt.Database.WithContext(ctx)
	)

	if err := query.Where("id = ?", data.ID).First(&result).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if data.Password != nil {
		result.Password = *data.Password
	}

	if data.Name != nil {
		result.Name = *data.Name
	}

	if data.Email != nil {
		result.Email = *data.Email
	}

	if data.Token != nil {
		result.Token = data.Token
	}

	if err := query.Updates(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s userRepository) Get(ctx context.Context, filter payload.UserGet) (*model.User, error) {
	var result *model.User

	query := s.opt.Database.WithContext(ctx)

	if filter.ID != nil {
		query = query.Where("id = ?", *filter.ID)
	}

	if filter.Password != nil {
		query = query.Where("password = ?", *filter.Password)
	}

	if filter.Name != nil {
		query = query.Where("name = ?", *filter.Name)
	}

	if filter.Token != nil {
		query = query.Where("token = ?", *filter.Token)
	}

	if err := query.First(&result).Error; err != nil {
		log.Printf("[GetUser] - %v", err)
		return nil, err
	}

	return result, nil
}

func (s userRepository) Create(ctx context.Context, data payload.UserCreate) (*model.User, error) {
	var (
		now   = time.Now()
		query = s.opt.Database.WithContext(ctx)
	)

	result := &model.User{
		ID: func() string {
			if data.ID == nil {
				id := uuid.NewString()
				return id
			}
			return *data.ID
		}(),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := query.Create(result).Error; err != nil {
		log.Printf("[Create] - %v", err)
		return nil, err
	}

	return result, nil
}

func NewUserRepository(opt commons.Options) interfaces.UserRepository {
	return &userRepository{
		opt: opt,
	}
}
