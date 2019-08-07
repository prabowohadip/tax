package config

import (
	"github.com/labstack/echo"
	"net/http"
)

func SuccessResponse(ctx echo.Context, data interface{}) error {

	responseData := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func ErrorResponse(ctx echo.Context, err error) error {

	responseData := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}
