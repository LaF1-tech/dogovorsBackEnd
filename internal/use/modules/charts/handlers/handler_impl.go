package handlers

import (
	"dogovorsBackEnd/internal/use/modules/charts/dto"
	"dogovorsBackEnd/internal/use/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) GetPeriodChart(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	dataForChart, err := h.controller.GetPeriodChart(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dataForChart)
}

func (h *handler) GetPeriodUserChart(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.PeriodUserChartRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	dataForChart, err := h.controller.GetPeriodUserChart(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dataForChart)
}

func (h *handler) GetEducationalEstablishmentChart(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	dataForChart, err := h.controller.GetEducationalEstablishmentChart(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dataForChart)
}

func (h *handler) GetSpecializationsChart(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	dataForChart, err := h.controller.GetSpecializationsChart(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dataForChart)
}

func (h *handler) GetTemplatesChart(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	dataForChart, err := h.controller.GetTemplatesChart(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dataForChart)
}
