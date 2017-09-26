package dealerService

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/log"
	"bitbucket.org/tekion/tbaas/mongoManager"
)

// metaData of HTTP API response
type metaData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// apiResponse complete structure of  HTTP response Meta + Data
type apiResponse struct {
	Meta metaData        `json:"meta"`
	Data json.RawMessage `json:"data,omitempty"`
}

//Global vars
var (
	testToken     = "TestToken"
	validClientID = "mobile"
	validContext  = apiContext.APIContext{Tenant: validTenantName, DealerID: validDealerID, ClientID: validClientID}
)

// Invalid vars
var (
	invalidTenantName       = "InvalidTenant"
	invalidDealerID         = bson.NewObjectId().String()
	invalidClientID         = bson.NewObjectId().String()
	invalidFixedOperationID = bson.NewObjectId().String()
	invalidContactID        = bson.NewObjectId().String()
	invalidGoalID           = bson.NewObjectId().String()
	invalidTenantContext    = apiContext.APIContext{Tenant: invalidTenantName, DealerID: validDealerID, ClientID: validClientID}
	invalidDealerIDContext  = apiContext.APIContext{Tenant: validTenantName, DealerID: invalidDealerID, ClientID: validClientID}
	invalidClintIDContext   = apiContext.APIContext{Tenant: validTenantName, DealerID: validDealerID, ClientID: invalidClientID}
)

func setHeaders(req *http.Request) {
	req.Header.Set("tenantname", validTenantName) // TODO : should be changed to Tenant-Name
	req.Header.Set("dealerid", validDealerID)     // TODO : should be changed to Dealer-ID
	req.Header.Set("tekion-api-token", testToken) // TODO : should be changed to Tekion-API-Token
	req.Header.Set("clientid", "mobile")          // TODO : should be changed to Client-ID

}

func setHeadersAndContext(req *http.Request) {
	setHeaders(req)
	context.Set(req, "apiContext", validContext)
}

func setHeadersAndInvalidTenantContext(req *http.Request) {
	setHeaders(req)
	context.Set(req, "apiContext", invalidTenantContext)
}

func setHeadersAndInvalidDealerIDContext(req *http.Request) {
	setHeaders(req)
	context.Set(req, "apiContext", invalidDealerIDContext)
}

func setHeadersAndInvalidClientIDContext(req *http.Request) {
	setHeaders(req)
	context.Set(req, "apiContext", invalidClintIDContext)
}

