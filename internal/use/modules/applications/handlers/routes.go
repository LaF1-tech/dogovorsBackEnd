package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, m *http.Middlewares) {
	applicationsGroup := group.Group("applications")

	applicationsGroup.POST("create", h.CreateApplication)
	applicationsGroup.GET("", m.Auth, h.GetApplications)
	applicationsGroup.PATCH("", m.Auth, h.PatchApplicationByID)
}
