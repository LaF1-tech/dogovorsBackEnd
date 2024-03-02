package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err == nil {
			return
		}

		status := errorToStatusMapper(err)
		ctx.JSON(status, err.JSON())
	}
}

func errorToStatusMapper(err error) (status int) {
	switch {
	case errors.Is(err, ErrForbidden):
		status = http.StatusForbidden
	default:
		status = http.StatusInternalServerError
	}

	return
}