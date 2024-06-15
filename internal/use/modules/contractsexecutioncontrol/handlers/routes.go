package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, m *http.Middlewares) {
	contractExecutionControlGroup := group.Group("cec")

	contractExecutionControlGroup.GET(":id", m.Auth, h.GetContractsExecutionControlByContractID)
	contractExecutionControlGroup.PATCH("", m.Auth, h.PatchContractExecutionControl)
	contractExecutionControlGroup.POST("", m.Auth, h.CreateContractExecutionControl)

}
