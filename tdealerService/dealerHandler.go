package dealerService

// This file contains handler functions

import (
	//standard libraries
	"net/http"

	//third party libraries
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	//tekion specific libraries
	"bitbucket.org/tekion/erratum"
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/tapi"
	"github.com/pkg/errors"
)

// ReadDealer - this handler function returns the Dealer Object. Reads the Dealer object from database identified by the Dealer.ID passed in header.
// By default ReadDealer returns complete Dealer object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /dealer?fields=dealerDoingBusinessAsName,vehicleDamage,dealerAddress
// @param - dealerid - (required) - unique identifier of the Dealer.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete Dealer object.
// @returns - json of Dealer object.
func ReadDealer(w http.ResponseWriter, r *http.Request) {
	//assuming logged in user has access to view all the dealers
	dealerID := r.Header.Get("dealerid") // should be corrected to Dealer-ID
	if len(dealerID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var dealer Dealer
	err := readOne(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerCollectionName(),
		bson.M{"_id": dealerID},
		fields,
		&dealer,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}
	// No need to check if some thing was found or not. readOne returns "not found".
	tapi.WriteHTTPResponse(w, http.StatusOK, "Document found", dealer)
}

// ReadFixedOperation - this handler function returns the FixedOperation object. Reads the FixedOperation object from database identified by FixedOperation.ID passed as part of url.
// By default ReadFixedOperation returns whole FixedOperation object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /fixedoperation/{id}?fields=serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity
// @param - {id} - (required) - unique identifier of the FixedOperation.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete FixedOperation object.
// @returns - json of  FixedOperation object.
func ReadFixedOperation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fixedOperationID := vars["id"]
	//vars = context.Get(r, 0).(map[string]string)
	//fixedOperationID = vars["id"]
	if len(fixedOperationID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("fixed operation id missing in request"))
		return
	}

	var fixedOperation FixedOperation
	fields := fetchFieldsFromRequest(r)
	err := readOne(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerFixedOperationCollectionName(),
		bson.M{"_id": fixedOperationID},
		fields,
		&fixedOperation,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if fixedOperation.ID == "" {
		msg = "Document not found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, fixedOperation)
}

// ReadFixedOperations - this handler function returns json array of FixedOperation objects. Reads the FixedOperation object from database identified by Dealer.ID passed in header.
// By default ReadFixedOperations returns complete FixedOperation objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /fixedoperations?field=serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity
// @param - dealerid - (required) - unique identifier of the Dealer.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete FixedOperation object.
// @returns - json array of FixedOperation objects.
func ReadFixedOperations(w http.ResponseWriter, r *http.Request) {
	dealerID := r.Header.Get("dealerid") // should be corrected to Dealer-ID
	if len(dealerID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("fixed operation id and dealer id missing in request"))
		return
	}

	var fixedOperations []FixedOperation
	fields := fetchFieldsFromRequest(r)
	err := readAllFixedOperations(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerFixedOperationCollectionName(),
		bson.M{"dealerID": dealerID},
		fields,
		&fixedOperations,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if len(fixedOperations) == 0 {
		msg = "No document found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, fixedOperations)
}

// ReadContact - this handler function returns the DealerContact Object. Reads the DealerContact object from database identified by the DealerContact.ID passed as part of url.
// By default ReadContact returns whole DealerContact object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /contact/{id}?fields=user,userDisplayName,userDisplayTitle
// @param - {id} - (required) - unique identifier of the DealerContact.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete DealerContact object.
// @returns - json of DealerContact object.
func ReadContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID := vars["id"]
	if len(contactID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("contact id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var contact DealerContact
	err := readOne(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerContactCollectionName(),
		bson.M{"_id": contactID},
		fields,
		&contact,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if contact.ID == "" {
		msg = "Document not found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, contact)
}

// ReadContacts - this handler function returns json array of DealerContact Objects. Reads the DealerContact object from database identified by the Dealer.ID passed as part of header.
// By default ReadContacts returns complete DealerContact objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /contacts?fields=user,userDisplayName,userDisplayTitle
// @param - dealerid - (required) - unique identifier of the Dealer.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete DealerContact object.
// @returns - json array of DealerContact objects.
func ReadContacts(w http.ResponseWriter, r *http.Request) {
	//assuming logged in user has access to view all the dealers
	dealerID := r.Header.Get("dealerid") // should be corrected to Dealer-ID
	if len(dealerID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var contacts []DealerContact
	err := readAllContacts(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerContactCollectionName(),
		bson.M{"dealerID": dealerID},
		fields,
		&contacts,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if len(contacts) == 0 {
		msg = "No document found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, contacts)
}

// ReadGoal - this handler function returns the DealerGoal Object. Reads the DealerGoal object from database identified by the DealerGoal.ID passed as part of url.
// By default ReadGoal returns complete DealerGoal object. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /goal/{id}?fields=hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal
// @param - {id} - (required) - unique identifier of the DealerGoal.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete DealerGoal object.
// @returns - json of DealerGoal object.
func ReadGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalID := vars["id"]
	if len(goalID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer goal id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var dealerGoal DealerGoal

	err := readOne(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerGoalCollectionName(),
		bson.M{"_id": goalID},
		fields,
		&dealerGoal,
	)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if dealerGoal.ID == "" {
		msg = "Document not found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, dealerGoal)
}

// ReadGoals - this handler function returns json array of DealerGoal Objects. Reads the DealerGoal object from database identified by the dealer id passed in header.
// By default ReadGoals returns complete DealerGoal objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /goals?fields=hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal
// @param - dealerid - (required) - unique identifier of the Dealer.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete DealerGoal objects.
// @returns - json array of DealerGoal objects.
func ReadGoals(w http.ResponseWriter, r *http.Request) {
	dealerID := r.Header.Get("dealerid") // should be corrected to Dealer-ID
	if len(dealerID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var goals []DealerGoal
	err := readAllGoals(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerGoalCollectionName(),
		bson.M{"dealerID": dealerID},
		fields,
		&goals,
	)

	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if len(goals) == 0 {
		msg = "No document found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, goals)
}

// ReadGroups - this handler function returns json array of DealerGroup objects. Reads the list of DealerGroup object from database identified by the Dealer.ID passed in header.
// By default ReadGroups returns complete DealerGroup objects. In case you need only certain fields, you can specify an optional query parameter "fields",
// passing a list of comma separated fields you want in response.
// E.g /groups?fields=dealerGroupName,dealerGroupName,dealers
// @param - dealerid - (required) - unique identifier of the Dealer.
// @param - fields - (optional) - list of comma separated fields you want in response instead of complete DealerGroup objects.
// @returns - json array of DealerGroup objects.
func ReadGroups(w http.ResponseWriter, r *http.Request) {
	dealerID := r.Header.Get("dealerid") // should be corrected to Dealer-ID
	if len(dealerID) == 0 {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, errors.New("dealer id missing in request"))
		return
	}

	fields := fetchFieldsFromRequest(r)
	var groups []DealerGroup
	err := readAllGroups(context.Get(r, "apiContext").(apiContext.APIContext).Tenant,
		getDealerGroupCollectionName(),
		bson.M{"dealers": dealerID},
		fields,
		&groups,
	)

	if err != nil {
		tapi.WriteHTTPErrorResponse(w, getModuleID(), erratum.ErrorQueryingDB, err)
		return
	}

	msg := "Document found"
	if len(groups) == 0 {
		msg = "No document found"
	}
	tapi.WriteHTTPResponse(w, http.StatusOK, msg, groups)
}
