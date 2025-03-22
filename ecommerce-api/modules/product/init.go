package product

import (
	"ecommerce-api/commons"
	"ecommerce-api/modules"
	"ecommerce-api/modules/product/handler"
	"ecommerce-api/modules/product/repository"
	"ecommerce-api/modules/product/service"
	"ecommerce-api/server"
)

type productModule struct {
}

func (s productModule) WireRepository(appRepo *commons.Model, opt commons.Options) {
	appRepo.Repository.Product = repository.NewProductRepository(opt)
}

func (s productModule) WireService(serviceOption *commons.Model) {
	serviceOption.Service.Product = service.NewProductService(*serviceOption)
}

func (s productModule) RegisterHandlers(opt *server.ModelHandler) {
	productHandler := handler.ProductHandler{}
	productHandler.ModelHandler = opt

	opt.Router.PublicApi.POST("/product", productHandler.Create)
	opt.Router.PublicApi.GET("/product", productHandler.GetAll)
}

func NewProductModule() modules.Module {
	return &productModule{}
}
