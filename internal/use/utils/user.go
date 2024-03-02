package utils

import (
	"dogovorsBackEnd/internal/use/models"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) models.User {
	p, _ := ctx.Get("user")
	return p.(models.User)
}
