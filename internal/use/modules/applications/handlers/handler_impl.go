package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"dogovorsBackEnd/internal/use/modules/applications/dto"
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
