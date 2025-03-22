package cart

import (
	"ecommerce-api/commons"
	"ecommerce-api/modules"
	"ecommerce-api/modules/cart/handler"
	"ecommerce-api/modules/cart/repository"
	"ecommerce-api/modules/cart/service"
	"ecommerce-api/server"
)

type cartModule struct {
}

func (s cartModule) WireRepository(appRepo *commons.Model, opt commons.Options) {
	appRepo.Repository.Cart = repository.NewCartRepository(opt)
}

func (s cartModule) WireService(serviceOption *commons.Model) {
	serviceOption.Service.Cart = service.NewCartService(*serviceOption)
}

func (s cartModule) RegisterHandlers(opt *server.ModelHandler) {
	cartHandler := handler.CartHandler{}
	cartHandler.ModelHandler = opt

	opt.Router.PrivateApi.POST("/cart", cartHandler.AddItemCart)
	opt.Router.PrivateApi.GET("/cart", cartHandler.Get)
}

func NewCartModule() modules.Module {
	return &cartModule{}
}
