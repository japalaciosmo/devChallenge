package app

import (
	"github.com/labstack/echo"
	"net/http"
)

// OK responds with a 200 ok JSON packet.
func OK(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusOK, Confirmation{Message: msg})
}

// Success responds with a 200 ok JSON packet.
func Success(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

// Error responds with a 500 internal server error.
func Error(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		Status:  http.StatusInternalServerError,
		Details: err.Error(),
	})
}

// BadRequest responds with a 400 bad request error.
func BadRequest(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, ErrorResponse{
		Status:  http.StatusBadRequest,
		Details: err.Error(),
	})
}
