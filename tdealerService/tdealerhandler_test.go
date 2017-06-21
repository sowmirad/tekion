package tdealerService

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tdealer/dealer"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	dealerCollectionName = "DealerMaster"

	//test tenantName and dealerID used in context
	correctTenantName   = "Buck"
	correctDealerID     = "3"
	incorrectTenantName = "ABCCD"
	incorrectDealerID   = "99"

	ctxD3        = apiContext.APIContext{Tenant: correctTenantName, DealerID: correctDealerID}
	ctxIncorrect = apiContext.APIContext{Tenant: incorrectTenantName, DealerID: incorrectDealerID}

	testDealerID          = "testDealerId" //this id should not exist in Database
	testDealerName        = "test Dealer Name"
	testTenantID          = "test tenant Id"
	testTenantDisplayName = "Buck"

	testDealerObject = dealer.Dealer{
		ID:                testDealerID,
		DealerName:        testDealerName,
		TenantID:          testTenantID,
		TenantDisplayName: testTenantDisplayName,
		SkillSet:          []string{"Engine"},
	}
)

//function to insert test data in Database
func setupTestData() error {
	session, err := mongoManager.GetS(ctxD3.Tenant)
	if err != nil {
		log.Error("mongo session error", err.Error())
		return err
	}
	defer session.Close()

	//inserting test dealer in database
	err = testDealerObject.Insert(ctxD3)
	if err != nil {
		log.Error("Unable to insert dealer into Database ", err.Error())
		return err
	}
	return err
}

//function to delete test data from Database
func clearTestData() error {
	session, err := mongoManager.GetS(ctxD3.Tenant)
	if err != nil {
		log.Error("mongo session error", err.Error())
		return err
	}
	defer session.Close()

	//deleting test dealer from Database
	err = session.DB(ctxD3.Tenant).C(dealerCollectionName).Remove(bson.M{"_id": testDealerID})
	if err != nil {
		log.Error("unable to delete dealer from DB", err.Error())
		return err
	}
	return err
}

func TestGetDealerByID(t *testing.T) {
	setupTestData()
	// Testing handler function
	req, err := http.NewRequest("GET", "http://localhost:8079/tdealer/getDealerByID", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := apiContext.APIContext{Tenant: "Buck", DealerID: testDealerID}

	context.Set(req, "apiContext", ctx)
	req.Header.Add("tekion-api-token", "TestToken")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	GetDealerByID(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Testing handler function for incorrect context
	req1, err := http.NewRequest("GET", "http://localhost:8079/tdealer/getDealerByID", nil)
	if err != nil {
		t.Fatal(err)
	}

	//setting incorrect context
	context.Set(req1, "apiContext", ctxIncorrect)
	req.Header.Add("tekion-api-token", "TestToken")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr1 := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	GetDealerByID(rr1, req1)

	// Check the status code is what we expect.
	if status := rr1.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
	clearTestData()
}
