package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, m *http.Middlewares) {
	chartsGroup := group.Group("charts")
	chartsGroup.GET("/period", m.Auth, h.GetPeriodChart)
	chartsGroup.GET("/educationalestablishments", m.Auth, h.GetEducationalEstablishmentChart)
	chartsGroup.GET("/specializations", m.Auth, h.GetSpecializationsChart)
	chartsGroup.GET("/templates", m.Auth, h.GetTemplatesChart)

}
