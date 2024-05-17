package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/controller"
)

type EducationalEstablishmentsHandler interface {
	http.IRoutes
	CreateEducationalEstablishment(ctx *gin.Context)
	GetSpecializations(ctx *gin.Context)
	GetEducationalEstablishmentByID(ctx *gin.Context)
	PatchEducationalEstablishment(ctx *gin.Context)
	DeleteEducationalEstablishmentByID(ctx *gin.Context)
}

type handler struct {
	controller controller.EducationalEstablishmentsController
}

func New(controller controller.EducationalEstablishmentsController) EducationalEstablishmentsHandler {
	return &handler{controller: controller}
}
