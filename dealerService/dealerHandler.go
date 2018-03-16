package dealerService

// This file contains handler functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/tekion/erratum"
	"bitbucket.org/tekion/tbaas/mongoManager"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tbaas/tapi"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	apiCtxKey = "apiContext"
)

var (
	errDealerName       = errors.New("dealer name is empty")
	errDealerID         = errors.New("empty dealer id")
	errFixedOperationID = errors.New("empty fixed Operation id")
)

func init() {
	time.Local = time.UTC
}

const (
	docFound = iota
	dealerDocNotFound
	fixedOpDocNotFound
)

// swagger:operation GET /dealer dealer readDealerH
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
func readDealerH(w http.ResponseWriter, r *http.Request) {
	ctx := getCustomCtx(r)

	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	fields := fetchFieldsFromRequest(r)
	var dealer dealer

	err := mongoManager.ReadOne(ctx.Tenant, dealerCollectionName, bson.M{"_id": dealerID}, selectedFields(fields), &dealer)
	if err == mgo.ErrNotFound {
		tapi.HTTPResponse(ctx.TContext, w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}
	// No need to check if some thing was found or not. readOne returns "not found".
	tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "Document found", dealer)
}

// swagger:operation POST /dealers dealer dealersList
//
// Returns list of dealers list
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
//   description: e.g /dealers?fields=dealerDoingBusinessAsName,vehicleDamage,dealerAddress
//   required: false
//   type: string
// - name: listDealersReq
//   in: body
//   description: listDealersReq object
//   required: true
//   schema:
//      "$ref": "#/definitions/listDealersReq"
// responses:
//   '200':
//     description: dealer object
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/dealer"
//   '204':
//     description: dealer not found in data base
//   '400':
//     description: error querying data base
func dealersListH(w http.ResponseWriter, r *http.Request) {
	ctx := getCustomCtx(r)

	var lstDealer listDealersReq
	err := json.NewDecoder(r.Body).Decode(&lstDealer)
	if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}
	findQuery := lstDealer.prepareFindQuery()
	selectQuery := lstDealer.prepareSelectQuery()
	var dealerLst []dealer

	if err := mMgr.ReadAll(ctx.Tenant, dealerCollectionName, findQuery, selectQuery, &dealerLst); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}
	if len(dealerLst) == 0 {
		tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "No document found, returning empty list", dealerLst)
		return
	}
	tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "dealer list", dealerLst)
}

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
func patchDealerH(w http.ResponseWriter, r *http.Request) {
	ctx, err := getUserCtx(r)
	if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}

	d := new(dealer)
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload,
			fmt.Errorf("error encountered while decoding userDetails payload: %v", err))
		return
	}

	if len(d.ID) == 0 {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload, errDealerID)
		return
	}

	d.populateMetaData(ctx)

	findQ := bson.M{"_id": d.ID}
	updateQ := d.prepareUpdateQuery(ctx)
	if err := mMgr.Update(ctx.Tenant, dealerCollectionName, findQ, updateQ); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorUpdatingMongoDoc,
			fmt.Errorf("error encountered while updating dealer details in db: %v", err))
		return
	}

	tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "dealer details updated", nil)
}

