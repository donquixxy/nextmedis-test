package order

import (
	"ecommerce-api/commons"
	"ecommerce-api/modules"
	"ecommerce-api/modules/order/handler"
	"ecommerce-api/modules/order/repository"
	"ecommerce-api/modules/order/service"
	"ecommerce-api/server"
)

type orderModule struct {
}

func (s orderModule) WireRepository(appRepo *commons.Model, opt commons.Options) {
	appRepo.Repository.Order = repository.NewOrderRepository(opt)
}

func (s orderModule) WireService(serviceOption *commons.Model) {
	serviceOption.Service.Order = service.NewOrderService(*serviceOption)
}

func (s orderModule) RegisterHandlers(opt *server.ModelHandler) {
	orderHandler := handler.OrderHandler{}
	orderHandler.ModelHandler = opt

	opt.Router.PrivateApi.POST("/order", orderHandler.SubmitOrder)
}

func NewOrderModule() modules.Module {
	return &orderModule{}
}
