package payload

import (
	"math"
)

type ResponseSuccessPagination struct {
	Message    string                `json:"message"`
	Data       any                   `json:"data"`
	Pagination ApiResponsePagination `json:"meta"`
}

type ResponseSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseFailed struct {
	Message string `json:"message"`
}

type ApiResponsePagination struct {
	Count       int `json:"count"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
}

func FailedResponse(msg string) ResponseFailed {
	return ResponseFailed{
		Message: msg,
	}
}

func SuccessResponse(data any, msg string) ResponseSuccess {
	return ResponseSuccess{
		Message: msg,
		Data:    data,
	}
}

func SuccessResponsePagination(data any, pgnt ApiResponsePagination) ResponseSuccessPagination {
	return ResponseSuccessPagination{
		Message:    "success",
		Data:       data,
		Pagination: pgnt,
	}
}

func Paginate(items int, page int, total int, limit int) ApiResponsePagination {
	totalPages := TotalPages(total, limit)

	return ApiResponsePagination{
		Count:       items,
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  totalPages,
	}
}

func TotalPages(items int, limit int) int {
	var totalPages float64

	if items < limit {
		totalPages = 1
	} else {
		if limit == 0 {
			limit = 10
		}

		totalPages = math.Ceil(float64(items) / float64(limit))
	}

	return int(math.Ceil(totalPages))
}
