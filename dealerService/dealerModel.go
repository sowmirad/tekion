package dealerService

// This file contains all the models related to dealer

import (
	"time"

	"bitbucket.org/tekion/tbaas/tapi"
	"bitbucket.org/tekion/tenums/constants"
)

// Collection names used by dealer service
const (
	serviceID                    = "tdealer"
	dealerCollectionName         = "DealerMaster"
	fixedOperationCollectionName = "FixedOperation"
	dealerContactCollectionName  = "DealerContact"
	dealerGoalCollectionName     = "DealerGoal"
	dealerGroupCollectionName    = "DealerGroup"
)

/*
Moved EPA and BAR numbers to fixed operations
Unclear at this point, we will handle CustomerCommunicationSource, DealerCommunication, OutGoingEmail,
IncomingEmail and Email later
// Customer Communication Source ? Any example
CustomerCommunicationSource string `bson:"CustomerCommunicationSource" json:"CustomerCommunicationSource"`
DealerCommunication 		[]dealerCommunication 	`bson:"dealerCommunication" json:"dealerCommunication"`
OutGoingEmail 				string 					`bson:"outGoingEmail" json:"outGoingEmail"`
IncomingEmail 				string 					`bson:"incomingEmail" json:"incomingEmail"`
Email                 		string  				`bson:"email" json:"email"`
Not sure about dealerShipCode mapping with dealerGroups
Created a new collection for dealer group
*/

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
	// DocumentTemplates list of dealer document templates
	DocumentTemplates []dealerDocumentTemplate `bson:"dealerDocumentTemplates" json:"dealerDocumentTemplates"`
	// OperationSchedule list of operation hours of different units like sales, parts etc
	OperationSchedule []dealerOperationSchedule `bson:"dealerOperationSchedule" json:"dealerOperationSchedule"`
	// Contact list of dealerContact ids
	Contact []string `bson:"dealerContact" json:"dealerContact"`
	// BannerImages dealer banner image
	BannerImages []image `bson:"bannerImages" json:"bannerImages"`
	// VideoURL dealer video url
	VideoURL string `bson:"videoURL" json:"videoURL"`
	// ServiceConfigs stores service related configs
	ServiceConfigs []serviceConfig `bson:"serviceConfigs" json:"serviceConfigs"`
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