//TODO : figure out how to test for lastupdatedtime
//Dealer vars
var (
	//testTime                          = time.Now()
	validDealerID                       = bson.NewObjectId().String()
	validTenantName                     = "Buck"
	validDealerName                     = "Valid Dealer Name"
	validDealerMakeCode                 = []string{"MakeCode1", "MakeCode2", "MakeCode3"}
	validDealerDoingBusinessAsName      = "Valid Business Name"
	validDealerStateIssuedNumber        = "12345"
	validDealerManufacturerIssuedNumber = "56789"
	validDealerWebsite                  = "https://valid-website.com"
	validDealerTimezone                 = "US/Pacific"
	validDealerCurrency                 = "USD"
	validDealerTenantID                 = "99"
	validDealerPhone                    = "+123456789"
	validDealerLogos                    = []image{
		{
			16,
			16,
			"Icon",
			"S3UUIDIcon_1_3",
		},
		{
			48,
			48,
			"Thumb",
			"S3UUIDThumb_1_3",
		},
		{
			256,
			256,
			"Original",
			"S3UUIDOriginal_1_3",
		},
	}
	validDealerVehicleDamage = []vehicleDamage{
		{
			"UUID_1",
			"https://s3-us-west-1.amazonaws.com/cdms-vehicle-damage-images/Icon-Scratch%403x.png",
			"Scratch",
			"Quisque id justo sit amet sapien dignissim vestibulum.",
			1,
		},
		{
			"UUID_2",
			"https://s3-us-west-1.amazonaws.com/cdms-vehicle-damage-images/Icon-Dent%403x.png",
			"Dent",
			"Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum. In hac habitasse platea dictumst. Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante.",
			2,
		},
		{
			"UUID_3",
			"https://s3-us-west-1.amazonaws.com/cdms-vehicle-damage-images/Icon-Chipped%403x.png",
			"Chip",
			"Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci.",
			3,
		},
	}
	validDealerDealershipCode = "987654321"
	validDealerGroup          = []string{"UUID_1", "UUID_2"}
	validDealerAddress        = []dealerAddress{
		{
			"UUID_1",
			"Service",
			"Fremont Street No. 999",
			"",
			"Fremont",
			"CA",
			"2365482",
			"US",
			"",
			true,
		},
		{
			"UUID_2",
			"Parts",
			"New York Street No. 420",
			"some_additional_address",
			"New York",
			"NY",
			"782373",
			"US",
			"some_county",
			true,
		},
	}
	validDealerDocumentTemplates = []dealerDocumentTemplate{
		{
			"UUID_1",
			"Appointment 1",
			"Appointment",
			"S3ImageID_123",
			true,
		},
		{
			"UUID_2",
			"Estimate 1",
			"Estimate",
			"S3ImageID_124",
			true,
		},
		{
			"UUID_3",
			"Repair Order 1",
			"Repair Order",
			"S3ImageID_125",
			false,
		},
	}
	validDealerOperationSchedule = []dealerOperationSchedule{{
		"UUID_1",
		"Sales",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
	},
		{
			"UUID_2",
			"Parts",
			"8:00 AM",
			"8:00 PM",
			"8:00 AM",
			"8:00 PM",
			"8:00 AM",
			"8:00 PM",
			"8:00 AM",
			"8:00 PM",
			"8:00 AM",
			"8:00 PM",
			"8:00 AM",
			"8:00 PM",
			"8:00 AM",
			"8:00 PM",
		},
	}
	validDealerContact                  = []string{"UUID_1", "UUID_2", "UUID_3"}
	validDealerIsActive                 = true
	validDealerLastUpdatedByUser        = "Test User"
	validDealerLastUpdatedByDisplayName = "tdealerTest"
	//validDealerLastUpdatedDateTime    = testTime
	validDealerDocumentVersion = float32(1.0)

	dealerFieldsSlice = []string{"dealerName", "makeCode", "dealerDoingBusinessAsName", "stateIssuedNumber", "manufacturerIssuedNumber", "website", "timeZone", "currency", "tenantID", "phone", "dealerLogos", "vehicleDamage", "dealershipCode", "dealerGroup", "dealerAddress", "dealerDocumentTemplates", "dealerOperationSchedule", "dealerContact"}
	dealerFields      = "dealerName,makeCode,dealerDoingBusinessAsName,stateIssuedNumber,manufacturerIssuedNumber,website,timeZone,currency,tenantID,phone,dealerLogos,vehicleDamage,dealershipCode,dealerGroup,dealerAddress,dealerDocumentTemplates,dealerOperationSchedule,dealerContact"
)

//Dealer objects
var (
	validDealer = dealer{
		ID:                       validDealerID,
		Name:                     validDealerName,
		MakeCode:                 validDealerMakeCode,
		DoingBusinessAsName:      validDealerDoingBusinessAsName,
		StateIssuedNumber:        validDealerStateIssuedNumber,
		ManufacturerIssuedNumber: validDealerManufacturerIssuedNumber,
		Website:                  validDealerWebsite,
		TimeZone:                 validDealerTimezone,
		Currency:                 validDealerCurrency,
		TenantID:                 validDealerTenantID,
		Phone:                    validDealerPhone,
		Logos:                    validDealerLogos,
		VehicleDamage:            validDealerVehicleDamage,
		DealershipCode:           validDealerDealershipCode,
		Group:                    validDealerGroup,
		Address:                  validDealerAddress,
		DocumentTemplates:        validDealerDocumentTemplates,
		OperationSchedule:        validDealerOperationSchedule,
		Contact:                  validDealerContact,
		IsActive:                 validDealerIsActive,
		LastUpdatedByUser:        validDealerLastUpdatedByUser,
		LastUpdatedByDisplayName: validDealerLastUpdatedByDisplayName,
		//LastUpdatedDateTime:    validDealerLastUpdatedDateTime,
		DocumentVersion: validDealerDocumentVersion,
	}

	validDealerWithFields = dealer{
		ID:                       validDealerID,
		Name:                     validDealerName,
		MakeCode:                 validDealerMakeCode,
		DoingBusinessAsName:      validDealerDoingBusinessAsName,
		StateIssuedNumber:        validDealerStateIssuedNumber,
		ManufacturerIssuedNumber: validDealerManufacturerIssuedNumber,
		Website:                  validDealerWebsite,
		TimeZone:                 validDealerTimezone,
		Currency:                 validDealerCurrency,
		TenantID:                 validDealerTenantID,
		Phone:                    validDealerPhone,
		Logos:                    validDealerLogos,
		VehicleDamage:            validDealerVehicleDamage,
		DealershipCode:           validDealerDealershipCode,
		Group:                    validDealerGroup,
		Address:                  validDealerAddress,
		DocumentTemplates:        validDealerDocumentTemplates,
		OperationSchedule:        validDealerOperationSchedule,
		Contact:                  validDealerContact,
	}
)

func dealerDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerCollectionName()).Insert(validDealer)
	if err != nil {
		log.Error(err)
	}
	err = session.DB(validContext.Tenant).C(getDealerCollectionName()).Insert(validDealer)
	if err != nil {
		log.Error(err)
	}
}

func clearDealerDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerCollectionName()).Remove(bson.M{"_id": validDealerID})
	if err != nil {
		log.Error(err)
	}
}

//FixedOperation vars
var (
	validFixedOperationID                = bson.NewObjectId().String()
	validFixedOperationDealerID          = validDealerID
	validFixedOperationEPANumber         = "1223"
	validFixedOperationBARNumber         = "12345"
	validFixedOperationManufacturerLogos = []image{
		{
			16,
			16,
			"Icon",
			"S3UUIDIcon_1",
		},
		{
			48,
			48,
			"Thumb",
			"S3UUIDThumb_1",
		},
		{
			256,
			256,
			"Original",
			"S3UUIDOriginal_1",
		},
	}
	validFixedOperationHolidays = []holiday{
		{
			"25Dec",
			"7:00AM",
			"10:00AM",
			true,
		},
		{
			"4Jul",
			"7:00AM",
			"10:00AM",
			false,
		},
	}
	validFixedOperationServiceAdvisors = []users{
		{
			validDealerID,
			"6",
			"ServiceAdvisor",
		},
		{
			validDealerID,
			"10",
			"ServiceAdvisor",
		},
		{
			validDealerID,
			"12",
			"ServiceAdvisor",
		},
	}
	validFixedOperationFloorCapacity = []floorCapacity{
		{
			"UUID_1",
			"Denting",
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
			"40",
		},
		{
			"UUID_2",
			"Painting",
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
			"40",
		},
	}
	validFixedOperationAppointmentHour = appointmentHour{
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
		"7:00 AM",
		"7:00 PM",
	}
	validFixedOperationAppointmentCapacity = []appointmentCapacity{
		{
			"UUID_3",
			"Express Lane",
			2,
			2,
			15,
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
		},
		{
			"UUID_4",
			"Recall",
			1,
			1,
			30,
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
			"8",
		},
	}
	validFixedOperationAmenities = []amenities{
		{
			"UUID_1",
			"Towing Service",
		},
		{
			"UUID_2",
			"Car Wash",
		},
	}
	validFixedOperationIsActive                 = true
	validFixedOperationLastUpdatedByUser        = "Test User"
	validFixedOperationLastUpdatedByDisplayName = "tdealerTest"
	//validFixedOperationLastUpdatedDateTime    = testTime
	validFixedOperationDocumentVersion = float32(1.0)

	fixedOperationFieldsSlice = []string{"dealerID", "EPANumber", "BARNumber", "manufacturerLogos", "holidays", "serviceAdvisors", "floorCapacity", "appointmentHour", "appointmentCapacity", "amenities"}
	fixedOperationFields      = "dealerID,EPANumber,BARNumber,manufacturerLogos,holidays,serviceAdvisors,floorCapacity,appointmentHour,appointmentCapacity,amenities"
)

