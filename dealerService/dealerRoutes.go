package dealerService

import (
	"bitbucket.org/tekion/tacl/acl"
	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/tapi"
)

//TODO : Need new admin scope

// Start add routes and start the service at specified port
func Start() {
	tapi.AddRoutes(
		"Read dealer info",
		"GET",
		"/dealer",
		ReadDealer,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads fixedoperation",
		"GET",
		"/fixedoperation/{id}",
		ReadFixedOperation,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads fixedoperations",
		"GET",
		"/fixedoperations",
		ReadFixedOperations,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads contact",
		"GET",
		"/contact/{id}",
		ReadContact,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads contacts",
		"GET",
		"/contacts",
		ReadContacts,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads goal",
		"GET",
		"/goal/{id}",
		ReadGoal,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads goals",
		"GET",
		"/goals",
		ReadGoals,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"Reads groups",
		"GET",
		"/groups",
		ReadGroups,
		acl.ACLStruct{
			PermittedRoles: []string{"ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	//log service start info
	log.GenericInfo("", "", "", "Started Tekion tdealer on port:8079")
	tapi.Start("8079", "/tdealer")
}
