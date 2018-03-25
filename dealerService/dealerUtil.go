package dealerService

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/tekion/tbaas/apiContext"
	l "bitbucket.org/tekion/tenums/login"

	"gopkg.in/mgo.v2/bson"
)

const (
	loginServiceID = "tuser"
	signUpEndPoint = "/tuser/username/"
	appJSON        = "application/json"
)

const (
	customerServiceID         = "tcustomer"
	getUserByUserNameEndPoint = "/tloginservice/getUserByUserName/"
)

var (
	docVersion = float32(1.0)
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
			if v == "dealerID" {
				selectQ["_id"] = 1
			} else {
				selectQ[v] = 1
			}
		}
		return selectQ
	}
	return nil
}

func (lstDealer *listDealersReq) prepareFindQuery() bson.M {
	findQ := make(bson.M)
	if len(lstDealer.IDs) != 0 {
		ids := make([]string, 0, len(lstDealer.IDs))
		for _, id := range lstDealer.IDs {
			ids = append(ids, id)
		}
		findQ["_id"] = bson.M{"$in": ids}
	}
	return findQ
}

// func getUserDtls(ctx *customCtx, r *http.Request, userDtlsRes *userDtlsRes) error {
// 	url := consulhelper.GetServiceNodes(ctx.TContext, loginServiceID) + signUpEndPoint + ctx.UserName
// 	resp, err := hwrap.MakeHTTPRequestWithCustomHeader(http.MethodGet, url, appJSON, r.Header, nil)
// 	if err != nil {
// 		err = fmt.Errorf("call to %s failed, error: %v", url, err)
// 		log.GenericError(ctx.TContext, err, nil)
// 		return err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		respBody, _ := ioutil.ReadAll(resp.Body)
// 		err := fmt.Errorf("call to %s returned error, response body: %s, code: %d", url, string(respBody), resp.StatusCode)
// 		log.GenericError(ctx.TContext, err, nil)
// 		return err
// 	}
// 	//Decode
// 	if err = json.NewDecoder(resp.Body).Decode(&userDtlsRes); err != nil {
// 		err = fmt.Errorf("error encountered while decoding %s reponse, error: %v", url, err)
// 		log.GenericError(ctx.TContext, err, nil)
// 		return err
// 	}

// 	return nil
// }

//prepareUpdateQuery is use to update the DealerMaster
func (d *dealer) prepareUpdateQuery(ctx *customCtx) bson.M {

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

	updateQuery["isActive"] = true
	updateQuery["lastUpdatedByUser"] = d.LastUpdatedByUser
	updateQuery["lastUpdatedByDisplayName"] = d.LastUpdatedByDisplayName
	updateQuery["lastUpdatedDateTime"] = time.Now()
	updateQuery["documentVersion"] = docVersion
	return bson.M{"$set": updateQuery}
}

func (fo *fixedOperation) prepareUpdateQuery(ctx *customCtx) bson.M {
	updateQuery := make(bson.M)
	if len(fo.DealerID) != 0 {
		updateQuery["dealerID"] = fo.DealerID
	}
	if len(fo.EPANumber) != 0 {
		updateQuery["EPANumber"] = fo.EPANumber
	}
	if len(fo.BARNumber) != 0 {
		updateQuery["BARNumber"] = fo.BARNumber
	}
	if len(fo.Disclaimer) != 0 {
		updateQuery["taxPercentage"] = fo.Disclaimer
	}
	updateQuery["isActive"] = true
	updateQuery["lastUpdatedByUser"] = fo.LastUpdatedByUser
	updateQuery["lastUpdatedDateTime"] = time.Now()
	updateQuery["documentVersion"] = docVersion
	return bson.M{"$set": updateQuery}
}

func (d *dealer) populateMetaData(ctx *customCtx) {
	d.IsActive = true
	d.LastUpdatedDateTime = time.Now()
	d.DocumentVersion = docVersion
	d.LastUpdatedByUser = ctx.UserName
	d.LastUpdatedByDisplayName = ctx.UserDisplayName
}

func (fo *fixedOperation) populateMetaData(ctx *customCtx) {
	fo.IsActive = true
	fo.LastUpdatedDateTime = time.Now()
	fo.DocumentVersion = docVersion
	fo.LastUpdatedByUser = ctx.UserName
	fo.LastUpdatedByDisplayName = ctx.UserDisplayName
}

/************************************ Context ************************************/
type customCtx struct {
	apiContext.TContext
	UserID          string
	UserDisplayName string
}

func getCustomCtx(r *http.Request) *customCtx {
	customCtx := new(customCtx)
	customCtx.TContext = apiContext.UpgradeCtx(r.Context())
	return customCtx
}

func getUserCtx(r *http.Request) (*customCtx, error) {
	customCtx := getCustomCtx(r)
	userIDC, _ := r.Cookie(l.UserIDC)
	userDisplayName, _ := r.Cookie(l.UserDisplayNameC)
	if userIDC == nil || userDisplayName == nil {
		usr, err := userByUserName(customCtx)
		if err != nil {
			err = fmt.Errorf("failed to get user details, userName: %s, err: %v ", customCtx.UserName, err)
			return customCtx, err
		}
		customCtx.UserID = usr.ID
		customCtx.UserDisplayName = usr.DisplayName
	} else {
		customCtx.UserID = userIDC.Value
		customCtx.UserDisplayName = userDisplayName.Value
	}

	return customCtx, nil
}
