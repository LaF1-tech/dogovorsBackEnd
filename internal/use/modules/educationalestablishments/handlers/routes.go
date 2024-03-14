package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, middlewares *http.Middlewares) {
	applicationsGroup := group.Group("ee")

	applicationsGroup.GET("", h.GetEducationalEstablishments)
}
