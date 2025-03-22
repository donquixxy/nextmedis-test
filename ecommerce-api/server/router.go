package server

import (
	"ecommerce-api/commons"
	"ecommerce-api/config"
	middleware2 "ecommerce-api/server/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ModelHandler struct {
	commons.Model
	Router *Router
}

type Router struct {
	PublicApi  *echo.Group
	PrivateApi *echo.Group
	Echo       *echo.Echo
}

func NewRouter(opt commons.Model) *Router {
	e := echo.New()

	e.Validator = config.NewValidator()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware2.LoggerMiddlware())

	publicApi := e.Group("")

	privateApi := e.Group("api")
	privateApi.Use(middleware2.ValidateJWT(opt))

	return &Router{
		PublicApi:  publicApi,
		PrivateApi: privateApi,
		Echo:       e,
	}
}
