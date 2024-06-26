package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, m *http.Middlewares) {
	contractsGroup := group.Group("contracts")
	contractsGroup.GET("", m.Auth, h.GetAllContracts)
	contractsGroup.GET(":id", m.Auth, h.GetContractByID)
	contractsGroup.GET("pdf/:id", m.Auth, h.GenerateContractPDF)

}
