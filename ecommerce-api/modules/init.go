package modules

import (
	"ecommerce-api/commons"
	"ecommerce-api/server"
)

type Module interface {
	WireRepository(appRepo *commons.Model, opt commons.Options)
	WireService(serviceOption *commons.Model)
	RegisterHandlers(opt *server.ModelHandler)
}
