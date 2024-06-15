package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/contracts/controller"
)

type ContractsHandler interface {
	http.IRoutes

	GetAllContracts(ctx *gin.Context)
	GetContractByID(ctx *gin.Context)

	GenerateContractPDF(ctx *gin.Context)
}

type handler struct {
	controller controller.ContractsController
}

func New(controller controller.ContractsController) ContractsHandler {
	return &handler{controller: controller}
}
