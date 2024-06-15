package handlers

import (
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/dto"
	"dogovorsBackEnd/internal/use/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handler) GetContractsExecutionControlByContractID(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	iddStr := ctx.Param("id")
	id, err := strconv.Atoi(iddStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cecs, err := h.controller.GetContractsExecutionControlByContractID(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, cecs)
}

func (h *handler) CreateContractExecutionControl(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	var request dto.CreateContractExecutionControlRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateContractExecutionControl(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *handler) PatchContractExecutionControl(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	var request dto.PatchContractExecutionControlRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.PatchContractExecutionControl(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
