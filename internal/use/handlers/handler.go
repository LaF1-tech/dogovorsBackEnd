package handlers

import (
	"dogovorsBackEnd/internal/use/controllers"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Routes() *gin.Engine

	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type handler struct {
	controller controllers.Controller
}

func New(controller controllers.Controller) Handler {
	return &handler{controller: controller}
}
