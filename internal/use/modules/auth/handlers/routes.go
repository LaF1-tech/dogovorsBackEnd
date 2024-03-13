package handlers

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/http/middlewares"
)

func (h *handler) Routes(group *gin.RouterGroup, m *http.Middlewares) {
	m.Auth = middlewares.Auth(h.controller)
	usersGroup := group.Group("users")

	usersGroup.POST("signup", m.Auth, h.SignUp)
	usersGroup.POST("login", h.Login)
	usersGroup.GET("", m.Auth, h.GetSelfUser)
	usersGroup.PATCH("", m.Auth, h.PatchUser)
	usersGroup.DELETE(":id", m.Auth, h.DeleteUser)
}
