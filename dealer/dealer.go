package dealer

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/log"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
	"gopkg.in/mgo.v2/bson"
	"bitbucket.org/tekion/tvehicle/vehicle"
)

var (
	dealerCollectionName string = "DealerMaster"
)

// Model to store master data about each dealer
type DealerMaster struct {
	Id                string   `bson:"_id" json:"dealerId"`
	DealerName        string   `bson:"dealerName" json:"dealerName"`
	DealerDisplayName string   `bson:"dealerDisplayName" json:"dealerDisplayName"`
	TenantId          string   `bson:"tenantId" json:"tenantId"`
	TenantDisplayName string   `bson:"tenantDisplayName" json:"tenantDisplayName"`
	EPANumber         string   `bson:"epaNumber" json:"epaNumber"` // 'ADB 1343857'
	BARNumber         string   `bson:"barNumber" json:"barNumber"` // 'CAL00234957'
	Email             string   `bson:"email" json:"email"`
	PhoneNumber       string   `bson:"phoneNumber" json:"phoneNumber"`
	StreetAddress1    string   `bson:"streetAddress1" json:"streetAddress1"`
	City              string   `bson:"city" json:"city"`
	State             string   `bson:"state" json:"state"`
	Country           string   `bson:"country" json:"country"`
	PostalCode        string   `bson:"postalCode" json:"postalCode"`
	Website           string   `bson:"website" json:"website"`
	VehicleDamageId   []string `bson:"vehicleDamage" json:"vehicleDamage"` //Note: Stores Id's of all vehicle Damages serviced by dealer. Multiple dealers can support same vehicle damage, so for improved fetch of vehicle damage, DealerMaster holds this array.
	ServicesId        []string `bson:"services" json:"services"`           //Note: Stores Id's of all Services provided by dealer. Multiple dealers can provide same service, so for improved fetch of services provided by a dealer, DealerMaster holds this array.
	TimeZone          string   `bson:"timeZone" json:"timeZone"`           //Used for time conversions.
}

func (dealer DealerMaster) Insert(ctx apiContext.APIContext) error {
	s, e := mMgr.GetS(ctx.Tenant)
	if e != nil {
		return e
	}
	defer s.Close()
	return s.DB(ctx.Tenant).C(dealerCollectionName).Insert(dealer)
}

type SelectDamageResponse struct {
	VehicleDamage []vehicle.VehicleDamageMaster `json:"vehicleDamage"`
}

func GetDamageTypes(ctx apiContext.APIContext, dealerid string) (interface{}, error) {
	dealerResult := []DealerMaster{}
	result := []SelectDamageResponse{}

	session, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.Error("Session error ", err.Error())
		return result, err
	}
	defer session.Close()
	err = session.DB(ctx.Tenant).C(dealerCollectionName).Find(bson.M{"_id": dealerid}).All(&dealerResult)

	for _, val := range dealerResult {
		resp := SelectDamageResponse{}
		vehicleDamageResult := []vehicle.VehicleDamageMaster{}

		//todo: add this query in vehicleDamage
		err2 := session.DB(ctx.Tenant).C(vehicle.VehicleDamageCollectionName).Find(bson.M{"_id": bson.M{"$in": val.VehicleDamageId}}).All(&vehicleDamageResult)
		if err2 != nil {
			log.Error("Query Error  ", err2.Error())
			return []SelectDamageResponse{}, err2
		}
		resp.VehicleDamage = vehicleDamageResult
		result = append(result, resp)
	}
	return result, err
}

// Function to get dealer by dealer ID
func GetDealerById(ctx apiContext.APIContext, dealerId string) (DealerMaster, error) {

	log.Info("GetDealerById dealerId : ", dealerId)
	dealer := DealerMaster{}

	session, err := mMgr.GetS(ctx.Tenant)
	if err != nil {
		log.Error("Session error ", err.Error())
		return dealer, err
	}
	defer session.Close()

	//Fetch timezone for dealer
	err = session.DB(ctx.Tenant).C(dealerCollectionName).Find(bson.M{"_id": dealerId}).One(&dealer)

	log.Warning("GetDealerById dealer ", dealer)

	//Return timezone
	return dealer, nil
}
