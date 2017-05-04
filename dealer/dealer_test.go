package dealer

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tbaas/log"
	"gopkg.in/mgo.v2/bson"
)

var (

	testDealerID = "testDealerId" //this id should not exist in Database
	testDealerName = "test Dealer Name"
	testTenantID = "test tenant Id"
	testTenantDisplayName = "Buck"

	//test tenantName and dealerID used in context
	correctTenantName   = "Buck"
	correctDealerID     = "3"
	incorrectTenantName = "ABCCD"
	incorrectDealerID   = "99"

	ctxD3        = apiContext.APIContext{Tenant: correctTenantName, DealerID: correctDealerID}
	ctxIncorrect = apiContext.APIContext{Tenant: incorrectTenantName, DealerID: incorrectDealerID}

	testDealerObject = Dealer{
		ID : testDealerID,
		DealerName: testDealerName,
		TenantID: testTenantID,
		TenantDisplayName: testTenantDisplayName,
		SkillSet: []string{"Engine"},
	}


)
//function to insert test data in Database
func setupTestData() (error){
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
func clearTestData() (error){
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

func TestGetDealerByIDMethod(t *testing.T) {
	Convey("Testing GetDealerByID function ", t, func(){
		err := setupTestData()
		Convey("setup test data should not give error ", func(){
			So(err, ShouldBeNil)
		})
		Convey("Testing for correct context ", func(){
			dealerObject, err := GetDealerByID(ctxD3, testDealerID)
			Convey("error returned by GetDealerByID function should be nil", func(){
				So(err, ShouldBeNil)
			})

			Convey("dealer object returned by GetDealerByID should be same as dealer object", func(){
				So(dealerObject, ShouldHaveSameTypeAs, Dealer{})
			})
			Convey("Returned dealer object's dealer Name should be same as test dealer Name", func(){
				So(dealerObject.DealerName, ShouldEqual, testDealerName)
			})
			Convey("Returned dealer object's TenantID should be same as TenantID", func(){
				So(dealerObject.TenantID, ShouldEqual, testTenantID)
			})
			Convey("Returned dealer object's TenantDisplayName should be same as testTenantDisplayName", func(){
				So(dealerObject.TenantDisplayName, ShouldEqual, testTenantDisplayName)
			})
		})
		Convey("Testing for incorrect context ", func(){
			_, err := GetDealerByID(ctxIncorrect, testDealerID)
			Convey("Error returned by GetDealerByID should return an error", func(){
				So(err, ShouldNotBeNil)
			})
		})
		err = clearTestData()
		Convey("Clear test data should not give error ", func(){
			So(err, ShouldBeNil)
		})
	})

}

func TestInsertMethod(t *testing.T) {
	Convey("Testing Insert function ", t ,func(){

		Convey("Inserting dealer object for correct context ", func(){
			err := testDealerObject.Insert(ctxD3)
			Convey("Insert test dealer should not give error", func(){
				So(err,ShouldBeNil)
			})
		})
		Convey("Inserting dealer object for incorrect context ", func(){
			err := testDealerObject.Insert(ctxIncorrect)
			Convey("Insert function should give error ", func(){
				So(err, ShouldNotBeEmpty)
			})
		})
		clearTestData()
	})
}

func TestGetDamageTypes(t *testing.T) {
	Convey("Testing function GetDamageTypes ", t, func(){
		err := setupTestData()
		Convey("setup test data should not give error ", func(){
			So(err, ShouldBeNil)
		})
		Convey("Testing for valid context", func(){
			vehicleDamageResponse, err := GetDamageTypes(ctxD3, testDealerID)
			Convey("Error should not be returned by GetDamageTypes function ", func(){
				So(err, ShouldBeNil)
			})
			Convey("vehicleDamageResponse object returned by GetDamageTypes should be same as VehicleDamageMaster", func(){
				So(vehicleDamageResponse, ShouldHaveSameTypeAs, []SelectDamageResponse{})
			})
		})
		Convey("Testing for invalid context ", func(){
			_, err := GetDamageTypes(ctxIncorrect, testDealerID)
			Convey("Error should be returned by GetDamageTypes function ", func(){
				So(err, ShouldNotBeNil)
			})
		})

		err = clearTestData()
		Convey("Clear test data should not give error ", func(){
			So(err, ShouldBeNil)
		})
	})
}