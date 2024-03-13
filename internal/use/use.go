package use

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/http"
	"dogovorsBackEnd/internal/use/modules"
	"dogovorsBackEnd/internal/use/types"
)

func New(database *sql.DB) (types.Runnable, types.Controllers) {
	routes, ctrls := modules.New(database)

	runnable := http.ConfigureRoutes(routes...)

	return runnable, ctrls
}
