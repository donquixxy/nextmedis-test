package user

import (
	"ecommerce-api/commons"
	"ecommerce-api/modules"
	"ecommerce-api/modules/user/handler"
	"ecommerce-api/modules/user/repository"
	"ecommerce-api/modules/user/service"
	"ecommerce-api/server"
)

type UserModule struct {
}

func (s UserModule) WireRepository(appRepo *commons.Model, opt commons.Options) {
	appRepo.Repository.User = repository.NewUserRepository(opt)
}

func (s UserModule) WireService(serviceOption *commons.Model) {
	serviceOption.Service.User = service.NewUserService(*serviceOption)
}

func (s UserModule) RegisterHandlers(opt *server.ModelHandler) {
	userHandler := handler.UserHandler{}
	userHandler.ModelHandler = opt

	opt.Router.PublicApi.POST("/register", userHandler.Create)
	opt.Router.PublicApi.POST("/login", userHandler.Login)
}

func NewUserModule() modules.Module {
	return &UserModule{}
}
