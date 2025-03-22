package handler

import (
	"ecommerce-api/modules/user/payload"
	payload2 "ecommerce-api/payload"
	"ecommerce-api/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	*server.ModelHandler
}

func (h *UserHandler) Create(c echo.Context) error {
	pyld := payload.UserCreate{}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, err := h.Service.User.Create(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *UserHandler) Login(c echo.Context) error {
	pyld := payload.Login{}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, msg, err := h.Service.User.Login(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(msg))
	}

	return c.JSON(http.StatusOK, payload2.SuccessResponse(result, msg))
}
