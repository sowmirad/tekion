package dealerService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/consulhelper"
	"bitbucket.org/tekion/tbaas/hwrap"
	"bitbucket.org/tekion/tbaas/log"

	"gopkg.in/mgo.v2/bson"
)

const (
	loginServiceID = "tuser"
	signupEndPoint = "/tuser/username/"
	appJSON        = "application/json"
)

// TODO : should be moved to some common library
// fetchFieldsFromRequest reads the query string and fetches "fields" parameter.
// It return slice of strings or nil if "fields" parameter was not found in query string.
func fetchFieldsFromRequest(r *http.Request) []string {
	q := r.URL.Query()
	requestedFields := q.Get("fields")
	if len(requestedFields) != 0 {
		return strings.Split(requestedFields, ",")
	}
	return nil
}

// TODO : should be moved to some common library
// selectedFields forms a map, key = selected field and value = 1
func selectedFields(fields []string) bson.M {
	selected := make(bson.M, len(fields))
	for _, s := range fields {
		selected[s] = 1
	}
	return selected
}

// prepareSelectQuery is to select query form listdealersReq.SelectedFields
func (lstDealer *listDealersReq) prepareSelectQuery() bson.M {
	if len(lstDealer.SelectedFields) != 0 {
		selectQ := make(bson.M)
		for _, v := range lstDealer.SelectedFields {
			selectQ[v] = 1
		}
		return selectQ
	}
	return nil
}
func getUserDtls(ctx apiContext.APIContext, r *http.Request, userDtlsRes *userDtlsRes) error {
	url := consulhelper.GetServiceNodes(loginServiceID) + signupEndPoint + ctx.UserName
	resp, err := hwrap.MakeHTTPRequestWithCustomHeader(http.MethodGet, url, appJSON, r.Header, nil)
	if err != nil {
		err = fmt.Errorf("call to %s failed, error: %v", url, err)
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("call to %s returned error, response body: %s, code: %d", url, string(respBody), resp.StatusCode)
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	//Decode
	if err = json.NewDecoder(resp.Body).Decode(&userDtlsRes); err != nil {
		err = fmt.Errorf("error encountered while decoding %s reponse, error: %v", url, err)
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}

	return nil
}

//prepareUpdateQuery is use to update the DealerMaster
func (d *dealer) prepareUpdateQuery(ctx apiContext.APIContext, r *http.Request) bson.M {

	updateQuery := make(bson.M)
	if len(d.Name) != 0 {
		updateQuery["dealerName"] = d.Name
	}
	if len(d.MakeCode) != 0 {
		updateQuery["makeCode"] = d.MakeCode
	}
	if len(d.DoingBusinessAsName) != 0 {
		updateQuery["dealerDisplayName"] = d.DoingBusinessAsName
	}
	if len(d.StateIssuedNumber) != 0 {
		updateQuery["stateIssuedNumber"] = d.StateIssuedNumber
	}
	if len(d.ManufacturerIssuedNumber) != 0 {
		updateQuery["manufacturerIssuedNumber"] = d.ManufacturerIssuedNumber
	}
	if len(d.Website) != 0 {
		updateQuery["website"] = d.Website
	}
	if len(d.TimeZone) != 0 {
		updateQuery["timeZone"] = d.TimeZone
	}
	if len(d.Currency) != 0 {
		updateQuery["currency"] = d.Currency
	}
	if len(d.TenantID) != 0 {
		updateQuery["tenantID"] = d.TenantID
	}
	if len(d.Phone) != 0 {
		updateQuery["phone"] = d.Phone
	}
	if len(d.VideoURL) != 0 {
		updateQuery["videoURL"] = d.VideoURL
	}

	updateQuery["lastUpdatedByUser"] = d.LastUpdatedByUser
	updateQuery["lastUpdatedByDisplayName"] = d.LastUpdatedByDisplayName
	updateQuery["lastUpdatedDateTime"] = d.LastUpdatedDateTime
	updateQuery["documentVersion"] = d.DocumentVersion
	return bson.M{"$set": updateQuery}
}

func (d *fixedOperation) prepareUpdateQuery(ctx apiContext.APIContext, r *http.Request) bson.M {
	updateQuery := make(bson.M)
	if len(d.DealerID) != 0 {
		updateQuery["dealerID"] = d.DealerID
	}
	if len(d.EPANumber) != 0 {
		updateQuery["EPANumber"] = d.EPANumber
	}
	if len(d.BARNumber) != 0 {
		updateQuery["BARNumber"] = d.BARNumber
	}
	if len(d.Disclaimer) != 0 {
		updateQuery["taxPercentage"] = d.Disclaimer
	}
	updateQuery["lastUpdatedByUser"] = d.LastUpdatedByUser
	updateQuery["lastUpdatedDateTime"] = d.LastUpdatedDateTime
	updateQuery["documentVersion"] = d.DocumentVersion
	return bson.M{"$set": updateQuery}
}
