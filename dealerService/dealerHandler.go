package dealerService

// This file contains handler functions

import (
	//standard libraries
	"errors"
	"net/http"

	//third party libraries
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//tekion specific libraries
	"bitbucket.org/tekion/erratum"
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/tapi"
)

// swagger:operation GET /dealer dealer readDealer
//
// Returns Dealer identified by the dealer id
//
// By default /dealer returns complete dealer object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /dealer?fields=dealerDoingBusinessAsName,vehicleDamage,dealerAddress
//   required: false
//   type: string
// responses:
//   '200':
//     description: dealer object
//     schema:
//         "$ref": "#/definitions/dealer"
//   '204':
//     description: dealer not found in data base
//   '400':
//     description: error querying data base
func readDealer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	//assuming logged in user has access to view all the dealers
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var dealer dealer

	err := fetchOne(ctx,
		getDealerCollectionName(),
		bson.M{"_id": dealerID},
		fields,
		&dealer,
	)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}
	// No need to check if some thing was found or not. readOne returns "not found".
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", dealer)
}

// swagger:operation GET /fixedoperation/{foid} fixedOperation readFixedOperation
//
// Returns fixed operation identified by fixed operation id passed as part of url
//
// By default /fixedoperation/{foid} returns complete fixed operation object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: foid
//   in: path
//   description: unique identifier of the fixed operation
//   required: true
//   type: string
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /fixedoperation/{foid}?fields=serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity
//   required: false
//   type: string
// responses:
//   '200':
//     description: fixed operation object
//     schema:
//         "$ref": "#/definitions/fixedOperation"
//   '204':
//     description: fixed operation not found in data base
//   '400':
//     description: error querying data base
func readFixedOperation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fixedOperationID := vars["foid"]
	if len(fixedOperationID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer fixed operation id missing in request"))
		return
	}

	var fo FixedOperation
	fields := fetchFieldsFromRequest(r)
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	err := fetchOne(ctx,
		getFixedOperationCollectionName(),
		bson.M{"_id": fixedOperationID},
		fields,
		&fo,
	)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", fo)
}

// swagger:operation GET /fixedoperations fixedOperations readFixedOperations
//
// Returns list of fixed operations identified by dealer id passed in header
//
// By default /fixedoperations returns list of complete fixed operation objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /fixedoperations?field=serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity
//   required: false
//   type: string
// responses:
//   '200':
//     description: list of fixed operations
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/fixedOperation"
//   '204':
//     description: fixed operations not found in data base
//   '400':
//     description: error querying data base
func readFixedOperations(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	var fixedOperations []FixedOperation
	fields := fetchFieldsFromRequest(r)
	err := fetchFixedOperations(ctx,
		bson.M{"dealerID": dealerID},
		fields,
		&fixedOperations,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	if len(fixedOperations) == 0 {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", fixedOperations)
}

// swagger:operation GET /contact/{cid} dealerContact readDealerContact
//
// Returns dealer contact identified by dealer contact id passed as part of url
//
// By default /contact/{cid} returns complete dealer contact object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: cid
//   in: path
//   description: unique identifier of the dealer contact
//   required: true
//   type: string
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /contact/{cid}?fields=user,userDisplayName,userDisplayTitle
//   required: false
//   type: string
// responses:
//   '200':
//     description: dealer contact object
//     schema:
//         "$ref": "#/definitions/dealerContact"
//   '204':
//     description: dealer contact not found in data base
//   '400':
//     description: error querying data base
func readDealerContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID := vars["cid"]
	if len(contactID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer contact id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var contact dealerContact
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	err := fetchOne(ctx,
		getDealerContactCollectionName(),
		bson.M{"_id": contactID},
		fields,
		&contact,
	)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", contact)
}

// swagger:operation GET /contacts dealerContacts readDealerContacts
//
// Returns list of dealer contacts identified by dealer id passed in header
//
// By default /contacts returns list of complete dealer contacts objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /contacts?fields=user,userDisplayName,userDisplayTitle
//   required: false
//   type: string
// responses:
//   '200':
//     description: list of dealer contacts
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/dealerContact"
//   '204':
//     description: dealer contacts not found in data base
//   '400':
//     description: error querying data base
func readDealerContacts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var contacts []dealerContact
	err := fetchDealerContacts(ctx,
		bson.M{"dealerID": dealerID},
		fields,
		&contacts,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	if len(contacts) == 0 {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", contacts)
}

// swagger:operation GET /goal/{gid} dealerGoal readDealerGoal
//
// Returns dealer goal identified by dealer goal id passed as part of url
//
// By default /goal/{gid} returns complete dealer goal object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: gid
//   in: path
//   description: unique identifier of the dealer goal
//   required: true
//   type: string
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /goal/{id}?fields=hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal
//   required: false
//   type: string
// responses:
//   '200':
//     description: dealer goal object
//     schema:
//         "$ref": "#/definitions/dealerGoal"
//   '204':
//     description: dealer goal not found in data base
//   '400':
//     description: error querying data base
func readDealerGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalID := vars["gid"]
	if len(goalID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer goal id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var goal dealerGoal
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	err := fetchOne(ctx,
		getDealerGoalCollectionName(),
		bson.M{"_id": goalID},
		fields,
		&goal,
	)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", goal)
}

// swagger:operation GET /goals dealerGoals readDealerGoals
//
// Returns list of dealer goals identified by dealer id passed in header
//
// By default /goals returns list of complete dealer goals objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /goals?fields=hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal
//   required: false
//   type: string
// responses:
//   '200':
//     description: list of dealer goals
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/dealerGoal"
//   '204':
//     description: dealer goals not found in data base
//   '400':
//     description: error querying data base
func readDealerGoals(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var goals []dealerGoal
	err := fetchDealerGoals(ctx,
		bson.M{"dealerID": dealerID},
		fields,
		&goals,
	)

	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	if len(goals) == 0 {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", goals)
}

// swagger:operation GET /groups dealerGroups readDealerGroups
//
// Returns list of dealer groups identified by dealer id passed in header
//
// By default /groups returns list of complete dealer groups objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
//
// ---
// produces:
// - application/json
// parameters:
// - name: dealerid
//   in: header
//   description: unique identifier of the dealer
//   required: true
//   type: string
// - name: clientid
//   in: header
//   description: client type
//   required: true
//   type: string
// - name: tenantname
//   in: header
//   description: current tenant name
//   required: true
//   type: string
// - name: tekion-api-token
//   in: header
//   description: auth token
//   required: true
//   type: string
// - name: fields
//   in: query
//   description: list of comma separated fields you want in response e.g /groups?fields=dealerGroupName,dealerGroupName,dealers
//   required: false
//   type: string
// responses:
//   '200':
//     description: list of dealer groups
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/dealerGroup"
//   '204':
//     description: dealer groups not found in data base
//   '400':
//     description: error querying data base
func readDealerGroups(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var groups []dealerGroup
	err := fetchDealerGroups(ctx,
		bson.M{"dealers": dealerID},
		fields,
		&groups,
	)

	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	if len(groups) == 0 {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", groups)
}
