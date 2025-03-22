package handler

import (
	"ecommerce-api/modules/product/payload"
	payload2 "ecommerce-api/payload"
	"ecommerce-api/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	*server.ModelHandler
}

func (h *ProductHandler) Create(c echo.Context) error {
	pyld := payload.ProductCreate{}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, err := h.Service.Product.Create(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload2.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, payload2.SuccessResponse(result, "success"))
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	pyld := payload.ProductFilter{}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, payload2.FailedResponse(err.Error()))
	}

	result, count, err := h.Service.Product.GetAll(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload2.FailedResponse(err.Error()))
	}

	var total = 10
	var page = 1

	if pyld.Limit != 0 {
		total = pyld.Limit
	}

	if pyld.Page != 0 {
		page = pyld.Page
	}

	pgnt := payload2.Paginate(len(result), page, int(count), total)

	return c.JSON(http.StatusOK, payload2.SuccessResponsePagination(result, pgnt))
}
