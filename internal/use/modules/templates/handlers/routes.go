package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, m *http.Middlewares) {
	templatesGroup := group.Group("templates")

	templatesGroup.POST("", m.Auth, h.CreateTemplate)
	templatesGroup.GET(":id", h.GetTemplateByID)
	templatesGroup.GET("", h.GetAllTemplatesPreview)
	templatesGroup.PATCH("", m.Auth, h.PatchTemplateByID)
	templatesGroup.DELETE(":id", m.Auth, h.DeleteTemplateByID)
}
