package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/templates/controller"
)

type TemplateHandler interface {
	http.IRoutes

	GetAllTemplatesPreview(ctx *gin.Context)
	GetTemplateByID(ctx *gin.Context)
}

type handler struct {
	controller controller.TemplatesController
}

func New(controller controller.TemplatesController) TemplateHandler {
	return &handler{controller: controller}
}
