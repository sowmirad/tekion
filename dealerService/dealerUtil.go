package dealerService

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/consulhelper"
	"bitbucket.org/tekion/tbaas/hwrap"
	"bitbucket.org/tekion/tbaas/log"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
<<<<<<< HEAD
	"gopkg.in/mgo.v2/bson"
	"bitbucket.org/tekion/tbaas/apiContext"
=======
>>>>>>> 8dc5389d1d12d4a9243a3a8e48156f6b12ae0980
	"time"
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
func (lstdealer *listDealerReq) prepareSelectQuery() bson.M {
	if len(lstdealer.SelectedFields) != 0 {
		selectQ := make(bson.M)
		for _, v := range lstdealer.SelectedFields {
			selectQ[v] = 1
		}
		return selectQ
	}
	return nil
}

func getUserdls(ctx apiContext.APIContext, r *http.Request, userdtlsRes *userdtlsRes) error {
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
	if err = json.NewDecoder(resp.Body).Decode(&userdtlsRes); err != nil {
		err = fmt.Errorf("error encountered while decoding %s reponse, error: %v", url, err)
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}

	return nil
}

//prepareUpdateQuery is use to update the Dealermaster
func (dealerdtls *dealer) prepareUpdateQuery(ctx apiContext.APIContext, r *http.Request) bson.M {

	updateQuery := make(bson.M)
	if len(dealerdtls.Name) != 0 {
		updateQuery["dealerName"] = dealerdtls.Name
	}
	if len(dealerdtls.DoingBusinessAsName) != 0 {
		updateQuery["dealerDisplayName"] = dealerdtls.DoingBusinessAsName
	}
	if len(dealerdtls.StateIssuedNumber) != 0 {
		updateQuery["stateIssuedNumber"] = dealerdtls.StateIssuedNumber
	}
	if len(dealerdtls.ManufacturerIssuedNumber) != 0 {
		updateQuery["manufacturerIssuedNumber"] = dealerdtls.ManufacturerIssuedNumber
	}
	if len(dealerdtls.Website) != 0 {
		updateQuery["website"] = dealerdtls.Website
	}
	if len(dealerdtls.TimeZone) != 0 {
		updateQuery["timeZone"] = dealerdtls.TimeZone
	}
	if len(dealerdtls.Currency) != 0 {
		updateQuery["currency"] = dealerdtls.Currency
	}
	if len(dealerdtls.TenantID) != 0 {
		updateQuery["tenantID"] = dealerdtls.TenantID
	}
	if len(dealerdtls.Phone) != 0 {
		updateQuery["phone"] = dealerdtls.Phone
	}
	if len(dealerdtls.DealershipCode) != 0 {
		updateQuery["dealershipCode"] = dealerdtls.DealershipCode
	}
	if len(dealerdtls.VideoURL) != 0 {
		updateQuery["videoURL"] = dealerdtls.VideoURL
	}
	if len(dealerdtls.ApplicationCode) != 0 {
		updateQuery["applicationCode"] = dealerdtls.ApplicationCode
	}
	updateQuery["lastUpdatedByUser"] = ctx.UserName
	updateQuery["lastUpdatedByDisplayName"] = dealerdtls.LastUpdatedByDisplayName
	updateQuery["lastUpdatedDateTime"] = time.Now().UTC()
	updateQuery["documentVersion"] = dealerdtls.DocumentVersion
	return bson.M{"$set": updateQuery}
}
