package http

import (
	"github.com/gin-gonic/gin"

	"dogovorsBackEnd/internal/use/http/middlewares"
)

func ConfigureRoutes(routes ...IRoutes) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Error())
	r.Use(middlewares.CORSMiddleware())
	v1 := r.Group("/api/v1")

	var m Middlewares

	for _, route := range routes {
		route.Routes(v1, &m)
	}

	return r
}
