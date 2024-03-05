package utils

import (
	"dogovorsBackEnd/internal/use/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(ctx *gin.Context) models.User {
	p, _ := ctx.Get("user")
	return p.(models.User)
}

func IntParam(ctx *gin.Context, key string) (int, error) {
	p := ctx.Param(key)
	return strconv.Atoi(p)
}
