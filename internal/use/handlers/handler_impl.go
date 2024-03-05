package handlers

import (
	"dogovorsBackEnd/internal/use/dto"
	"dogovorsBackEnd/internal/use/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) SignUp(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	var request dto.SignUpRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.SignUp(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *handler) Login(ctx *gin.Context) {
	var request dto.LoginRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.Login(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) PatchUser(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	var request dto.UpdateRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.PatchUser(ctx, user, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) DeleteUser(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	id, err := utils.IntParam(ctx, "id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = h.controller.DeleteUserByID(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
