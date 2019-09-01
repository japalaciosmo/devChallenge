package app

import (
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/kyani-inc/kms"
	"github.com/kyani-inc/kms-object-pro/src/app/log"
	"github.com/labstack/echo"
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

// ParseSNS parses the sns payload data into a target data structure.
func ParseSNS(ctx echo.Context, dest interface{}) error {
	sns := SNSPayload{}

	if err := ctx.Bind(&sns); err != nil {
		log.With(kms.Log{"SNSPayload": sns}).Error(fmt.Sprintf("Failed to bind request payload. Error: %s. Content Type: %s\n", err.Error(), ctx.Request().Header.Get("Content-Type")))
		return err
	}

	if err := json.Unmarshal([]byte(fmt.Sprintf("%v", sns.Message)), &dest); err != nil {
		log.With(kms.Log{"dest": dest}).Error(err)
		return err
	}
	return nil
}

