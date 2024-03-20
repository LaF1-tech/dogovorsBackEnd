package types

import (
	"dogovorsBackEnd/internal/use/http"
	applications "dogovorsBackEnd/internal/use/modules/applications/controller"
	auth "dogovorsBackEnd/internal/use/modules/auth/controller"
	contracts "dogovorsBackEnd/internal/use/modules/contracts/controller"
	ee "dogovorsBackEnd/internal/use/modules/educationalestablishments/controller"
	specializations "dogovorsBackEnd/internal/use/modules/specializations/controller"
	templates "dogovorsBackEnd/internal/use/modules/templates/controller"
)

type Controllers struct {
	auth.AuthController
	templates.TemplatesController
	applications.ApplicationsController
	ee.EducationalEstablishmentsController
	specializations.SpecializationsController
	contracts.ContractsController
}

type Routes []http.IRoutes

type Runnable interface {
	Run(addr ...string) error
}
