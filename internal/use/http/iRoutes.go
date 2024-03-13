package http

import (
	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	Auth gin.HandlerFunc
}

type IRoutes interface {
	Routes(group *gin.RouterGroup, middlewares *Middlewares)
}
