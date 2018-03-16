package dealerService

// This file contains all the models related to dealer

import (
	"time"

	"bitbucket.org/tekion/tenums/constants"
)

// Collection names used by dealer service
const (
	serviceID                    = "tdealer"
	dealerCollectionName         = "DealerMaster"
	fixedOperationCollectionName = "FixedOperation"
)

// dealer struct contains dealer details
// swagger:model dealer
type dealer struct {
	// ID dealer unique identifier
	ID string `bson:"_id" json:"dealerID"`
	// Name dealer name
	Name string `bson:"dealerName" json:"dealerName"`
	// MakeCode car manufacturer code
	MakeCode []string `bson:"makeCode" json:"makeCode"`
	// DoingBusinessAsName this may or may not be government registered name for the business
	DoingBusinessAsName string `bson:"dealerDoingBusinessAsName" json:"dealerDoingBusinessAsName"`
	// StateIssuedNumber state issued/registration number
	StateIssuedNumber string `bson:"stateIssuedNumber" json:"stateIssuedNumber"`
	// ManufacturerIssuedNumber issued/registered number
	ManufacturerIssuedNumber string `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber"`
	// Website dealer website URL
	Website string `bson:"website" json:"website"`
	// TimeZone dealer timezone like PST (Pacific standard Time)
	TimeZone string `bson:"timeZone" json:"timeZone"`
	// Currency dealer currency -  DEFAULT 'USD'
	Currency string `bson:"currency" json:"currency"`
	// TenantID tenants unique identifier
	TenantID string `bson:"tenantID" json:"tenantID"`
	// Phone dealer phone contact
	Phone string `bson:"phone" json:"phone"`
	// Logos dealer logos
	Logos []image `bson:"dealerLogos" json:"dealerLogos"`
	// VehicleDamage dealer vehicle damage types
	VehicleDamage []vehicleDamage `bson:"vehicleDamage" json:"vehicleDamage"`
	// DealershipCode
	// A dealership can have one or more dealers in it.(Requested to change to dealerCode.
	// But this is not one to one as dealerID, thats the reason we put it as dealershipCode,
	// This is kind of dealer GroupCode)
	DealershipCode string `bson:"dealershipCode" json:"dealershipCode"`
	// Application code of dealer
	ApplicationCode string `bson:"applicationCode" json:"applicationCode"`
	// Group list of groups dealer is part of
	Group []string `bson:"dealerGroup" json:"dealerGroup"`
	// Address list of dealer addresses
	Address []dealerAddress `bson:"dealerAddress" json:"dealerAddress"`

	// Contact list of dealerContact ids
	Contact []string `bson:"dealerContact" json:"dealerContact"`
	// BannerImages dealer banner image
	BannerImages []image `bson:"bannerImages" json:"bannerImages"`
	// VideoURL dealer video url
	VideoURL string `bson:"videoURL" json:"videoURL"`

	// IsActive is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
	// LastUpdatedByUser data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// LastUpdatedByDisplayName this is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// LastUpdatedDateTime when was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// DocumentVersion to keep track of the changes - DEFAULT 1.0
	DocumentVersion       float32 `bson:"documentVersion" json:"documentVersion"`
	LateAppointemntMins   int     `bson:"lateAppointmentMins" json:"lateAppointmentMins"`
	MissedAppointemntMins int     `bson:"missedAppointmentMins" json:"missedAppointmentMins"`
}

// Embedded objects in dealer -- start

// image struct contains details of the image stored in S3 bucket, stored as slice of embedded objects in dealer struct
// swagger:model image
type image struct {
	// Width of the stored image in pixels
	Width int32 `bson:"width" json:"width"`
	// Height of the stored image in pixels
	Height int32 `bson:"height" json:"height"`
	// Title image title - e.g Dublin logo
	Title string `bson:"title" json:"title"`
	// ImageID image id - unique identifier of the image in S3 bucket
	ImageID string `bson:"imageID" json:"imageID"`
}

