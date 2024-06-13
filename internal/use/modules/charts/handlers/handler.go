package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/charts/controller"
)

type ChartsHandler interface {
	http.IRoutes

	GetPeriodChart(ctx *gin.Context)
	GetPeriodUserChart(ctx *gin.Context)
	GetEducationalEstablishmentChart(ctx *gin.Context)
	GetSpecializationsChart(ctx *gin.Context)
	GetTemplatesChart(ctx *gin.Context)
}

type handler struct {
	controller controller.ChartsController
}

func New(controller controller.ChartsController) ChartsHandler {
	return &handler{controller: controller}
}
