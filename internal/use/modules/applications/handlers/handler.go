package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/applications/controller"
)

type ApplicationsHandler interface {
	http.IRoutes

	CreateApplication(ctx *gin.Context)
}

type handler struct {
	controller controller.ApplicationsController
}

func New(controller controller.ApplicationsController) ApplicationsHandler {
	return &handler{controller: controller}
}
