package dealerService

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

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
	} else {
		return nil
	}
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
// readOne reads only one object from mongo and populates it in data parameter passed to the function.
// Returns 1st object if multiple objects are selected by selector parameter.
func readOne(tenantName, collectionName string, selector bson.M, fields []string, data interface{}) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}

	// Collection
	c := mongo.DB(tenantName).C(collectionName)

	err = c.Find(selector).Select(selectedFields(fields)).One(data)
	if err != nil {
		return err
	}
	return nil
}

// readDealerGroups reads list of dealer groups from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func readDealerGroups(tenantName string, selector bson.M, fields []string, data *[]DealerGroup) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(getDealerGroupCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}

// readDealerContacts reads list of dealer contacts from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func readDealerContacts(tenantName string, selector bson.M, fields []string, data *[]DealerContact) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(getDealerContactCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}

// readFixedOperations reads list of dealer fixed operations from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func readFixedOperations(tenantName string, selector bson.M, fields []string, data *[]FixedOperation) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(getDealerFixedOperationCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}

// readDealerGoals reads list of dealer goals from mongo based on the selector passed.
// Populates it in data parameter passed to the function.
func readDealerGoals(tenantName string, selector bson.M, fields []string, data *[]DealerGoal) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(getDealerGoalCollectionName())

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}
