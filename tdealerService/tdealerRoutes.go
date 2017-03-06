package tdealerService

import (
	"bitbucket.org/tekion/tbaas/tapi"
)

//Start : function to start route
func Start() {
	tapi.AddRoutes(
		"Login",
		"GET",
		"/getDealerById",
		GetDealerByID,
		tapi.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
}
