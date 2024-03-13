package types

import (
	"dogovorsBackEnd/internal/use/http"
	auth "dogovorsBackEnd/internal/use/modules/auth/controller"
	templates "dogovorsBackEnd/internal/use/modules/templates/controller"
)

type Controllers struct {
	auth.AuthController
	templates.TemplatesController
}

type Routes []http.IRoutes

type Runnable interface {
	Run(addr ...string) error
}
