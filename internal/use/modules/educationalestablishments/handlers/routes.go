package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
)

func (h *handler) Routes(group *gin.RouterGroup, middlewares *http.Middlewares) {
	educationalEstablishmentsGroup := group.Group("ee")

	educationalEstablishmentsGroup.POST("", middlewares.Auth, h.CreateEducationalEstablishment)
	educationalEstablishmentsGroup.GET("", h.GetSpecializations)
	educationalEstablishmentsGroup.GET(":id", h.GetEducationalEstablishmentByID)
	educationalEstablishmentsGroup.PATCH("", middlewares.Auth, h.PatchEducationalEstablishment)
	educationalEstablishmentsGroup.DELETE(":id", middlewares.Auth, h.DeleteEducationalEstablishmentByID)
}
