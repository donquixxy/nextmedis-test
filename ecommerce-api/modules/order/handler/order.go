package handler

import (
	"ecommerce-api/modules/order/payload"
	payload2 "ecommerce-api/payload"
	"ecommerce-api/server"
	"ecommerce-api/server/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderHandler struct {
	*server.ModelHandler
}

func (h *OrderHandler) SubmitOrder(c echo.Context) error {
	user := middleware.CurrentUser(c)

	pyld := payload.SubmitOrder{
		UserID: user.ID,
	}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, msg, err := h.Service.Order.SubmitOrder(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(msg))
	}

	return c.JSON(http.StatusOK, payload2.SuccessResponse(result, msg))
}
