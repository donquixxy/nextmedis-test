package service

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/user/interfaces"
	"ecommerce-api/modules/user/payload"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userService struct {
	opt commons.Model
}

func (s userService) Get(ctx context.Context, request payload.UserGet) (*model.User, error) {
	return s.opt.Repository.User.Get(ctx, request)
}

func (s userService) generateToken(u *model.User, tType int) (string, error) {
	secret := s.opt.Config.JwtCfg.Secret

	expAt := jwt.NewNumericDate(time.Now().Add(72 * time.Hour))

	// Refresh token
	if tType == 1 {
		secret = s.opt.Config.JwtCfg.SecretRefresh
		expAt = jwt.NewNumericDate(time.Now().Add(240 * time.Hour))
	}

	claims := payload.TokenPayload{
		ID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ari", // Ideally from env
			Subject:   "test",
			ExpiresAt: expAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        u.ID,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, err := t.SignedString([]byte(secret))

	if err != nil {
		log.Errorf("[generateToken] - %v", err)
		return "", err
	}

	return str, nil
}

func (s userService) Login(ctx context.Context, request payload.Login) (*payload.LoginResponse, string, error) {
	// Retrieve user by email
	userData, err := s.opt.Repository.User.Get(ctx, payload.UserGet{
		Email: &request.Email,
	})

	if err != nil {
		log.Errorf("[Login] - %v | %+v", err, fmt.Sprintf("+%v", request))
		return nil, "user data not found", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password)); err != nil {
		log.Errorf("[Login] - %v | %+v", err, fmt.Sprintf("+%v", request))
		return nil, "invalid password", err
	}

	// Generate  token
	token, err := s.generateToken(userData, 0)

	if err != nil {
		return nil, "token generate error", err
	}

	// Refresh token
	rToken, err := s.generateToken(userData, 1)

	if err != nil {
		log.Errorf("[Login] - %v | %+v", err, fmt.Sprintf("+%v", request))
		return nil, "token generate error", err
	}

	// update user token
	_, err = s.opt.Repository.User.Update(ctx, payload.UserUpdate{
		Token: &token,
		ID:    userData.ID,
	})

	if err != nil {
		log.Errorf("[Login] - %v | %+v", err, fmt.Sprintf("+%v", request))
		return nil, "failed to generate token", err
	}

	return &payload.LoginResponse{
		Token:        token,
		RefreshToken: rToken,
	}, "success", nil
}

func (s userService) Create(ctx context.Context, request payload.UserCreate) (*model.User, error) {
	if request.ID == nil {
		id := uuid.NewString()

		request.ID = &id
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)

	if err != nil {
		log.Errorf("[Create] - %v", err)
		return nil, errors.New("an error occurred in our system. please try again later")
	}

	request.Password = string(hashedPassword)

	return s.opt.Repository.User.Create(ctx, request)
}

func NewUserService(opt commons.Model) interfaces.UserService {
	return &userService{opt: opt}
}
