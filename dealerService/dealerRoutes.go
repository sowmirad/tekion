// Package dealerService implements dealer micro service
// title: Dealer micro service
//
// dealerService implements dealer micro service
//
// The purpose of this application is to provides api's to perform CURD operations on dealer object.
// Currently only get endpoints are available.
// dealerService is divided into 4 file.
//  1. dealerRoutes.go  -> contain routes.
//  2. dealerHandler.go -> containing handler functions.
//  3. dealerModel.go   -> containing models.
//  4. dealerUtils.go   -> containing util functions.
//
// Terms Of Service:
//
//     Schemes: https
//     BasePath: /tdealer
//     Version: 1.0.0
//     Contact: Qasim Hasnain<mqhasnain@tekion.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package dealerService

import (
	"net/http"

	"bitbucket.org/tekion/tacl/acl"
	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/tapi"
)

//TODO : Need new admin scope

// Start add routes and start the service at specified port
func Start() {
	tapi.AddRoutes(
		"readDealer",
		http.MethodGet,
		"/dealer",
		readDealer,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor", "ServiceDirector", "Technician", "Dispatcher", "BDCSpecialist"},
		},
	)
	tapi.AddRoutes(
		"dealersList",
		http.MethodPost,
		"/dealers",
		dealersList,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor"},
		},
	)
	tapi.AddRoutes(
		"patchDealer",
		http.MethodPatch,
		"/dealer",
		patchDealer,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor"},
		},
	)
	// todo create and update should be one function. Figure out why and write one
	/*	tapi.AddRoutes(
		"createDealer",
		http.MethodPost,
		"/createDealer",
		createDealer,
		acl.ACLStruct{
			// TODO PremittedRoles (SuperAdmin)
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor"},
		},
	)*/
	tapi.AddRoutes(
		"saveDealer",
		http.MethodPost,
		"/dealer",
		saveDealer,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor"},
		},
	)
	tapi.AddRoutes(
		"readFixedOperation",
		http.MethodGet,
		"/fixedoperation",
		readFixedOperation,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher", "BDCSpecialist"},
		},
	)
	tapi.AddRoutes(
		"patchFixedOperation",
		http.MethodPatch,
		"/fixedoperation",
		patchFixedOperation,
		acl.ACLStruct{
			// TODO PremittedRoles (SuperAdmin)
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor"},
		},
	)
	tapi.AddRoutes(
		"readDealerContact",
		http.MethodGet,
		"/contact/{cid}",
		readDealerContact,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerContacts",
		http.MethodGet,
		"/contacts",
		readDealerContacts,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGoal",
		http.MethodGet,
		"/goal/{gid}",
		readDealerGoal,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGoals",
		http.MethodGet,
		"/goals",
		readDealerGoals,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGroups",
		http.MethodGet,
		"/groups",
		readDealerGroups,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGroups",
		http.MethodGet,
		"/aggregate/dealer/fixedoperation",
		aggregateDealerFixedOp,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)

	//log service start info
	log.GenericInfo("", "", "", "Started Tekion tdealer on port:8079")
	tapi.Start("8079", "/tdealer")
}
