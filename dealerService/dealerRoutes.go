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

	"bitbucket.org/tekion/tbaas/apiContext"
	com "bitbucket.org/tekion/tenums/common"

	log "bitbucket.org/tekion/tbaas/log/v1"
	"bitbucket.org/tekion/tbaas/tapi"
)

//TODO : Need new admin scope

// Start add routes and start the service at specified port
func Start() {

	tapi.AddRoute(
		"readDealer",
		http.MethodGet,
		"/dealer",
		map[string]uint8{com.DealerResourceName: com.Read},
		readDealerH,
	)

	tapi.AddRoute(
		"dealersList",
		http.MethodPost,
		"/dealers",
		map[string]uint8{com.DealerResourceName: com.Read},
		dealersListH,
	)

	tapi.AddRoute(
		"patchDealer",
		http.MethodPatch,
		"/dealer",
		map[string]uint8{com.DealerResourceName: com.Update},
		patchDealerH,
	)

	tapi.AddRoute(
		"saveDealer",
		http.MethodPost,
		"/dealer",
		map[string]uint8{com.DealerResourceName: com.Create + com.Update},
		saveDealerH,
	)

	tapi.AddRoute(
		"readFixedOperation",
		http.MethodGet,
		"/fixedoperation",
		map[string]uint8{com.FixedOperationResourceName: com.Read},
		readFixedOperationH,
	)

	tapi.AddRoute(
		"patchFixedOperation",
		http.MethodPatch,
		"/fixedoperation",
		map[string]uint8{com.FixedOperationResourceName: com.Update},
		patchFixedOperationH,
	)

	tapi.AddRoute(
		"readDealerGroups",
		http.MethodGet,
		"/aggregate/dealer/fixedoperation",
		map[string]uint8{com.DealerResourceName: com.Read, com.FixedOperationResourceName: com.Read},
		aggregateDealerFixedOpH,
	)
	tapi.AddRoute(
		"readDealerGoal",
		http.MethodGet,
		"/goal/{gid}",
		map[string]uint8{com.GoalResourceName: com.Read},
		readDealerGoalH,
	)
	tapi.AddRoute(
		"readDealerGoals",
		http.MethodGet,
		"/goals",
		map[string]uint8{com.GoalResourceName: com.Read},
		readDealerGoalsH,
	)

	tapi.AddRoute(
		"readDealerGroups",
		http.MethodGet,
		"/groups",
		map[string]uint8{com.GroupResourceName: com.Read},
		readDealerGroupsH,
	)

	//log service start info
	log.GenericInfo(apiContext.TContext{}, "Started Tekion tdealer on port:8079", nil)
	tapi.Start("8079", "/tdealer")
}
