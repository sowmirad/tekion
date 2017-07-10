package dealerService

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
)

// fetchFieldsFromRequest -
func fetchFieldsFromRequest(r *http.Request) []string {
	q := r.URL.Query()
	requestedFields := q.Get("fields")
	if len(requestedFields) != 0 {
		return strings.Split(requestedFields, ",")
	} else {
		return nil
	}
}

// selectedFields -
func selectedFields(fields []string) bson.M {
	selected := make(bson.M, len(fields))
	for _, s := range fields {
		selected[s] = 1
	}
	return selected
}

// ReadOne - reads only one object from mongo. Returns 1st object if multiple objects are selected by selector parameter
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

// readAllGroups -
func readAllGroups(tenantName, collectionName string, selector bson.M, fields []string, data *[]DealerGroup) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(collectionName)

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}

// readAllContacts -
func readAllContacts(tenantName, collectionName string, selector bson.M, fields []string, data *[]DealerContact) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(collectionName)

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}

// readAllFixedOperations -
func readAllFixedOperations(tenantName, collectionName string, selector bson.M, fields []string, data *[]FixedOperation) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(collectionName)

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}

// readAllGoals -
func readAllGoals(tenantName, collectionName string, selector bson.M, fields []string, data *[]DealerGoal) error {
	mongo, err := mMgr.GetS(tenantName)
	if err != nil {
		return err
	}
	// Collection
	c := mongo.DB(tenantName).C(collectionName)

	err = c.Find(selector).Select(selectedFields(fields)).All(data)
	if err != nil {
		return err
	}
	return nil
}
