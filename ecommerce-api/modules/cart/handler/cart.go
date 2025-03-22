package handler

import (
	"ecommerce-api/modules/cart/payload"
	payload2 "ecommerce-api/payload"
	"ecommerce-api/server"
	"ecommerce-api/server/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CartHandler struct {
	*server.ModelHandler
}

func (h *CartHandler) AddItemCart(c echo.Context) error {
	currentUser := middleware.CurrentUser(c)

	pyld := payload.CartUpsert{
		UserID: currentUser.ID,
	}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, msg, err := h.Service.Cart.AddItemCart(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, payload2.SuccessResponse(result, msg))
}

func (h *CartHandler) Get(c echo.Context) error {
	currentUser := middleware.CurrentUser(c)

	pyld := payload.CartFilter{
		UserID:    &currentUser.ID,
		WithItems: true,
	}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, msg, err := h.Service.Cart.Get(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(msg))
	}

	return c.JSON(http.StatusOK, payload2.SuccessResponse(result, msg))
}
