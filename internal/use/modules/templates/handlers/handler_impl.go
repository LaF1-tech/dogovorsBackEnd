package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
