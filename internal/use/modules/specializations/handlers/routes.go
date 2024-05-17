package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, middlewares *http.Middlewares) {
	specializationsGroup := group.Group("specializations")

	specializationsGroup.POST("", middlewares.Auth, h.CreateSpecialization)
	specializationsGroup.GET("", h.GetSpecializations)
	specializationsGroup.GET(":id", h.GetSpecializationByID)
	specializationsGroup.PATCH("", middlewares.Auth, h.PatchSpecialization)
	specializationsGroup.DELETE(":id", middlewares.Auth, h.DeleteSpecializationByID)
}
