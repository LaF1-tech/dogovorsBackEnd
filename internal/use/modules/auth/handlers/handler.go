package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/modules/auth/controller"

	"dogovorsBackEnd/internal/use/http"
)

type AuthHandler interface {
	http.IRoutes

	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetSelfUser(ctx *gin.Context)
	PatchUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type handler struct {
	controller controller.AuthController
}

func New(controller controller.AuthController) AuthHandler {
	return &handler{controller: controller}
}
