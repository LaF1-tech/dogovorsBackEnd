package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"dogovorsBackEnd/internal/use/modules/applications/dto"
	"dogovorsBackEnd/internal/use/utils"
)

func (h *handler) CreateApplication(ctx *gin.Context) {
	var request dto.CreateApplicationRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateApplication(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response)

}

func (h *handler) GetApplications(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	response, err := h.controller.GetAllApplications(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) PatchApplicationByID(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.PatchApplicationRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.PatchApplicationByID(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
