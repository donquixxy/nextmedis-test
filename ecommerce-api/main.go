package main

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/config"
	"ecommerce-api/model"
	"ecommerce-api/modules/cart"
	"ecommerce-api/modules/order"
	"ecommerce-api/modules/product"
	"ecommerce-api/modules/user"
	"ecommerce-api/server"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	op := bootstrapApp()
	router := wire(op)

	go func() {
		if err := router.Echo.Start(fmt.Sprintf(":%v", op.Config.AppCfg.AppPort)); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	router.Echo.Shutdown(context.Background())
}

func bootstrapApp() commons.Options {

	cfg := config.NewConfig()
	db, err := config.GetDatabaseInstance(cfg)

	if err != nil {
		panic(err)
	}

	runAutoMigrate(db)
	log.Infof("App is running on environment %v", cfg.AppCfg.AppEnv)
	return commons.Options{
		Config:   cfg,
		Database: db,
	}
}

func wire(opt commons.Options) *server.Router {
	srvOpt := commons.Model{
		Options:    opt,
		Service:    &commons.Service{},
		Repository: &commons.Repository{},
	}

	handlerOpt := server.ModelHandler{
		Model:  srvOpt,
		Router: server.NewRouter(srvOpt),
	}

	userModule := user.NewUserModule()
	userModule.WireRepository(&srvOpt, opt)
	userModule.WireService(&srvOpt)
	userModule.RegisterHandlers(&handlerOpt)

	productModule := product.NewProductModule()
	productModule.WireRepository(&srvOpt, opt)
	productModule.WireService(&srvOpt)
	productModule.RegisterHandlers(&handlerOpt)

	cartModule := cart.NewCartModule()
	cartModule.WireRepository(&srvOpt, opt)
	cartModule.WireService(&srvOpt)
	cartModule.RegisterHandlers(&handlerOpt)

	orderModule := order.NewOrderModule()
	orderModule.WireRepository(&srvOpt, opt)
	orderModule.WireService(&srvOpt)
	orderModule.RegisterHandlers(&handlerOpt)

	return handlerOpt.Router
}

func runAutoMigrate(db *gorm.DB) {
	// For the sake of test. run auto  migrate by default
	// Ideally should be written in sql query
	db.AutoMigrate(&model.User{},
		&model.Product{},
		&model.Cart{},
		&model.CartItem{},
		&model.Order{},
		&model.OrderItem{})
}