// dealerContact
// swagger:model dealerContact
type dealerContact struct {
	// ID dealer contact unique identifier
	ID string `bson:"_id" json:"dealerContactID"`
	// DealerID dealer identification number
	DealerID string `bson:"dealerID" json:"dealerID"`
	// DealerOperationType - fixed operations , sales, parts, management etc
	DealerOperationType constants.DealerOperationType `bson:"dealerOperationType" json:"dealerOperationType"`
	// User id or login name ( ex: sig@tekion.com )
	User string `bson:"user" json:"user"`
	// UserDisplayName dealer contact/user display name like "Scott Hertler"
	UserDisplayName string `bson:"userDisplayName" json:"userDisplayName"`
	// UserDisplayTitle dealer contact or user title like "General Manager" or "Parts Clerk"
	UserDisplayTitle string `bson:"userDisplayTitle" json:"userDisplayTitle"`
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

// TODO : still unclear where and how to store it

// dealerCommunication struct contains details of the dealer communication
// swagger:model dealerCommunication
type dealerCommunication struct {
	// ID dealer communication unique identifier
	ID string `bson:"dealerCommunicationID" json:"dealerCommunicationID"`
	// DealerID dealer unique identifier
	DealerID string `bson:"dealerID" json:"dealerID"`
	// CustomerCommunicationSource ? any example
	CustomerCommunicationSource string `bson:"customerCommunicationSource" json:"customerCommunicationSource"`
	// OutGoingEmail customer communication out going email
	OutGoingEmail string `bson:"outGoingEmail" json:"outGoingEmail"`
	// IncomingEmail customer communication incoming email
	IncomingEmail string `bson:"incomingEmail" json:"incomingEmail"`
	// Phone delaer phone contact
	Phone string `bson:"phone" json:"phone"`
	// IsActive is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// dealerDocumentTemplate struct contains details of the templates specific to the dealer,
// stored as slice of embedded objects in dealer struct
// swagger:model dealerDocumentTemplate
type dealerDocumentTemplate struct {
	// ID dealer document template unique identifier
	ID string `bson:"dealerDocumentTemplateID" json:"dealerDocumentTemplateID"`
	// TemplateName dealer document template name
	TemplateName string `bson:"templateName" json:"templateName"`
	// TemplateType dealer document template type like appointment, estimate, repair order, invoice etc
	TemplateType constants.DealerDocumentTemplateType `bson:"templateType" json:"templateType"`
	// TemplateImageID unique identifier of dealer document template image stored in S3 bucket
	TemplateImageID string `bson:"templateImageID" json:"templateImageID"`
	// IsActive is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// dealerOperationSchedule struct contains details of the dealer operation schedule,
// stored as slice of embedded objects in dealer struct
// swagger:model dealerOperationSchedule
type dealerOperationSchedule struct {
	// ID dealer operation schedule unique identifier
	ID string `bson:"dealerOperationScheduleID" json:"dealerOperationScheduleID"` //
	// DealerOperationType like service, sales, parts etc
	DealerOperationType constants.DealerOperationType `bson:"dealerOperationType" json:"dealerOperationType"`
	// SundayOpenHour business open hour or office start time today : 7:00, 24 hr format
	SundayOpenHour string `bson:"sundayOpenHour" json:"sundayOpenHour"`
	// SundayCloseHour business close hour or office end time today: 16:00, 24 hr format
	SundayCloseHour string `bson:"sundayCloseHour" json:"sundayCloseHour"`
	// MondayOpenHour business open hour or office start time today : 7:00, 24 hr format
	MondayOpenHour string `bson:"mondayOpenHour" json:"mondayOpenHour"`
	// MondayCloseHour business close hour or office end time today: 16:00, 24 hr format
	MondayCloseHour string `bson:"mondayCloseHour" json:"mondayCloseHour"`
	// TuesdayOpenHour business open hour or office start time today : 7:00, 24 hr format
	TuesdayOpenHour string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour"`
	// TuesdayCloseHour business close hour or office end time today: 16:00, 24 hr format
	TuesdayCloseHour string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour"`
	// WednesdayOpenHour business open hour or office start time today : 7:00, 24 hr format
	WednesdayOpenHour string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour"`
	// WednesdayCloseHour business close hour or office end time today: 16:00, 24 hr format
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour"`
	// ThursdayOpenHour business open hour or office start time today : 7:00, 24 hr format
	ThursdayOpenHour string `bson:"thursdayOpenHour" json:"thursdayOpenHour"`
	// ThursdayCloseHour business close hour or office end time today: 16:00, 24 hr format
	ThursdayCloseHour string `bson:"thursdayCloseHour" json:"thursdayCloseHour"`
	// FridayOpenHour business open hour or office start time today : 7:00, 24 hr format
	FridayOpenHour string `bson:"fridayOpenHour" json:"fridayOpenHour"`
	// FridayCloseHour business close hour or office end time today: 16:00, 24 hr format
	FridayCloseHour string `bson:"fridayCloseHour" json:"fridayCloseHour"`
	// SaturdayOpenHour business open hour or office start time today : 7:00, 24 hr format
	SaturdayOpenHour string `bson:"saturdayOpenHour" json:"saturdayOpenHour"`
	// SaturdayCloseHour business close hour or office end time today: 16:00, 24 hr format
	SaturdayCloseHour string `bson:"saturdayCloseHour" json:"saturdayCloseHour"`
}

// micro services will add their dealer related configurations here
// swagger:model serviceConfig
type serviceConfig struct {
	// ID is service id like tscheduling, tcheckin etc
	ID string `bson:"id" json:"serviceID"`
	// Config is an interface you can store what ever you want here
	Config interface{} `bson:"config" json:"config"`
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
	//enable sent welcome message for customer portal
	EnableCustomerPortal bool `bson:"enableCustomerPortal" json:"enableCustomerPortal"`
	// ManufacturerLogos list of manufacturer logos
	ManufacturerLogos []image `bson:"manufacturerLogos" json:"manufacturerLogos"`
	// Holidays list of holidays
	Holidays []holiday `bson:"holidays" json:"holidays"`
	// ServiceAdvisors list of service advisors
	ServiceAdvisors []users `bson:"serviceAdvisors" json:"serviceAdvisors"`
	// FloorCapacity list of floor capacities
	FloorCapacity []floorCapacity `bson:"floorCapacity" json:"floorCapacity"`
	// AppointmentHour fixed operation appointment hrs
	AppointmentHour appointmentHour `bson:"appointmentHour" json:"appointmentHour"`
	// AppointmentCapacity list of fixed operation appointment capacities
	AppointmentCapacity []appointmentCapacity `bson:"appointmentCapacity" json:"appointmentCapacity"`
	// DefaultOperationCodes dealers default operation codes
	DefaultOperationCodes []string `bson:"defaultOperationCodes" json:"defaultOperationCodes"`
	// RecommendedOperationCodes dealers recommended operation codes
	RecommendedOperationCodes []string `bson:"recommendedOperationCodes" json:"recommendedOperationCodes"`
	// Amenities list of amenities provided by dealer
	Amenities []amenities `bson:"amenities" json:"amenities"`
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

	CustomConcernOpcode string `bson:"customConcernOpcode" json:"customConcernOpcode"`
	RecallOpCodeMapping string `bson:"recallOpCodeMapping" json:"recallOpCodeMapping"`

	PayTypeMapping payTypeMapping `bson:"payTypeMapping" json:"payTypeMapping"`
	PayTypes       payTypes       `bson:"payTypes" json:"payTypes"`

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

type payTypes map[string]payType
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

// model for default payType codes
// swagger:model payTypeMapping
type payTypeMapping struct {
	CustomerPay string `bson:"CP" json:"CP"`
	InternalPay string `bson:"I" json:"I"`
	WarrantyPay string `bson:"W" json:"W"`
}

// Embedded structures in fixed operations-- start

// TODO : need inputs from Venkat on start end and carryover fields

// holiday struct contains details of holidays, stored as slice of embedded objects in fixed operation struct
// swagger:model holiday
type holiday struct {
	// Date of the holiday
	Date *time.Time `bson:"date" json:"date"` // ask for date format used in Tekion
	// StartOperationHour Operation hrs start time in 24 hr format - "7:00"
	StartOperationHour string `bson:"startOperationHour" json:"startOperationHour"` // not sure about its use ??
	// EndOperationHour Operation hrs end time in 24 hr format - "17:00"
	EndOperationHour string `bson:"endOperationHour" json:"endOperationHour"` // not sure about its use ??
	// CarryOver Not sure what it means ?
	CarryOver bool `bson:"carryOver" json:"carryOver"` // not sure about its use ??
}

// users struct contains details of the users detail specific to the dealer,
// stored as slice of embedded objects in fixed operation struct
// Using generic user type instead of serviceAdvisor type
// swagger:model users
type users struct {
	// DealerID dealers unique identifier
	DealerID string `bson:"dealerID" json:"dealerID"` // maps to Dealer.ID
	// UserID users unique identifier
	UserID string `bson:"userID" json:"userID"` // maps to User.ID
	// JobTitle users job title
	JobTitle string `bson:"jobTitle" json:"jobTitle"` // maps to User.jobTitle
}

// floorCapacity struct contains details of dealer floor capacities,
// stored as slice of embedded objects in fixed operation struct
// swagger:model floorCapacity
type floorCapacity struct {
	// SkillCode maps to SkillMaster.ID
	SkillCode string `bson:"skillCode" json:"skillCode"`
	// SkillName name of skill, maps to skillMaster.Name
	// rarely going to change, keeping a copy here to avoid extra call to skillMaster
	SkillName string `bson:"skillName" json:"skillName"`
	// SundayHours skill hrs available on sunday
	SundayHours string `bson:"sundayHours" json:"sundayHours"`
	// MondayHour skill hrs available on monday
	MondayHour string `bson:"mondayHour" json:"mondayHour"`
	// TuesdayHour skill hrs available on tuesday
	TuesdayHour string `bson:"tuesdayHour" json:"tuesdayHour"`
	// WednesdayHour skill hrs available on wednesday
	WednesdayHour string `bson:"wednesdayHour" json:"wednesdayHour"`
	// ThursdayHour skill hrs available on thursday
	ThursdayHour string `bson:"thursdayHour" json:"thursdayHour"`
	// FridayHour skill hrs available on friday
	FridayHour string `bson:"fridayHour" json:"fridayHour"`
	// SaturdayHour skill hrs available on saturday
	SaturdayHour string `bson:"saturdayHour" json:"saturdayHour"`
	// Total weekly hrs available
	Total string `bson:"total" json:"total"` // added - Can also be calculated by front end
}

// appointmentHour struct contains details appointment hours, stored as embedded object in fixed operation struct
// swagger:model appointmentHour
type appointmentHour struct {
	SundayOpenHour     string `bson:"sundayOpenHour" json:"sundayOpenHour"`
	SundayCloseHour    string `bson:"sundayCloseHour" json:"sundayCloseHour"`
	MondayOpenHour     string `bson:"mondayOpenHour" json:"mondayOpenHour"`
	MondayCloseHour    string `bson:"mondayCloseHour" json:"mondayCloseHour"`
	TuesdayOpenHour    string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour"`
	TuesdayCloseHour   string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour"`
	WednesdayOpenHour  string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour"`
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour"`
	ThursdayOpenHour   string `bson:"thursdayOpenHour" json:"thursdayOpenHour"`
	ThursdayCloseHour  string `bson:"thursdayCloseHour" json:"thursdayCloseHour"`
	FridayOpenHour     string `bson:"fridayOpenHour" json:"fridayOpenHour"`
	FridayCloseHour    string `bson:"fridayCloseHour" json:"fridayCloseHour"`
	SaturdayOpenHour   string `bson:"saturdayOpenHour" json:"saturdayOpenHour"`
	SaturdayCloseHour  string `bson:"saturdayCloseHour" json:"saturdayCloseHour"`
}

// appointmentCapacity struct contains details of the number of hours present for each skill,
// stored as slice of embedded objects in fixed operation struct
// swagger:model appointmentCapacity
type appointmentCapacity struct {
	// SkillCode maps to SkillMaster.ID
	SkillCode string `bson:"skillCode" json:"skillCode"`
	// SkillName name of skill, maps to skillMaster.Name
	// rarely going to change, keeping a copy here to avoid extra call to skillMaster
	SkillName string `bson:"skillName" json:"skillName"`
	// NumberOfServiceAdvisors number of service advisors assigned
	NumberOfServiceAdvisors int16 `bson:"numberOfServiceAdvisors" json:"numberOfServiceAdvisors"`
	// AppointmentsPerSlot number of appointments per slot
	AppointmentsPerSlot int16 `bson:"appointmentsPerSlot" json:"appointmentsPerSlot"`
	// AppointmentSlotDuration appointment slot duration
	AppointmentSlotDuration int16 `bson:"appointmentSlotDuration" json:"appointmentSlotDuration"`
	// Sunday available hrs on sunday
	Sunday string `bson:"sunday" json:"sunday"`
	// Monday available hrs on monday
	Monday string `bson:"monday" json:"monday"`
	// Tuesday available hrs on tuesday
	Tuesday string `bson:"tuesday" json:"tuesday"`
	// Wednesday available hrs on wednesday
	Wednesday string `bson:"wednesday" json:"wednesday"`
	// Thursday available hrs on thursday
	Thursday string `bson:"thursday" json:"thursday"`
	// Friday available hrs on friday
	Friday string `bson:"friday" json:"friday"`
	// Saturday available hrs on saturday
	Saturday string `bson:"saturday" json:"saturday"`
}

// amenities struct contains list of dealer amenities, stored as slice of embedded objects in fixed operation struct
// swagger:model amenities
type amenities struct {
	// ID maps to AmenitiesMaster._id
	ID string `bson:"amenityID" json:"amenityID"`
	// Name maps to AmenitiesMaster.name
	Name    string `bson:"name" json:"name"`
	IconURL string `bson:"iconURL" json:"iconURL"`
}

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

// dealerGoal
// swagger:model dealerGoal
type dealerGoal struct {
	// ID goal unique identifier
	ID string `bson:"_id" json:"dealerGoalID"`
	// DealerID dealer unique identifier
	DealerID string `bson:"dealerID" json:"dealerID"`
	// HoursPerRepairOrder time to be spent by service advisor per RO
	HoursPerRepairOrder string `bson:"hoursPerRepairOrder" json:"hoursPerRepairOrder"`
	// TotalHours total time spent by service advisor
	TotalHours string `bson:"totalHours" json:"totalHours"`
	// AverageLaborRate average labor rate for service advisor
	AverageLaborRate string `bson:"averageLaborRate" json:"averageLaborRate"`
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

// dealerGroup
// swagger:model dealerGroup
type dealerGroup struct {
	// ID unique identifier of dealer group
	ID string `bson:"_id" json:"groupID"`
	// Name of group
	Name string `bson:"groupName" json:"groupName"`
	// Dealers list of dealer ids
	Dealers []string `bson:"dealers" json:"dealers"`
	// Description of group
	Desc string `bson:"description" json:"description"`
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

// swagger:model listDealersReq
type listDealersReq struct {
	IDs            []string `json:"dealerIDs"`
	SelectedFields []string `json:"selectedFields"`
	SortBy         string   `json:"SortBy"`
	Limit          int      `json:"limit"`
}

// swagger:model userDtlsRes
type userDtlsRes struct {
	Meta tapi.MetaData `json:"meta"`
	Data userData      `json:"data"`
}

// this is the response we get from signup user endpoint
// swagger:model userData
type userData struct {
	DisplayName string `json:"displayName"`
}

type readDealerAndFixedOpRes struct {
	Dealer         *dealer         `json:"dealer"`
	FixedOperation *fixedOperation `json:"fixedOperation"`
}
