package tdealerService

import (
	"bitbucket.org/tekion/tbaas/tapi"
)

//Start : function to start route
func Start() {
	tapi.AddRoutes(
		"API to get dealer by dealer ID",
		"GET",
		"/getDealerById",
		GetDealerByID,
		tapi.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
}
