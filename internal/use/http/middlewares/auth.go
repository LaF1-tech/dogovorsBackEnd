package middlewares

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/models"
)

var (
	ErrForbidden    = errors.New("forbidden")
	ErrUnauthorized = errors.New("unauthorized")
	ErrValidation   = errors.New("validation error")
	ErrNotFound     = errors.New("not found")
)

type IAuth interface {
	Auth(ctx context.Context, token string) (models.User, error)
}

func Auth(ctrl IAuth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			_ = ctx.Error(ErrUnauthorized)
			ctx.Abort()
			return
		}

		user, err := ctrl.Auth(ctx, cookie)
		if err != nil {
			_ = ctx.Error(ErrUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
	}
}
