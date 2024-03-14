package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) GetEducationalEstablishments(ctx *gin.Context) {
	ee, err := h.controller.GetSpecializations(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, ee)
}
