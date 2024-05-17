package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/specializations/controller"
)

type SpecializationsHandler interface {
	http.IRoutes
	CreateSpecialization(ctx *gin.Context)
	GetSpecializations(ctx *gin.Context)
	GetSpecializationByID(ctx *gin.Context)
	PatchSpecialization(ctx *gin.Context)
	DeleteSpecializationByID(ctx *gin.Context)
}

type handler struct {
	controller controller.SpecializationsController
}

func New(controller controller.SpecializationsController) SpecializationsHandler {
	return &handler{controller: controller}
}
