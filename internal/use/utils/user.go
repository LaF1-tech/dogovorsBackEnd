package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"dogovorsBackEnd/internal/use/models"
)

func GetUser(ctx *gin.Context) models.User {
	p, _ := ctx.Get("user")
	return p.(models.User)
}

func IntParam(ctx *gin.Context, key string) (int, error) {
	p := ctx.Param(key)
	return strconv.Atoi(p)
}
