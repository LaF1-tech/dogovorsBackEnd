package handlers

import (
	"dogovorsBackEnd/internal/use/modules/templates/dto"
	"dogovorsBackEnd/internal/use/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handler) CreateTemplate(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.CreateTemplateRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateTemplate(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) GetAllTemplatesPreview(ctx *gin.Context) {
	preview, err := h.controller.GetTemplatesPreview(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, preview)
}

func (h *handler) GetTemplateByID(ctx *gin.Context) {
	iddStr := ctx.Param("id")
	id, err := strconv.Atoi(iddStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	preview, err := h.controller.GetTemplateByID(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, preview)
}

func (h *handler) PatchTemplateByID(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.PatchTemplateRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.PatchTemplateByID(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *handler) DeleteTemplateByID(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	id, err := utils.IntParam(ctx, "id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = h.controller.DeleteTemplateByID(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