//FixedOperation objects
var (
	validFixedOperation = fixedOperation{
		ID:                       validFixedOperationID,
		DealerID:                 validFixedOperationDealerID,
		EPANumber:                validFixedOperationEPANumber,
		BARNumber:                validFixedOperationBARNumber,
		ManufacturerLogos:        validFixedOperationManufacturerLogos,
		Holidays:                 validFixedOperationHolidays,
		ServiceAdvisors:          validFixedOperationServiceAdvisors,
		FloorCapacity:            validFixedOperationFloorCapacity,
		AppointmentHour:          validFixedOperationAppointmentHour,
		AppointmentCapacity:      validFixedOperationAppointmentCapacity,
		Amenities:                validFixedOperationAmenities,
		IsActive:                 validFixedOperationIsActive,
		LastUpdatedByUser:        validFixedOperationLastUpdatedByUser,
		LastUpdatedByDisplayName: validFixedOperationLastUpdatedByDisplayName,
		//LastUpdatedDateTime:	  validFixedOperationLastUpdatedDateTime,
		DocumentVersion: validFixedOperationDocumentVersion,
	}

	validFixedOperationWithFields = fixedOperation{
		ID:                  validFixedOperationID,
		DealerID:            validFixedOperationDealerID,
		EPANumber:           validFixedOperationEPANumber,
		BARNumber:           validFixedOperationBARNumber,
		ManufacturerLogos:   validFixedOperationManufacturerLogos,
		Holidays:            validFixedOperationHolidays,
		ServiceAdvisors:     validFixedOperationServiceAdvisors,
		FloorCapacity:       validFixedOperationFloorCapacity,
		AppointmentHour:     validFixedOperationAppointmentHour,
		AppointmentCapacity: validFixedOperationAppointmentCapacity,
		Amenities:           validFixedOperationAmenities,
	}
)

func fixedOperationDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getFixedOperationCollectionName()).Insert(validFixedOperation)
	if err != nil {
		log.Error(err)
	}
}
func clearFixedOperationDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getFixedOperationCollectionName()).Remove(bson.M{"_id": validFixedOperationID})
	if err != nil {
		log.Error(err)
	}
}

//DealerContact vars
var (
	validContactID                       = bson.NewObjectId().String()
	validContactDealerID                 = validDealerID
	validContactDealerOperationType      = "Sales"
	validContactUser                     = "test.contact@tekion.com"
	validContactUserDisplayName          = "Test"
	validContactUserDisplayTitle         = "Mr."
	validContactLastUpdatedByUser        = "Test User"
	validContactLastUpdatedByDisplayName = "tdealerTest"
	validContactDocumentVersion          = float32(1.0)

	contactFieldsSlice = []string{"dealerID", "dealerOperationType", "user", "userDisplayName", "userDisplayTitle"}
	contactFields      = "dealerID,dealerOperationType,user,userDisplayName,userDisplayTitle"
)

//DealerContact objects
var (
	validContact = dealerContact{
		ID:                       validContactID,
		DealerID:                 validContactDealerID,
		DealerOperationType:      validContactDealerOperationType,
		User:                     validContactUser,
		UserDisplayName:          validContactUserDisplayName,
		UserDisplayTitle:         validContactUserDisplayTitle,
		LastUpdatedByUser:        validContactLastUpdatedByUser,
		LastUpdatedByDisplayName: validContactLastUpdatedByDisplayName,
		DocumentVersion:          validContactDocumentVersion,
	}

	validContacts = []dealerContact{validContact}

	validContactWithFields = dealerContact{
		ID:                  validContactID,
		DealerID:            validContactDealerID,
		DealerOperationType: validContactDealerOperationType,
		User:                validContactUser,
		UserDisplayName:     validContactUserDisplayName,
		UserDisplayTitle:    validContactUserDisplayTitle,
	}

	validContactsWithFields = []dealerContact{validContactWithFields}
)

func contactDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(Coll()).Insert(validContact)
	if err != nil {
		log.Error(err)
	}
}
func clearContactDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerContactCollectionName()).Remove(bson.M{"_id": validContactID})
	if err != nil {
		log.Error(err)
	}
}