// saveDealer dealer details
func saveDealerH(w http.ResponseWriter, r *http.Request) {
	ctx, err := getUserCtx(r)
	if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}

	d := new(dealer)
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload,
			fmt.Errorf("error encountered while decoding save dealer payload: %v", err))
		return
	}

	d.populateMetaData(ctx)

	if len(d.ID) == 0 {
		// create new dealer
		// generating customerID from GetNextSequence function
		if len(strings.TrimSpace(d.Name)) != 0 {
			findQ := bson.M{"dealerName": d.Name}
			count, err := mMgr.Count(ctx.TContext, dealerCollectionName, findQ)
			if err != nil {
				tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload,
					fmt.Errorf("failed to generate dealer id for new dealer, error: %v", err))
				return
			}
			if count != 0 {
				tapi.WriteCustomHTTPErrorResponse(w, serviceID, "400", http.StatusBadRequest,
					"dealer name already exists", 1,
					fmt.Errorf("dealer name already exists"))
				return
			}
		}

		id, err := mMgr.GetNextSequence(ctx.Tenant, dealerCollectionName)
		if err != nil {
			tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload,
				fmt.Errorf("failed to generate dealer id for new dealer, error: %v", err))
			return
		}
		d.ID = id
		if err := mMgr.Create(ctx.Tenant, dealerCollectionName, d); err != nil {
			tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorUpdatingMongoDoc,
				fmt.Errorf("error encountered while creating dealer details in db: %v", err))
			return

		}
		tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "dealer created", d)
		return
	}
	// update existing dealer
	findQ := bson.M{"_id": d.ID}
	if err := mMgr.Update(ctx.Tenant, dealerCollectionName, findQ, d); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorUpdatingMongoDoc,
			fmt.Errorf("error encountered while updating dealer details in db: %v", err))
		return
	}

	tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "dealer details updated", d)
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
//   description: e.g /fixedoperation?field=serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity
//   required: false
//   type: string
// responses:
//   '200':
//     description: list of fixed operations
//     schema:
//         "$ref": "#/definitions/fixedOperation"
//   '204':
//     description: fixed operations not found in data base
//   '400':
//     description: error querying data base
func readFixedOperationH(w http.ResponseWriter, r *http.Request) {
	ctx := getCustomCtx(r)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	var fixedOperation fixedOperation
	fields := fetchFieldsFromRequest(r)
	err := mongoManager.ReadOne(ctx.Tenant, fixedOperationCollectionName,
		bson.M{"dealerID": dealerID}, selectedFields(fields), &fixedOperation)
	if err == mgo.ErrNotFound {
		tapi.HTTPResponse(ctx.TContext, w, http.StatusNoContent, "No document found", nil)
		return
	} else if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "Document found", fixedOperation)
}

//patchFixedOperation is use to update patchfixed operation
func patchFixedOperationH(w http.ResponseWriter, r *http.Request) {
	ctx, err := getUserCtx(r)
	if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}

	fo := new(fixedOperation)
	if err := json.NewDecoder(r.Body).Decode(fo); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload,
			fmt.Errorf("error encountered while decoding fixed operation payload: %v", err))
		return
	}
	if len(fo.ID) == 0 {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorDecodingPayload, errFixedOperationID)
		return
	}

	fo.populateMetaData(ctx)

	findQ := bson.M{"_id": fo.ID}
	updateQ := fo.prepareUpdateQuery(ctx)
	if err := mMgr.Update(ctx.Tenant, fixedOperationCollectionName, findQ, updateQ); err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorUpdatingMongoDoc,
			fmt.Errorf("error encountered while updating fixed operation details in db: %v", err))
		return
	}

	tapi.HTTPResponse(ctx.TContext, w, http.StatusOK, "fixed operations details updated successfully", nil)
}

func aggregateDealerFixedOpH(w http.ResponseWriter, r *http.Request) {
	ctx := getCustomCtx(r)
	dealerID := ctx.DealerID // should be corrected to Dealer-ID

	var dealer *dealer
	err := mongoManager.ReadOne(ctx.Tenant, dealerCollectionName, bson.M{"_id": dealerID}, nil, &dealer)
	if err == mgo.ErrNotFound {
		tapi.CustomHTTPResponse(ctx.TContext, w, http.StatusOK, "dealer doc not found", dealerDocNotFound, nil)
		return
	} else if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	var fixedOp *fixedOperation
	err = mongoManager.ReadOne(ctx.Tenant, fixedOperationCollectionName,
		bson.M{"dealerID": dealerID}, nil, &fixedOp)
	if err == mgo.ErrNotFound {
		tapi.CustomHTTPResponse(ctx.TContext, w, http.StatusNoContent, "fixed operation doc not found",
			fixedOpDocNotFound, nil)
		return
	} else if err != nil {
		tapi.HTTPErrorResponse(ctx.TContext, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	var dealerAndFixedOp readDealerAndFixedOpRes
	dealerAndFixedOp.Dealer = dealer
	dealerAndFixedOp.FixedOperation = fixedOp

	tapi.CustomHTTPResponse(ctx.TContext, w, http.StatusOK, "document found", docFound, dealerAndFixedOp)
}
