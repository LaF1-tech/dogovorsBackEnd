package handlers

import (
	"dogovorsBackEnd/internal/use/handlers/middlewares"
	"github.com/gin-gonic/gin"
)

func (h *handler) Routes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Error())
	v1 := r.Group("/api/v1")

	auth := middlewares.Auth(h.controller)

	{
		usersGroup := v1.Group("users")

		usersGroup.POST("signup", auth, h.SignUp)
		usersGroup.POST("login", h.Login)
		usersGroup.PATCH("", auth, h.PatchUser)
		usersGroup.DELETE(":id", auth, h.DeleteUser)
	}

	return r
}
