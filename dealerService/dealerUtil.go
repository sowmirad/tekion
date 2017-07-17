package dealerService

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/log"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
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

// TODO : should be moved to some common library
// fetchOne reads only one object from mongo and populates it in data parameter passed to the function.
// Returns 1st object if multiple objects are selected by selector parameter.
func fetchOne(ctx apiContext.APIContext, collectionName string, selector bson.M, fields []string, data interface{}) error {
	tenantName := ctx.Tenant
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}

	// Collection
	c := mongo.DB(tenantName).C(collectionName)

	err = c.Find(selector).Select(selectedFields(fields)).One(data)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	return err
}

// fetchDealerGroups reads list of dealer groups from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func fetchDealerGroups(ctx apiContext.APIContext, selector bson.M, fields []string, data *[]DealerGroup) error {
	mongo, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	// Collection
	c := mongo.DB(ctx.Tenant).C(getDealerGroupCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	return err
}

// fetchDealerContacts reads list of dealer contacts from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func fetchDealerContacts(ctx apiContext.APIContext, selector bson.M, fields []string, data *[]DealerContact) error {
	mongo, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	// Collection
	c := mongo.DB(ctx.Tenant).C(getDealerContactCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	return err
}

// fetchFixedOperations reads list of dealer fixed operations from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func fetchFixedOperations(ctx apiContext.APIContext, selector bson.M, fields []string, data *[]FixedOperation) error {
	mongo, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	// Collection
	c := mongo.DB(ctx.Tenant).C(getFixedOperationCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	return err
}

// fetchDealerGoals reads list of dealer goals from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func fetchDealerGoals(ctx apiContext.APIContext, selector bson.M, fields []string, data *[]DealerGoal) error {
	mongo, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	// Collection
	c := mongo.DB(ctx.Tenant).C(getDealerGoalCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	return err
}