// DealerGoal vars
var (
	validGoalID                             = bson.NewObjectId().String()
	validGoalDealerID                       = validDealerID
	validGoalHoursPerRepairOrderAdvisorGoal = "1.5"
	validGoalTotalHoursAdvisorGoal          = "8"
	validGoalAverageLaborRateAdvisorGoal    = "5"
	validGoalLastUpdatedByUser              = "Test User"
	validGoalLastUpdatedByDisplayName       = "tdealerTest"
	validGoalDocumentVersion                = float32(1.0)
	//validGoalLastUpdatedDateTime          = testTime

	goalFieldsSlice = []string{"dealerID", "hoursPerRepairOrderAdvisorGoal", "totalHoursAdvisorGoal", "averageLaborRateAdvisorGoal"}
	goalFields      = "dealerID,hoursPerRepairOrderAdvisorGoal,totalHoursAdvisorGoal,averageLaborRateAdvisorGoal"
)

// DealerGoal objects
var (
	validGoal = dealerGoal{
		ID:                             validGoalID,
		DealerID:                       validGoalDealerID,
		HoursPerRepairOrderAdvisorGoal: validGoalHoursPerRepairOrderAdvisorGoal,
		TotalHoursAdvisorGoal:          validGoalTotalHoursAdvisorGoal,
		AverageLaborRateAdvisorGoal:    validGoalAverageLaborRateAdvisorGoal,
		LastUpdatedByUser:              validGoalLastUpdatedByUser,
		LastUpdatedByDisplayName:       validGoalLastUpdatedByDisplayName,
		DocumentVersion:                validGoalDocumentVersion,
	}

	validGoals = []dealerGoal{validGoal}

	validGoalWithFields = dealerGoal{
		ID:                             validGoalID,
		DealerID:                       validGoalDealerID,
		HoursPerRepairOrderAdvisorGoal: validGoalHoursPerRepairOrderAdvisorGoal,
		TotalHoursAdvisorGoal:          validGoalTotalHoursAdvisorGoal,
		AverageLaborRateAdvisorGoal:    validGoalAverageLaborRateAdvisorGoal,
	}

	validGoalsWithFields = []dealerGoal{validGoalWithFields}
)

func goalDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerGoalCollectionName()).Insert(validGoal)
	if err != nil {
		log.Error(err)
	}
}

func clearGoalDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerGoalCollectionName()).Remove(bson.M{"_id": validGoalID})
	if err != nil {
		log.Error(err)
	}
}

// DealerGroup vars
var (
	validGroupID                       = bson.NewObjectId().String()
	validGroupName                     = "Test Group"
	validGroupDealers                  = []string{validDealerID, "2", "3"}
	validGroupDesc                     = "Test Group"
	validGroupLastUpdatedByUser        = "Test User"
	validGroupLastUpdatedByDisplayName = "tdealerTest"
	validGroupDocumentVersion          = float32(1.0)

	groupFieldsSlice = []string{"dealerGroupName", "dealers", "description"}
	groupFields      = "dealerGroupName,dealers,description"
)

// DealerGroup objects
var (
	validGroup = dealerGroup{
		ID:                       validGroupID,
		Name:                     validGroupName,
		Dealers:                  validGroupDealers,
		Desc:                     validGroupDesc,
		LastUpdatedByUser:        validGroupLastUpdatedByUser,
		LastUpdatedByDisplayName: validGroupLastUpdatedByDisplayName,
		DocumentVersion:          validGroupDocumentVersion,
	}

	validGroups = []dealerGroup{validGroup}

	validGroupWithFields = dealerGroup{
		ID:      validGroupID,
		Name:    validGroupName,
		Dealers: validGroupDealers,
		Desc:    validGroupDesc,
	}

	validGroupsWithFields = []dealerGroup{validGroupWithFields}
)

func groupDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerGroupCollectionName()).Insert(validGroup)
	if err != nil {
		log.Error(err)
	}
}

func clearGroupDataSetup() {
	session, err := mongoManager.GetS(validContext.Tenant)
	if err != nil {
		log.Error(err)
	}
	defer session.Close()
	err = session.DB(validContext.Tenant).C(getDealerGroupCollectionName()).Remove(bson.M{"_id": validGroupID})
	if err != nil {
		panic(err)
	}
}
