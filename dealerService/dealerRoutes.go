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
	"bitbucket.org/tekion/tacl/acl"
	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/tapi"
)

//TODO : Need new admin scope

// Start add routes and start the service at specified port
func Start() {
	tapi.AddRoutes(
		"readDealer",
		"GET",
		"/dealer",
		readDealer,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readFixedOperation",
		"GET",
		"/fixedoperation",
		readFixedOperation,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerContact",
		"GET",
		"/contact/{cid}",
		readDealerContact,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerContacts",
		"GET",
		"/contacts",
		readDealerContacts,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGoal",
		"GET",
		"/goal/{gid}",
		readDealerGoal,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGoals",
		"GET",
		"/goals",
		readDealerGoals,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGroups",
		"GET",
		"/groups",
		readDealerGroups,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"dealerList",
		"Post",
		"/dealer",
		 dealerList,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor"},
		},
	)
	tapi.AddRoutes(
		"updateDealer",
		"MethodPatch",
		"/dealer",
		 updateDealer,
		acl.ACLStruct{
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor"},
		},
	)
	//log service start info
	log.GenericInfo("", "", "", "Started Tekion tdealer on port:8079")
	tapi.Start("8079", "/tdealer")
}
