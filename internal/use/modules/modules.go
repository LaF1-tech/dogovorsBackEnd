package modules

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/auth"
	"dogovorsBackEnd/internal/use/modules/templates"
	"dogovorsBackEnd/internal/use/types"
)

func New(database *sql.DB) (types.Routes, types.Controllers) {
	routes := make([]http.IRoutes, 2)
	var controllers types.Controllers

	routes[0], controllers.AuthController = auth.New(database)
	routes[1], controllers.TemplatesController = templates.New(database)

	return routes, controllers
}
