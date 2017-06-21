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
	ID                         string                       `bson:"_id" json:"dealerID"`
	DealerName                 string                       `bson:"dealerName" json:"dealerName"`
	DealerDisplayName          string                       `bson:"dealerDisplayName" json:"dealerDisplayName"`
	TenantID                   string                       `bson:"tenantId" json:"tenantID"`
	TenantDisplayName          string                       `bson:"tenantDisplayName" json:"tenantDisplayName"`
	EPANumber                  string                       `bson:"epaNumber" json:"epaNumber"` // 'ADB 1343857'
	BARNumber                  string                       `bson:"barNumber" json:"barNumber"` // 'CAL00234957'
	Email                      string                       `bson:"email" json:"email"`
	PhoneNumber                string                       `bson:"phoneNumber" json:"phoneNumber"`
	StreetAddress1             string                       `bson:"streetAddress1" json:"streetAddress1"`
	City                       string                       `bson:"city" json:"city"`
	State                      string                       `bson:"state" json:"state"`
	Country                    string                       `bson:"country" json:"country"`
	PostalCode                 string                       `bson:"postalCode" json:"postalCode"`
	Website                    string                       `bson:"website" json:"website"`
	VehicleDamageID            []string                     `bson:"vehicleDamage" json:"vehicleDamage"` //Note: Stores ID's of all vehicle Damages serviced by dealer. Multiple dealers can support same vehicle damage, so for improved fetch of vehicle damage, DealerMaster holds this array.
	TimeZone                   string                       `bson:"timeZone" json:"timeZone"`           //Used for time conversions.
	Currency                   string                       `bson:"currency" json:"currency"`
	Logo                       string                       `bson:"logo" json:"logo"`
	WorkingDaysAndHours        string                       `bson:"workingDaysAndHours" json:"workingDaysAndHours"`
	TaxPercentage              float32                      `bson:"taxPercentage" json:"taxPercentage"`
	Disclaimer                 string                       `bson:"disclaimer" json:"disclaimer"`
	SkillSet                   []string                     `bson:"skillSet" json:"skillSet"`
	VehicleComponentInspection []VehicleComponentInspection `bson:"vehicleComponentInspection" json:"vehicleComponentInspection"`
	ServiceGroup               []string                     `bson:"serviceGroup" json:"serviceGroup"`
}

//VehicleComponentInspection - Structure of vehicle component inspection for a dealer
type VehicleComponentInspection struct {
	InspectionName string `bson:"inspectionName" json:"inspectionName"` //'Under Hood Inspection',
	Status         string `bson:"status" json:"status"`                 //'OK', // OK, Need Attn., N/A
}

// Insert : function to insert dealers to DB
func (dealer Dealer) Insert(ctx apiContext.APIContext) error {
	session, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return err
	}
	defer session.Close()
	return session.DB(ctx.Tenant).C(dealerCollectionName).Insert(dealer)
}

//SelectDamageResponse : structure for SelectDamageResponse
type SelectDamageResponse struct {
	VehicleDamage []vehicle.VehicleDamageMaster `json:"vehicleDamage"`
}

//GetDamageTypes : function to get DamageTypes based on dealerID
func GetDamageTypes(ctx apiContext.APIContext, dealerID string) ([]SelectDamageResponse, error) {

	//variable for storing list of dealer
	dealerResult := []Dealer{}

	//final response given to te client
	result := []SelectDamageResponse{}

	session, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return result, err
	}
	defer session.Close()

	err = session.DB(ctx.Tenant).C(dealerCollectionName).Find(bson.M{"_id": dealerID}).All(&dealerResult)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return result, err
	}

	//looping through the dealer list to get vehicle damage
	for _, val := range dealerResult {
		resp := SelectDamageResponse{}
		vehicleDamageResult := []vehicle.VehicleDamageMaster{}

		//query to find list of vehicle damage to be appended in response
		err = session.DB(ctx.Tenant).C(vehicle.VehicleDamageCollectionName).Find(bson.M{"_id": bson.M{"$in": val.VehicleDamageID}}).All(&vehicleDamageResult)
		if err != nil {
			log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
			return []SelectDamageResponse{}, err
		}
		resp.VehicleDamage = vehicleDamageResult
		result = append(result, resp)
	}
	return result, err
}

//GetDealerByID : Function to get dealer by dealer ID
func GetDealerByID(ctx apiContext.APIContext, dealerID string) (Dealer, error) {

	dealer := Dealer{}

	session, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return dealer, err
	}
	defer session.Close()

	//Fetch dealer object based on dealerID passed as agrument
	err = session.DB(ctx.Tenant).C(dealerCollectionName).Find(bson.M{"_id": dealerID}).One(&dealer)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return dealer, err
	}
	//Return dealer detail
	return dealer, err
}
