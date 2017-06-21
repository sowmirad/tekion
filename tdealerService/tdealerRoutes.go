package tdealerService

import (
	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/tapi"
)

//Start : function to start route
func Start() {
	tapi.AddRoutes(
		"API to get dealer by dealer ID",
		"GET",
		"/getDealerByID",
		GetDealerByID,
		tapi.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)

	tapi.AddRoutes(
		"API to get dealer by dealer ID",
		"GET",
		"/selectdamage",
		getDamageTypes,
		tapi.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)

	log.GenericInfo("", "", "", "tdealer service started on port :8079")
	tapi.Start("8079", "/tdealer")
}
