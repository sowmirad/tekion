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

	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/tapi"
)

const (
	dealerResourceName        = "Dealer"
	fixedOpResourceName       = "Fixed_Operations"
	dealerContactResourceName = "Dealer_Contact"
	dealerGoalResourceName    = "Dealer_Goal"
	dealerGroupResourceName   = "Dealer_Group"
	read                      = uint64(1)
	create                    = uint64(2)
	update                    = uint64(4)
	delete                    = uint64(8)
)

//TODO : Need new admin scope

// Start add routes and start the service at specified port
func Start() {

	tapi.AddRoute(
		"readDealer",
		http.MethodGet,
		"/dealer",
		[]string{dealerResourceName},
		read,
		readDealerH,
	)

	tapi.AddRoute(
		"dealersList",
		http.MethodPost,
		"/dealers",
		[]string{dealerResourceName},
		read,
		dealersListH,
	)

	tapi.AddRoute(
		"patchDealer",
		http.MethodPatch,
		"/dealer",
		[]string{dealerResourceName},
		update,
		patchDealerH,
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
	tapi.AddRoute(
		"saveDealer",
		http.MethodPost,
		"/dealer",
		[]string{dealerResourceName},
		create+update,
		saveDealerH,
	)

	tapi.AddRoute(
		"readFixedOperation",
		http.MethodGet,
		"/fixedoperation",
		[]string{fixedOpResourceName},
		read,
		readFixedOperationH,
	)

	tapi.AddRoute(
		"patchFixedOperation",
		http.MethodPatch,
		"/fixedoperation",
		[]string{fixedOpResourceName},
		update,
		patchFixedOperationH,
	)

	tapi.AddRoute(
		"readDealerContact",
		http.MethodGet,
		"/contact/{cid}",
		[]string{dealerContactResourceName},
		read,
		readDealerContactH,
	)

	tapi.AddRoute(
		"readDealerContacts",
		http.MethodGet,
		"/contacts",
		[]string{dealerContactResourceName},
		read,
		readDealerContactsH,
	)

	tapi.AddRoute(
		"readDealerGoal",
		http.MethodGet,
		"/goal/{gid}",
		[]string{dealerGoalResourceName},
		read,
		readDealerGoalH,
	)
	tapi.AddRoute(
		"readDealerGoals",
		http.MethodGet,
		"/goals",
		[]string{dealerGoalResourceName},
		read,
		readDealerGoalsH,
	)

	tapi.AddRoute(
		"readDealerGroups",
		http.MethodGet,
		"/groups",
		[]string{dealerGroupResourceName},
		read,
		readDealerGroupsH,
	)

	tapi.AddRoute(
		"readDealerGroups",
		http.MethodGet,
		"/aggregate/dealer/fixedoperation",
		[]string{dealerResourceName, fixedOpResourceName},
		read,
		aggregateDealerFixedOpH,
	)

	//log service start info
	log.GenericInfo("", "", "", "Started Tekion tdealer on port:8079")
	tapi.Start("8079", "/tdealer")
}
