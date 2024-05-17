package handlers

import (
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"dogovorsBackEnd/internal/use/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handler) CreateEducationalEstablishment(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	var request dto.CreateEducationalEstablishmentRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateEducationalEstablishment(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *handler) GetSpecializations(ctx *gin.Context) {
	ee, err := h.controller.GetEducationalEstablishments(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, ee)
}

func (h *handler) GetEducationalEstablishmentByID(ctx *gin.Context) {
	iddStr := ctx.Param("id")
	id, err := strconv.Atoi(iddStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ee, err := h.controller.GetEducationalEstablishmentByID(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ee)
}

func (h *handler) PatchEducationalEstablishment(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	var request dto.PatchEducationalEstablishmentRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.PatchEducationalEstablishmentByID(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *handler) DeleteEducationalEstablishmentByID(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = h.controller.DeleteEducationalEstablishmentByID(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
