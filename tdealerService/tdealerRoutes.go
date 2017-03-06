package tdealerService

import (
	"bitbucket.org/tekion/tbaas/tapi"
)

func Start() {
	tapi.AddRoutes(
		"Login",
		"GET",
		"/getDealerById",
		GetDealerById,
		tapi.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
}
