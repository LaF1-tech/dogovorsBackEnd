package handlers

import (
	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"dogovorsBackEnd/internal/use/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handler) CreateSpecialization(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.CreateSpecializationRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateSpecialization(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) GetSpecializations(ctx *gin.Context) {
	ee, err := h.controller.GetSpecializations(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, ee)
}

func (h *handler) GetSpecializationByID(ctx *gin.Context) {
	iddStr := ctx.Param("id")
	id, err := strconv.Atoi(iddStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	specialization, err := h.controller.GetSpecializationByID(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, specialization)
}

func (h *handler) PatchSpecialization(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.PatchSpecializationRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.PatchSpecializationByID(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *handler) DeleteSpecializationByID(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = h.controller.DeleteSpecializationByID(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
