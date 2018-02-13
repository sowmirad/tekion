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
	"bitbucket.org/tekion/tbaas/apiContext"

	log "bitbucket.org/tekion/tbaas/log/v1"
	"bitbucket.org/tekion/tbaas/tapi"
)

//TODO : Need new admin scope

// Start add routes and start the service at specified port
func Start() {

	//tapi.AddRoute(
	//	"readDealer",
	//	http.MethodGet,
	//	"/dealer",
	//	[]string{com.DealerResourceName},
	//	[]uint64{com.Read},
	//	readDealerH,
	//)
	//
	//tapi.AddRoute(
	//	"dealersList",
	//	http.MethodPost,
	//	"/dealers",
	//	[]string{com.DealerResourceName},
	//	[]uint64{com.Read},
	//	dealersListH,
	//)
	//
	//tapi.AddRoute(
	//	"patchDealer",
	//	http.MethodPatch,
	//	"/dealer",
	//	[]string{com.DealerResourceName},
	//	[]uint64{com.Update},
	//	patchDealerH,
	//)
	//// todo create and update should be one function. Figure out why and write one
	///*	tapi.AddRoutes(
	//	"createDealer",
	//	http.MethodPost,
	//	"/createDealer",
	//	createDealer,
	//	acl.ACLStruct{
	//		// TODO PremittedRoles (SuperAdmin)
	//		PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "ServiceDirector"},
	//	},
	//)*/
	//tapi.AddRoute(
	//	"saveDealer",
	//	http.MethodPost,
	//	"/dealer",
	//	[]string{com.DealerResourceName},
	//	[]uint64{com.Create + com.Update},
	//	saveDealerH,
	//)
	//
	//tapi.AddRoute(
	//	"readFixedOperation",
	//	http.MethodGet,
	//	"/fixedoperation",
	//	[]string{com.FixedOperationResourceName},
	//	[]uint64{com.Read},
	//	readFixedOperationH,
	//)
	//
	//tapi.AddRoute(
	//	"patchFixedOperation",
	//	http.MethodPatch,
	//	"/fixedoperation",
	//	[]string{com.FixedOperationResourceName},
	//	[]uint64{com.Update},
	//	patchFixedOperationH,
	//)
	//
	//tapi.AddRoute(
	//	"readDealerContact",
	//	http.MethodGet,
	//	"/contact/{cid}",
	//	[]string{com.ContactResourceName},
	//	[]uint64{com.Read},
	//	readDealerContactH,
	//)
	//
	//tapi.AddRoute(
	//	"readDealerContacts",
	//	http.MethodGet,
	//	"/contacts",
	//	[]string{com.ContactResourceName},
	//	[]uint64{com.Read},
	//	readDealerContactsH,
	//)
	//
	//tapi.AddRoute(
	//	"readDealerGoal",
	//	http.MethodGet,
	//	"/goal/{gid}",
	//	[]string{com.GoalResourceName},
	//	[]uint64{com.Read},
	//	readDealerGoalH,
	//)
	//tapi.AddRoute(
	//	"readDealerGoals",
	//	http.MethodGet,
	//	"/goals",
	//	[]string{com.GoalResourceName},
	//	[]uint64{com.Read},
	//	readDealerGoalsH,
	//)
	//
	//tapi.AddRoute(
	//	"readDealerGroups",
	//	http.MethodGet,
	//	"/groups",
	//	[]string{com.GroupResourceName},
	//	[]uint64{com.Read},
	//	readDealerGroupsH,
	//)
	//
	//tapi.AddRoute(
	//	"readDealerGroups",
	//	http.MethodGet,
	//	"/aggregate/dealer/fixedoperation",
	//	[]string{com.DealerResourceName, com.FixedOperationResourceName},
	//	[]uint64{com.Read},
	//	aggregateDealerFixedOpH,
	//)

	tapi.AddRoutes(
		"readDealer",
		http.MethodGet,
		"/dealer",
		readDealerH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor", "ServiceDirector", "Technician", "Dispatcher", "BDCSpecialist", "PartsClerk"},
		},
	)
	tapi.AddRoutes(
		"dealersList",
		http.MethodPost,
		"/dealers",
		dealersListH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor", "ServiceDirector"},
		},
	)
	tapi.AddRoutes(
		"patchDealer",
		http.MethodPatch,
		"/dealer",
		patchDealerH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor", "ServiceDirector"},
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
			PermittedRoles: []string{"SystemUser", "ServiceAdvisor", "ServiceDirector"},
		},
	)*/
	tapi.AddRoutes(
		"saveDealer",
		http.MethodPost,
		"/dealer",
		saveDealerH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor", "ServiceDirector"},
		},
	)
	tapi.AddRoutes(
		"readFixedOperation",
		http.MethodGet,
		"/fixedoperation",
		readFixedOperationH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher", "BDCSpecialist", "PartsClerk"},
		},
	)
	tapi.AddRoutes(
		"patchFixedOperation",
		http.MethodPatch,
		"/fixedoperation",
		patchFixedOperationH,
		acl.ACLStruct{
			// TODO PremittedRoles (SuperAdmin)
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "SystemAdmin", "ServiceAdvisor", "ServiceDirector"},
		},
	)
	tapi.AddRoutes(
		"readDealerContact",
		http.MethodGet,
		"/contact/{cid}",
		readDealerContactH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerContacts",
		http.MethodGet,
		"/contacts",
		readDealerContactsH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGoal",
		http.MethodGet,
		"/goal/{gid}",
		readDealerGoalH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGoals",
		http.MethodGet,
		"/goals",
		readDealerGoalsH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGroups",
		http.MethodGet,
		"/groups",
		readDealerGroupsH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)
	tapi.AddRoutes(
		"readDealerGroups",
		http.MethodGet,
		"/aggregate/dealer/fixedoperation",
		aggregateDealerFixedOpH,
		acl.ACLStruct{
			PermittedRoles: []string{"Accountant", "Manager", "SystemUser", "ServiceAdvisor", "ServiceDirector", "SystemAdmin", "Technician", "Dispatcher"},
		},
	)

	//log service start info
	log.GenericInfo(apiContext.TContext{}, "Started Tekion tdealer on port:8079", nil)
	tapi.Start("8079", "/tdealer")
}
