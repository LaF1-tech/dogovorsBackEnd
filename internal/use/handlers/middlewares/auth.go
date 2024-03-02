package middlewares

import (
	"context"
	"dogovorsBackEnd/internal/use/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type IAuth interface {
	Auth(ctx context.Context, token string) (models.User, error)
}

func Auth(ctrl IAuth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			_ = ctx.Error(fmt.Errorf("no auth header provided: %w", ErrForbidden))
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			_ = ctx.Error(fmt.Errorf("could not find bearer token: %w", ErrForbidden))
			ctx.Abort()
			return
		}

		user, err := ctrl.Auth(ctx, token)
		if err != nil {
			_ = ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
	}
}
