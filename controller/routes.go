package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prabowohadip/tax/config"
	"net/http"
)

func Routing() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	e.Any("/test-ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
		})
	})

	v1 := e.Group("/v1")
	order := v1.Group("/order")

	order.POST("", CreateOrder)
	order.GET("", ListOrder)

	e.Logger.Fatal(e.Start(config.Config.App.Port))
}

