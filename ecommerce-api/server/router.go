package server

import (
	"ecommerce-api/commons"
	"ecommerce-api/config"
	middleware2 "ecommerce-api/server/middleware"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
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
	now := time.Now()
	e.Validator = config.NewValidator()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware2.LoggerMiddlware())

	publicApi := e.Group("")

	publicApi.GET("/health-check", func(c echo.Context) error {
		uptime := time.Since(now)

		uptimeString := formatUptime(uptime)
		return c.JSON(200, map[string]interface{}{
			"status": "ok",
			"uptime": uptimeString,
		})
	})

	privateApi := e.Group("api")
	privateApi.Use(middleware2.ValidateJWT(opt))

	return &Router{
		PublicApi:  publicApi,
		PrivateApi: privateApi,
		Echo:       e,
	}
}

func formatUptime(uptime time.Duration) string {
	days := int(uptime.Hours() / 24)
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	if days > 0 {
		return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
	} else if hours > 0 {
		return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d minutes, %d seconds", minutes, seconds)
	} else {
		return fmt.Sprintf("%d seconds", seconds)
	}
}