// dealerAddress struct contains details of the dealer address, stored as embedded objects in dealer struct
// swagger:model dealerAddress
type dealerAddress struct {
	// ID dealer address unique identifier
	ID string `bson:"dealerAddressID" json:"dealerAddressID"`
	// AddressType dealer address type like service, sales, parts etc
	AddressType constants.DealerOperationType `bson:"addressType" json:"addressType"`
	// StreetAddress1 dealer address
	StreetAddress1 string `bson:"streetAddress1" json:"streetAddress1"`
	// StreetAddress2 dealer street address - additional address field
	StreetAddress2 string `bson:"streetAddress2" json:"streetAddress2"`
	// City dealer location city
	City string `bson:"city" json:"city"`
	// State dealer Location state
	State string `bson:"state" json:"state"`
	// ZipCode dealer zip code or postal code
	ZipCode string `bson:"zipCode" json:"zipCode"`
	// Country dealer country
	Country string `bson:"country" json:"country"`
	// County dealer location county
	County string `bson:"county" json:"county"`
	// ISDCode extension according to dealer location
	ISDCode string `bson:"isdCode" json:"isdCode"`
	// LocationURL google maps url of the dealer location
	LocationURL string `bson:"locationUrl" json:"locationUrl"`
	// IsActive is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// vehicleDamage struct contains details of the dealer vehicle damage types,
// stored as slice of embedded objects in dealer struct
// swagger:model vehicleDamage
type vehicleDamage struct {
	// ID vehicle damage unique identifier
	ID string `bson:"vehicleDamageID" json:"vehicleDamageID"`
	// ImageURL url of the vehicle damage image
	ImageURL string `bson:"imageURL" json:"imageURL"`
	// DamageType vehicle damage type like scratch, dent, chip etc
	DamageType string `bson:"damageType" json:"damageType"`
	// Description of damage type
	Description string `bson:"description" json:"description"`
	// Priority decides the sequence in which damage types will be displayed on UI
	Priority int `bson:"priority" json:"priority"`
	// IsActive is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// Embedded objects in Dealer -- end

// When get appointment hrs check it its a holiday or not
// date is stored in utc
// date should be converted to dealer time zone

// fixedOperation struct contains dealer fixed operation details
// swagger:model fixedOperation
type fixedOperation struct {
	// ID fixed operation unique identifier
	ID string `bson:"_id" json:"fixedOperationID"`
	// DealerID dealers unique identifier
	DealerID string `bson:"dealerID" json:"dealerID"`
	// EPANumber Environmental Protection Agency Number
	EPANumber string `bson:"EPANumber" json:"EPANumber"`
	// BARNumber Bureau of Automotive Repair Number
	BARNumber string `bson:"BARNumber" json:"BARNumber"`
	//pdi user info
	PDIDetail pdiDetail `bson:"PDIDetail" json:"PDIDetail"`
	//dealerShip working hours weekly
	WorkingHours string `bson:"workingHours" json:"workingHours"`
	//enable sent welcome message for customer portal
	EnableCustomerPortal bool `bson:"enableCustomerPortal" json:"enableCustomerPortal"`
	//flags for dealerTire , mimic ro status update
	Flags map[string]bool `bson:"flags" json:"flags"`
	// DefaultOperationCodes dealers default operation codes
	DefaultOperationCodes []string `bson:"defaultOperationCodes" json:"defaultOperationCodes"`
	// RecommendedOperationCodes dealers recommended operation codes
	RecommendedOperationCodes []string `bson:"recommendedOperationCodes" json:"recommendedOperationCodes"`
	// TaxPercentage fixed operation tax percentage
	TaxPercentage float64 `bson:"taxPercentage" json:"taxPercentage"`
	// DoorRates fixed operation door rates
	DoorRates []doorRate `bson:"doorRates" json:"doorRates"`

	// Disclaimer dealers disclaimer message
	Disclaimer     string   `bson:"disclaimer" json:"disclaimer"`
	DefaultPrinter string   `bson:"defaultPrinter" json:"defaultPrinter"`
	ConcernType    []string `bson:"concernType" json:"concernType"`

	ServiceMenuDisclaimer string `bson:"serviceMenuDisclaimer" json:"serviceMenuDisclaimer"`
	PrinterEmail          string `bson:"printerEmail" json:"printerEmail"`
	// List of printer types and their email addresses
	Printers Printers `bson:"printers" json:"printers"`

	CustomConcernOpcode string   `bson:"customConcernOpcode" json:"customConcernOpcode"`
	RecallOpCodeMapping string   `bson:"recallOpCodeMapping" json:"recallOpCodeMapping"`
	PayTypes            payTypes `bson:"payTypes" json:"payTypes"`

	ApplicationURLs map[string]string `bson:"applicationURLs" json:"applicationURLs"`

	// IsActive is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
	// LastUpdatedByUser data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// LastUpdatedByDisplayName this is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// LastUpdatedDateTime when was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// DocumentVersion to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion"`
}

type pdiDetail struct {
	//pdi customer id of customer master
	CustomerID string `bson:"customerID" json:"customerID"`
	//default opcode for PDI
	DefaultOpcodes []string `bson:"defaultOperationCodes" json:"defaultOperationCodes"`
}

type payTypes map[string]map[string]payType
type payType struct {
	ID               string     `bson:"id" json:"payTypeID"`
	Code             string     `bson:"code" json:"code"`
	Description      string     `bson:"description" json:"description"`
	DefaultLaborType laborType  `bson:"defaultLaborType" json:"defaultLaborType"`
	LaborTypes       laborTypes `bson:"laborTypes" json:"laborTypes"`
}

type laborTypes []laborType
type laborType struct {
	ID          string `bson:"id" json:"laborTypeID"`
	Code        string `bson:"code" json:"code"`
	Description string `bson:"description" json:"description"`
}

// model for printer email addresses
// swagger:model Printers
type Printers struct {
	PartsPullSheet string `bson:"partsPullSheet" json:"partsPullSheet"`
	CheckInSummary string `bson:"checkInSummary" json:"checkInSummary"`
}

// Embedded structures in fixed operations-- start
// doorRate struct
// swagger:model doorRate
type doorRate struct {
	StartDate   *time.Time `bson:"startDate" json:"startDate"`
	EndDate     *time.Time `bson:"endDate" json:"endDate"`
	CustomerPay *float64   `bson:"customerPay" json:"customerPay"`
	Warranty    *float64   `bson:"warranty" json:"warranty"`
	Internal    *float64   `bson:"internal" json:"internal"`
}

// Embedded structures in fixed operations-- end

// swagger:model listDealersReq
type listDealersReq struct {
	IDs            []string `json:"dealerIDs"`
	SelectedFields []string `json:"selectedFields"`
	SortBy         string   `json:"SortBy"`
	Limit          int      `json:"limit"`
}
type readDealerAndFixedOpRes struct {
	Dealer         *dealer         `json:"dealer"`
	FixedOperation *fixedOperation `json:"fixedOperation"`
}
