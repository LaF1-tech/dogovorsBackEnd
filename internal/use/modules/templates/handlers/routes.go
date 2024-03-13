package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, _ *http.Middlewares) {
	templatesGroup := group.Group("templates")

	templatesGroup.GET(":id", h.GetTemplateByID)
	templatesGroup.GET("", h.GetAllTemplatesPreview)
}
