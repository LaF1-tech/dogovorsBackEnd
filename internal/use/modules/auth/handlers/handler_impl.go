package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"dogovorsBackEnd/internal/use/modules/auth/dto"

	"dogovorsBackEnd/internal/use/utils"
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

	ctx.SetCookie("token", response.Token, 3600, "/", "", false, false)
	ctx.JSON(http.StatusOK, response)
}

func (h *handler) GetSelfUser(ctx *gin.Context) {
	user := utils.GetUser(ctx)
	ctx.JSON(http.StatusOK, user)
}

func (h *handler) PatchUser(ctx *gin.Context) {
	user := utils.GetUser(ctx)

	var request dto.UserUpdateRequestDTO
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
