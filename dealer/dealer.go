package dealer

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/log"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tvehicle/vehicle"
	"gopkg.in/mgo.v2/bson"
)

//dealerCollectionName : collection name of DealerMaster
var (
	dealerCollectionName = "DealerMaster"
)

//Dealer : Model to store master data about each dealer
type Dealer struct {
	ID                  string   `bson:"_id" json:"dealerId"`
	DealerName          string   `bson:"dealerName" json:"dealerName"`
	DealerDisplayName   string   `bson:"dealerDisplayName" json:"dealerDisplayName"`
	TenantID            string   `bson:"tenantId" json:"tenantId"`
	TenantDisplayName   string   `bson:"tenantDisplayName" json:"tenantDisplayName"`
	EPANumber           string   `bson:"epaNumber" json:"epaNumber"` // 'ADB 1343857'
	BARNumber           string   `bson:"barNumber" json:"barNumber"` // 'CAL00234957'
	Email               string   `bson:"email" json:"email"`
	PhoneNumber         string   `bson:"phoneNumber" json:"phoneNumber"`
	StreetAddress1      string   `bson:"streetAddress1" json:"streetAddress1"`
	City                string   `bson:"city" json:"city"`
	State               string   `bson:"state" json:"state"`
	Country             string   `bson:"country" json:"country"`
	PostalCode          string   `bson:"postalCode" json:"postalCode"`
	Website             string   `bson:"website" json:"website"`
	VehicleDamageID     []string `bson:"vehicleDamage" json:"vehicleDamage"` //Note: Stores Id's of all vehicle Damages serviced by dealer. Multiple dealers can support same vehicle damage, so for improved fetch of vehicle damage, DealerMaster holds this array.
	ServicesID          []string `bson:"services" json:"services"`           //Note: Stores Id's of all Services provided by dealer. Multiple dealers can provide same service, so for improved fetch of services provided by a dealer, DealerMaster holds this array.
	TimeZone            string   `bson:"timeZone" json:"timeZone"`           //Used for time conversions.
	Currency            string   `bson:"currency" json:"currency"`
	Logo                string   `bson:"logo" json:"logo"`
	WorkingDaysAndHours string   `bson:"workingDaysAndHours" json:"workingDaysAndHours"`
	TaxPercentage       float32   `bson:"taxPercentage" json:"taxPercentage"`
	Disclaimer          string   `bson:"disclaimer" json:"disclaimer"`
	SkillSet            []string `bson:"skillSet" json:"skillSet"`
	//todo : add need []VehicleComponentInspection here ?
}

// Insert : function to insert dealers to DB
func (dealer Dealer) Insert(ctx apiContext.APIContext) error {
	s, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		return err
	}
	defer s.Close()
	return s.DB(ctx.Tenant).C(dealerCollectionName).Insert(dealer)
}

//SelectDamageResponse : structure for SelectDamageResponse
type SelectDamageResponse struct {
	VehicleDamage []vehicle.VehicleDamageMaster `json:"vehicleDamage"`
}

//GetDealerByID : Function to get dealer by dealer ID
func GetDealerByID(ctx apiContext.APIContext, dealerID string) (Dealer, error) {

	dealer := Dealer{}

	session, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.Error("Session error ", err.Error())
		return dealer, err
	}
	defer session.Close()

	//Fetch timezone for dealer
	err = session.DB(ctx.Tenant).C(dealerCollectionName).Find(bson.M{"_id": dealerID}).One(&dealer)
	if err != nil {
		log.Error("not found dealer", err.Error())
		return dealer, err
	}
	//Return dealer detail
	return dealer, nil
}
