package handlers

import (
	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/controller"
	"github.com/gin-gonic/gin"
)

type ContractsExecutionControlHandler interface {
	http.IRoutes

	GetContractsExecutionControlByContractID(ctx *gin.Context)
	CreateContractExecutionControl(ctx *gin.Context)
	PatchContractExecutionControl(ctx *gin.Context)
}

type handler struct {
	controller controller.ContractsExecutionControlController
}

func New(controller controller.ContractsExecutionControlController) ContractsExecutionControlHandler {
	return &handler{controller: controller}
}
