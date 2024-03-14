package modules

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules/applications"
	"dogovorsBackEnd/internal/use/modules/auth"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments"
	"dogovorsBackEnd/internal/use/modules/specializations"
	"dogovorsBackEnd/internal/use/modules/templates"
	"dogovorsBackEnd/internal/use/types"
)

func New(database *sql.DB) (types.Routes, types.Controllers) {
	routes := make([]http.IRoutes, 5)
	var controllers types.Controllers

	routes[0], controllers.AuthController = auth.New(database)
	routes[1], controllers.TemplatesController = templates.New(database)
	routes[2], controllers.ApplicationsController = applications.New(database)
	routes[3], controllers.EducationalEstablishmentsController = educationalestablishments.New(database)
	routes[4], controllers.SpecializationsController = specializations.New(database)

	return routes, controllers
}
