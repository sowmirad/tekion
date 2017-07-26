package dealerService

// This file contains all the models related to dealer

import (
	"time"
)

// Collection names used by dealer service
const (
	dealerCollectionName         = "DealerMaster"
	fixedOperationCollectionName = "FixedOperation"
	dealerContactCollectionName  = "DealerContact"
	dealerGoalCollectionName     = "DealerGoal"
	dealerGroupCollectionName    = "DealerGroup"
)

// getDealerCollectionName returns dealer collection name
func getDealerCollectionName() string {
	return dealerCollectionName
}

// getFixedOperationCollectionName returns dealer fixed operation collection name
func getFixedOperationCollectionName() string {
	return fixedOperationCollectionName
}

// getDealerContactCollectionName returns the dealer contact collection name
func getDealerContactCollectionName() string {
	return dealerContactCollectionName
}

// getDealerGoalCollectionName returns the dealer goal collection name
func getDealerGoalCollectionName() string {
	return dealerGoalCollectionName
}

// dealerGroupCollectionName returns the dealer group collection name
func getDealerGroupCollectionName() string {
	return dealerGroupCollectionName
}

// getModuleID returns the module id
func getModuleID() string {
	return "tdealer"
}

/*
Moved EPA and BAR numbers to fixed operations
Unclear at this point, we will handle CustomerCommunicationSource, DealerCommunication, OutGoingEmail, IncomingEmail and Email later
CustomerCommunicationSource string					`bson:"CustomerCommunicationSource" json:"CustomerCommunicationSource"` // Customer Communication Source ? Any example
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
	// Dealer Identification  - Keep it Unique across the tenant
	ID string `bson:"_id" json:"dealerID,omitempty"`
	// Dealer Name
	Name string `bson:"dealerName" json:"dealerName,omitempty"`
	// Make Code - Car manufacturer code
	MakeCode []string `bson:"makeCode" json:"makeCode,omitempty"`
	// Dealer Doing Business As Name. This may or may not be government registered Name for the business
	DoingBusinessAsName string `bson:"dealerDoingBusinessAsName" json:"dealerDoingBusinessAsName,omitempty"`
	// State Government registered or issued Number
	StateIssuedNumber string `bson:"stateIssuedNumber" json:"stateIssuedNumber,omitempty"`
	// Manufacturer (OEM) registered-issued Number
	ManufacturerIssuedNumber string `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber,omitempty"`
	// Dealer website URL
	Website string `bson:"website" json:"website,omitempty"`
	// Dealer's timezone like PST ( Pacific standard Time)
	Timezone string `bson:"timezone" json:"timezone,omitempty"`
	// Dealer Currency --  DEFAULT 'USD'
	Currency string `bson:"currency" json:"currency,omitempty"`
	// Tenant Identification Number
	TenantID string `bson:"tenantID" json:"tenantID,omitempty"`
	// Dealership phone Contact
	Phone string `bson:"phone" json:"phone,omitempty"`
	// Dealer logos
	Logos []image `bson:"dealerLogos" json:"dealerLogos,omitempty"`
	// Dealer vehicle damage types
	VehicleDamage []vehicleDamage `bson:"vehicleDamage" json:"vehicleDamage,omitempty"`
	// Dealership Code
	DealershipCode string `bson:"dealershipCode" json:"dealershipCode,omitempty"` // A dealership can have one or more dealers in it.( Requested to change to dealerCode. But this is not one to one as dealerID, thats the reason we put it as dealershipCode, This is kind of dealer GroupCode)
	// Dealer groups
	Group []string `bson:"dealerGroup" json:"dealerGroup,omitempty"`
	// Dealer addresses
	Address []dealerAddress `bson:"dealerAddress" json:"dealerAddress,omitempty"`
	// Dealer document templates
	DocumentTemplates []dealerDocumentTemplate `bson:"dealerDocumentTemplates" json:"dealerDocumentTemplates,omitempty"`
	// Dealer operation schedule
	OperationSchedule []dealerOperationSchedule `bson:"dealerOperationSchedule" json:"dealerOperationSchedule,omitempty"`
	// Dealer contacts
	Contact []string `bson:"dealerContact" json:"dealerContact,omitempty"`
	// Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
	// Last Updated By User
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime *time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes -- DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// dealerContact
// swagger:model dealerContact
type dealerContact struct {
	ID string `bson:"_id" json:"dealerContactID,omitempty"`
	// Dealer Identification Number
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"`
	// Dealer Operation Type - Fixed Operations-Services, Sales, Parts, Management etc
	DealerOperationType string `bson:"dealerOperationType" json:"dealerOperationType,omitempty"`
	// User Id or Login name ( ex: sig@tekion.com )
	User string `bson:"user" json:"user,omitempty"`
	// Dealer Contact/User Display Name like "Scott Hertler "
	UserDisplayName string `bson:"userDisplayName" json:"userDisplayName,omitempty"`
	// Dealer Contact or User Title like "General Manager" or "Parts Clerk"
	UserDisplayTitle string `bson:"userDisplayTitle" json:"userDisplayTitle,omitempty"`
	// Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
	// Last Updated By User
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime *time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes -- DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// Embedded objects in dealer -- start

// image struct contains details of the image stored in S3 bucket, stored as slice of embedded objects in dealer struct
// swagger:model image
type image struct {
	Width  string `bson:"width" json:"width,omitempty"`
	Height string `bson:"height" json:"height,omitempty"`
	Title  string `bson:"title" json:"title,omitempty"`
	// Image id
	ImageID string `bson:"imageID" json:"imageID,omitempty"`
}

// dealerAddress struct contains details of the dealer address, stored as embedded objects in dealer struct
// swagger:model dealerAddress
type dealerAddress struct {
	// Dealer Identification
	ID string `bson:"dealerAddressID" json:"dealerAddressID,omitempty"`
	// Dealer Address Type like Service, Sales, Parts etc
	AddressType string `bson:"addressType" json:"addressType,omitempty"`
	// Dealer Address1
	StreetAddress1 string `bson:"streetAddress1" json:"streetAddress1,omitempty"`
	// Dealer Street Address2
	StreetAddress2 string `bson:"streetAddress2" json:"streetAddress2,omitempty"`
	// Dealer location City
	City string `bson:"city" json:"city,omitempty"`
	// Dealer Location State
	State string `bson:"state" json:"state,omitempty"`
	// Dealer Zip Code - Postal Code
	ZipCode string `bson:"zipCode" json:"zipCode,omitempty"`
	// Dealer Country
	Country string `bson:"country" json:"country,omitempty"`
	// Dealer Location County
	County string `bson:"county" json:"county,omitempty"`
	// Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
}

// TODO : still unclear where and how to store it

// dealerCommunication struct contains details of the dealer communication
// swagger:model dealerCommunication
type dealerCommunication struct {
	ID string `bson:"dealerCommunicationID" json:"dealerCommunicationID,omitempty"`
	// Dealer Identification  - Keep it Unique across the tenant
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"`
	// Customer Communication Source
	CustomerCommunicationSource string `bson:"customerCommunicationSource" json:"customerCommunicationSource,omitempty"` // Customer Communication Source ? Any example
	// Customer communication Out going email
	OutGoingEmail string `bson:"outGoingEmail" json:"outGoingEmail,omitempty"`
	// Customer communication incoming email
	IncomingEmail string `bson:"incomingEmail" json:"incomingEmail,omitempty"`
	// Dealership phone Contact
	Phone string `bson:"phone" json:"phone,omitempty"`
	// Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
}

// dealerDocumentTemplate struct contains details of the templates specific to the dealer, stored as slice of embedded objects in dealer struct
// swagger:model dealerDocumentTemplate
type dealerDocumentTemplate struct {
	ID string `bson:"dealerDocumentTemplateID" json:"dealerDocumentTemplateID,omitempty"`
	// Template Name
	TemplateName string `bson:"templateName" json:"templateName,omitempty"` // Template Name ( What is the use of this templates)
	// Template Type like Appointment, Estimate, Repair Order, Invoice etc
	TemplateType string `bson:"templateType" json:"templateType,omitempty"`
	// TemplateImageID stored in S3 bucket
	TemplateImageID string `bson:"templateImageID" json:"templateImageID,omitempty"`
	// Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
}

// dealerOperationSchedule struct contains details of the dealer operation schedule, stored as slice of embedded objects in dealer struct
// swagger:model dealerOperationSchedule
type dealerOperationSchedule struct {
	ID string `bson:"dealerOperationScheduleID" json:"dealerOperationScheduleID,omitempty"` //
	// Dealer Operation Type - Fixed Operations-Services, Sales, Parts, Management etc
	DealerOperationType string `bson:"dealerOperationType" json:"dealerOperationType,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	SundayOpenHour string `bson:"sundayOpenHour" json:"sundayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	SundayCloseHour string `bson:"sundayCloseHour" json:"sundayCloseHour,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	MondayOpenHour string `bson:"mondayOpenHour" json:"mondayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	MondayCloseHour string `bson:"mondayCloseHour" json:"mondayCloseHour,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	TuesdayOpenHour string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	TuesdayCloseHour string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	WednesdayOpenHour string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	ThursdayOpenHour string `bson:"thursdayOpenHour" json:"thursdayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	ThursdayCloseHour string `bson:"thursdayCloseHour" json:"thursdayCloseHour,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	FridayOpenHour string `bson:"fridayOpenHour" json:"fridayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	FridayCloseHour string `bson:"fridayCloseHour" json:"fridayCloseHour,omitempty"`
	// Business Open Hour or Office start time today : 7:00 AM
	SaturdayOpenHour string `bson:"saturdayOpenHour" json:"saturdayOpenHour,omitempty"`
	// Business Close Hour or Office end time today: 5:00 PM
	SaturdayCloseHour string `bson:"saturdayCloseHour" json:"saturdayCloseHour,omitempty"`
}

// vehicleDamage struct contains details of the dealer vehicle damage types, stored as slice of embedded objects in dealer struct
// swagger:model vehicleDamage
type vehicleDamage struct {
	ID          string `bson:"vehicleDamageID" json:"vehicleDamageID"`
	ImageURL    string `bson:"imageURL" json:"imageURL"`
	DamageType  string `bson:"damageType" json:"damageType"`
	Description string `bson:"description" json:"description"`
	// Decided the sequence in which damage types will be displayed on UI
	Priority int `bson:"priority" json:"priority"`
}

// Embedded objects in Dealer -- end

// When get appointment hrs check it its a holiday or not
// date is stored in utc
// date should be converted to dealer time zone

// fixedOperation struct contains dealer fixed operation details
// swagger:model fixedOperation
type FixedOperation struct {
	ID                string    `bson:"_id" json:"fixedOperationID,omitempty"`
	DealerID          string    `bson:"dealerID" json:"dealerID,omitempty"`
	EPANumber         string    `bson:"EPANumber" json:"EPANumber,omitempty"`
	BARNumber         string    `bson:"BARNumber" json:"BARNumber,omitempty"`
	ManufacturerLogos []image   `bson:"manufacturerLogos" json:"manufacturerLogos,omitempty"`
	Holidays          []holiday `bson:"holidays" json:"holidays,omitempty"`
	// List of service advisors
	ServiceAdvisors     []users               `bson:"serviceAdvisors" json:"serviceAdvisors,omitempty"` // make sure during insertion that only service advisors are stored in here
	FloorCapacity       []floorCapacity       `bson:"floorCapacity" json:"floorCapacity,omitempty"`
	AppointmentHour     appointmentHour       `bson:"appointmentHour" json:"appointmentHour,omitempty"`
	AppointmentCapacity []appointmentCapacity `bson:"appointmentCapacity" json:"appointmentCapacity,omitempty"`
	Amenities           []amenities           `bson:"amenities" json:"amenities,omitempty"`
	// Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
	// Last Updated By User
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime *time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes -- DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
	//Moved AppointmentsPerSlot, AppointmentSlotDuration to appointmentCapacity
}

// Embedded structures in fixed operations-- start

// TODO : need inputs from Venkat on start end and carryover fields

// holiday struct contains details of holidays, stored as slice of embedded objects in fixed operation struct
// swagger:model holiday
type holiday struct {
	Date               string `bson:"date" json:"date,omitempty"`                             // ask for date format used in Tekion
	StartOperationHour string `bson:"startOperationHour" json:"startOperationHour,omitempty"` // not sure about its use ??
	EndOperationHour   string `bson:"endOperationHour" json:"endOperationHour,omitempty"`     // not sure about its use ??
	CarryOver          bool   `bson:"carryOver" json:"carryOver,omitempty"`                   // not sure about its use ??
}

// users struct contains details of the users detail specific to the dealer, stored as slice of embedded objects in fixed operation struct
// Using generic user type instead of serviceAdvisor type
// swagger:model users
type users struct {
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"` // maps to Dealer.ID
	UserID   string `bson:"userID" json:"userID,omitempty"`     // maps to User.ID
	JobTitle string `bson:"jobTitle" json:"jobTitle,omitempty"` // maps to User.jobTitle
}

// floorCapacity struct contains details of dealer floor capacities, , stored as slice of embedded objects in fixed operation struct
// swagger:model floorCapacity
type floorCapacity struct {
	SkillCode     string `bson:"skillCode" json:"skillCode,omitempty"`         // maps to SkillMaster.ID
	SkillName     string `bson:"skillName" json:"skillName,omitempty"`         // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call to skillMaster
	SundayHours   string `bson:"sundayHours" json:"sundayHours,omitempty"`     //
	MondayHour    string `bson:"mondayHour" json:"mondayHour,omitempty"`       //
	TuesdayHour   string `bson:"tuesdayHour" json:"tuesdayHour,omitempty"`     //
	WednesdayHour string `bson:"wednesdayHour" json:"wednesdayHour,omitempty"` //
	ThursdayHour  string `bson:"thursdayHour" json:"thursdayHour,omitempty"`   //
	FridayHour    string `bson:"fridayHour" json:"fridayHour,omitempty"`       //
	SaturdayHour  string `bson:"saturdayHour" json:"saturdayHour,omitempty"`   //
	Total         string `bson:"total" json:"total"`                           // added ask // Can also be calculated by front end
}

// appointmentHour struct contains details appointment hours, stored as embedded object in fixed operation struct
// swagger:model appointmentHour
type appointmentHour struct {
	SundayOpenHour     string `bson:"sundayOpenHour" json:"sundayOpenHour,omitempty"`         //
	SundayCloseHour    string `bson:"sundayCloseHour" json:"sundayCloseHour,omitempty"`       //
	MondayOpenHour     string `bson:"mondayOpenHour" json:"mondayOpenHour,omitempty"`         //
	MondayCloseHour    string `bson:"mondayCloseHour" json:"mondayCloseHour,omitempty"`       //
	TuesdayOpenHour    string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour,omitempty"`       //
	TuesdayCloseHour   string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour,omitempty"`     //
	WednesdayOpenHour  string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour,omitempty"`   //
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour,omitempty"` //
	ThursdayOpenHour   string `bson:"thursdayOpenHour" json:"thursdayOpenHour,omitempty"`     //
	ThursdayCloseHour  string `bson:"thursdayCloseHour" json:"thursdayCloseHour,omitempty"`   //
	FridayOpenHour     string `bson:"fridayOpenHour" json:"fridayOpenHour,omitempty"`         //
	FridayCloseHour    string `bson:"fridayCloseHour" json:"fridayCloseHour,omitempty"`       //
	SaturdayOpenHour   string `bson:"saturdayOpenHour" json:"saturdayOpenHour,omitempty"`     //
	SaturdayCloseHour  string `bson:"saturdayCloseHour" json:"saturdayCloseHour,omitempty"`   //
}

// appointmentCapacity struct contains details of the number of hours present for each skill, stored as slice of embedded objects in fixed operation struct
// swagger:model appointmentCapacity
type appointmentCapacity struct {
	SkillCode               string `bson:"skillCode" json:"skillCode,omitempty"`                             // maps to SkillMaster._id
	SkillName               string `bson:"skillName" json:"skillName,omitempty"`                             // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call
	NumberOfServiceAdvisors int16  `bson:"numberOfServiceAdvisors" json:"numberOfServiceAdvisors,omitempty"` // data type changed to int16 on Prameet's input
	AppointmentsPerSlot     int16  `bson:"appointmentsPerSlot" json:"appointmentsPerSlot,omitempty"`         // data type changed to int16 on Prameet's input
	AppointmentSlotDuration int16  `bson:"appointmentSlotDuration" json:"appointmentSlotDuration,omitempty"` // data type changed to int16 on Prameet's input
	Sunday                  string `bson:"sunday" json:"sunday,omitempty"`                                   //
	Monday                  string `bson:"monday" json:"monday,omitempty"`                                   //
	Tuesday                 string `bson:"tuesday" json:"tuesday,omitempty"`                                 //
	Wednesday               string `bson:"wednesday" json:"wednesday,omitempty"`                             //
	Thursday                string `bson:"thursday" json:"thursday,omitempty"`                               //
	Friday                  string `bson:"friday" json:"friday,omitempty"`                                   //
	Saturday                string `bson:"saturday" json:"saturday,omitempty"`                               //
}

// amenities struct contains list of dealer amenities, stored as slice of embedded objects in fixed operation struct
// swagger:model amenities
type amenities struct {
	ID   string `bson:"amenityID" json:"amenityID,omitempty"` //maps to AmenitiesMaster._id
	Name string `bson:"name" json:"name,omitempty"`           // maps to AmenitiesMaster.name
}

// Embedded structures in fixed operations-- end

// dealerGoal
// swagger:model dealerGoal
type dealerGoal struct {
	// dealerGoalID
	ID string `bson:"_id" json:"dealerGoalID,omitempty"`
	// dealerID
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"`
	// hoursPerRepairOrderAdvisorGoal
	HoursPerRepairOrderAdvisorGoal string `bson:"hoursPerRepairOrderAdvisorGoal" json:"hoursPerRepairOrderAdvisorGoal,omitempty"`
	// totalHoursAdvisorGoal
	TotalHoursAdvisorGoal string `bson:"totalHoursAdvisorGoal" json:"totalHoursAdvisorGoal,omitempty"`
	// averageLaborRateAdvisorGoal
	AverageLaborRateAdvisorGoal string `bson:"averageLaborRateAdvisorGoal" json:"averageLaborRateAdvisorGoal,omitempty"`
	// Last Updated By User
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes -- DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// dealerGroup
// swagger:model dealerGroup
type dealerGroup struct {
	//dealerGroupID
	ID string `bson:"_id" json:"dealerGroupID,omitempty"`
	// Name of group
	Name string `bson:"dealerGroupName" json:"dealerGroupName,omitempty"`
	// List of dealer id's
	Dealers []string `bson:"dealers" json:"dealers,omitempty"`
	// Description of group
	Desc string `bson:"description" json:"description,omitempty"`
	// Last Updated By User
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	// format: date-time
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes -- DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

//Missing collections in mongo -- start
/*
type SkillMaster struct {
	LastUpdatedByUser        string    `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`               // Last Updated By User
	LastUpdatedByDisplayName string    `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"` // This is to Display the Name in the application
	LastUpdatedDateTime      time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`           // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion          float32   `bson:"documentVersion" json:"documentVersion"`                   // Document version to keep track of the changes -- DEFAULT 1.0
}

type AmenitiesMaster struct {
	LastUpdatedByUser        string    `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`               // Last Updated By User
	LastUpdatedByDisplayName string    `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"` // This is to Display the Name in the application
	LastUpdatedDateTime      time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`           // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion          float32   `bson:"documentVersion" json:"documentVersion"`                   // Document version to keep track of the changes -- DEFAULT 1.0
}

// metaData of HTTP API response
// swagger:model metaData
type metaData struct {
	// code
	Code int `json:"code"`
	// msg
	Msg string `json:"msg"`
}

// dealerContactsRespObj
// swagger:response dealerContactsRespObj
type dealerContactsRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data []DealerContact `json:"data,omitempty"`
}

// dealerGoalRespObj
// swagger:response dealerGoalRespObj
type dealerGoalRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data DealerGoal `json:"data,omitempty"`
}

// dealerGoalsRespObj
// swagger:response dealerGoalsRespObj
type dealerGoalsRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data []DealerGoal `json:"data,omitempty"`
}

// fixedOperationRespObj
// swagger:response fixedOperationRespObj
type fixedOperationRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data FixedOperation `json:"data,omitempty"`
}

// fixedOperationsRespObj
// swagger:response fixedOperationsRespObj
type fixedOperationsRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data []FixedOperation `json:"data,omitempty"`
}

// dealerContactRespObj
// swagger:response dealerContactRespObj
type dealerContactRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data DealerContact `json:"data,omitempty"`
}

// dealerRespObj
// swagger:response dealerRespObj
type dealerRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data Dealer `json:"data,omitempty"`
}

// dealerGroupsRespObj
// swagger:response dealerGroupsRespObj
type dealerGroupsRespObj struct {
	// in: body
	Meta metaData `json:"meta"`
	// in: body
	Data []DealerGroup `json:"data,omitempty"`
}

*/
//Missing collections in mongo -- end
