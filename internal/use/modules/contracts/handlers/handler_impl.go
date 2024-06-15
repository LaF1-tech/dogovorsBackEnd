package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"dogovorsBackEnd/internal/use/utils"
)

func (h *handler) GetAllContracts(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	response, err := h.controller.GetAllContracts(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) GetContractByID(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	iddStr := ctx.Param("id")
	id, err := strconv.Atoi(iddStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	preview, err := h.controller.GetContractByID(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, preview)
}

func (h *handler) GenerateContractPDF(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	iddStr := ctx.Param("id")

	id, err := strconv.Atoi(iddStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	filename, err := h.controller.GenerateContractPDF(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.File(filename)
}
