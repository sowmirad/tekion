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
	"bitbucket.org/tekion/tbaas/mongoManager"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tbaas/tapi"
	"encoding/json"
	"fmt"
)

var (
	errDealerID = errors.New("empty user id")
)

const (
	apiCtxKey = "apiContext"
)

// swagger:operation GET /dealer dealer readDealer
//
// Returns Dealer identified by the dealer id
//
// By default /dealer returns complete dealer object.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /dealer?fields=dealerDoingBusinessAsName,vehicleDamage,dealerAddress
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
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	//assuming logged in user has access to view all the dealers
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var dealer dealer

	err := mongoManager.ReadOne(ctx.Tenant, dealerCollectionName, bson.M{"_id": dealerID}, selectedFields(fields), &dealer)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}
	// No need to check if some thing was found or not. readOne returns "not found".
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", dealer)
}

//dealerList is to query list of dealers from Dealermaster
// swagger:operation POST /dealers listDealersReq dealerList
//
// Returns Dealer list
//
// By default /lstDealer returns complete dealer list.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /dealer?fields=dealerDoingBusinessAsName,vehicleDamage,dealerAddress
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
func dealerList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)

	var lstDealer listDealersReq
	err := json.NewDecoder(r.Body).Decode(&lstDealer)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}
	findQuery := bson.M{}
	selectQuery := lstDealer.prepareSelectQuery()
	var dealerLst []dealer

	if err := mMgr.ReadAll(ctx.Tenant, dealerCollectionName, findQuery, selectQuery, &dealerLst); err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}
	if len(dealerLst) == 0 {
		tapi.WriteHTTPResponse(w, http.StatusOK, "No document found, returning empty list", dealerLst)
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "dealer list", dealerLst)
	return

}

// patchDealer is to query list of dealers from Dealermaster
// swagger:operation PATCH /dealer dealer patchDealer
//
// Returns dealer list of columns to update
//
// By default /dealerDtls returns complete dealer details.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /dealer?fields=dealerDoingBusinessAsName,vehicleDamage,dealerAddress
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
func patchDealer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	var dealerDtls dealer
	if err := json.NewDecoder(r.Body).Decode(&dealerDtls); err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorDecodingPayload,
			fmt.Errorf("error encountered while decoding userDetails payload: %v", err))
		return
	}
	if len(dealerDtls.ID) == 0 {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorDecodingPayload, errDealerID)
		return
	}
	var userDtls userDtlsRes
	if err := getUserDtls(ctx, r, &userDtls); err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorDocumentNotFound,
			fmt.Errorf("failed to get user id in db: %v", err))
		return
	}
	findQ := bson.M{"_id": dealerDtls.ID}
	dealerDtls.LastUpdatedByDisplayName = userDtls.Data.DisplayName
	updateQ := dealerDtls.prepareUpdateQuery(ctx, r)
	if err := mMgr.Update(ctx.Tenant, dealerCollectionName, findQ, updateQ); err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorUpdatingMongoDoc,
			fmt.Errorf("error encountered while updating dealer details in db: %v", err))
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "dealer details updated", nil)
}

//createNewDealer is for creating new dealer
func createDealer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)

	var dealerDtls dealer

	err := json.NewDecoder(r.Body).Decode(&dealerDtls)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorDecodingPayload,
			fmt.Errorf("error while decoding creating new dealer payload: %v", err))
		return
	}

	if err = validateNewDealer(ctx, &dealerDtls); err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorWithQueryParams,
			fmt.Errorf("missing fields required to create new dealer: %v", err))
		return
	}

	//Inserting dealers into DB
	if err = mMgr.Create(ctx.Tenant, dealerCollectionName, &dealerDtls); err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorInsertingMongoDoc,
			fmt.Errorf("failed to insert new dealer in db: %v", err))
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "dealer created successfully", nil)
}

// swagger:operation GET /fixedoperation fixedOperation readFixedOperation
//
// Returns list of fixed operations identified by dealer id passed in header
//
// By default /fixedoperation returns list of complete fixed operation objects.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /fixedoperations?field=serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity
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
func readFixedOperation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	var fixedOperation fixedOperation
	fields := fetchFieldsFromRequest(r)
	err := mongoManager.ReadOne(ctx.Tenant, fixedOperationCollectionName,
		bson.M{"dealerID": dealerID}, selectedFields(fields), &fixedOperation)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", fixedOperation)
}

// swagger:operation GET /contact/{cid} dealerContact readDealerContact
//
// Returns dealer contact identified by dealer contact id passed as part of url
//
// By default /contact/{cid} returns complete dealer contact object.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /contact/{cid}?fields=user,userDisplayName,userDisplayTitle
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

	fields := fetchFieldsFromRequest(r)
	var contact dealerContact
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	err := mongoManager.ReadOne(ctx.Tenant, dealerContactCollectionName,
		bson.M{"_id": contactID, "dealerID": ctx.DealerID}, selectedFields(fields), &contact)
	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", contact)
}

// swagger:operation GET /contacts dealerContacts readDealerContacts
//
// Returns list of dealer contacts identified by dealer id passed in header
//
// By default /contacts returns list of complete dealer contacts objects.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /contacts?fields=user,userDisplayName,userDisplayTitle
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
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var contacts []dealerContact
	err := mongoManager.ReadAll(ctx.Tenant, dealerContactCollectionName,
		bson.M{"dealerID": dealerID}, selectedFields(fields), &contacts)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
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
// By default /goal/{gid} returns complete dealer goal object.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /goal/{id}?fields=hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal
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

	fields := fetchFieldsFromRequest(r)
	var goal dealerGoal
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	err := mongoManager.ReadOne(ctx.Tenant, dealerGoalCollectionName,
		bson.M{"_id": goalID, "dealerID": ctx.DealerID}, selectedFields(fields), &goal)

	if err == mgo.ErrNotFound {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", goal)
}

// swagger:operation GET /goals dealerGoals readDealerGoals
//
// Returns list of dealer goals identified by dealer id passed in header
//
// By default /goals returns list of complete dealer goals objects.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /goals?fields=hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal
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
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var goals []dealerGoal
	err := mongoManager.ReadAll(ctx.Tenant, dealerGoalCollectionName,
		bson.M{"dealerID": dealerID}, selectedFields(fields), &goals)

	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
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
// By default /groups returns list of complete dealer groups objects.
// In case you need only certain fields, you can specify an optional query parameter "fields",
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
//   description: e.g /groups?fields=dealerGroupName,dealerGroupName,dealers
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
	ctx := context.Get(r, apiCtxKey).(apiContext.APIContext)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var groups []dealerGroup
	err := mongoManager.ReadAll(ctx.Tenant, dealerGroupCollectionName,
		bson.M{"dealers": dealerID}, selectedFields(fields), &groups)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	if len(groups) == 0 {
		tapi.WriteHTTPResponse(w, http.StatusNoContent, "No document found", nil)
		return
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", groups)
}
