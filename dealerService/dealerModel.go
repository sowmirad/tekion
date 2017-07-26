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
	// Dealer identification - Keep it Unique across the tenant
	ID string `bson:"_id" json:"dealerID,omitempty"`
	// Dealer name
	Name string `bson:"dealerName" json:"dealerName,omitempty"`
	// Make code - Car manufacturer code
	MakeCode []string `bson:"makeCode" json:"makeCode,omitempty"`
	// Dealer doing business as name. This may or may not be government registered Name for the business
	DoingBusinessAsName string `bson:"dealerDoingBusinessAsName" json:"dealerDoingBusinessAsName,omitempty"`
	// State government registered or issued Number
	StateIssuedNumber string `bson:"stateIssuedNumber" json:"stateIssuedNumber,omitempty"`
	// Manufacturer (OEM) registered-issued Number
	ManufacturerIssuedNumber string `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber,omitempty"`
	// Dealer website URL
	Website string `bson:"website" json:"website,omitempty"`
	// Dealer's timezone like PST (Pacific standard Time)
	TimeZone string `bson:"timeZone" json:"timeZone,omitempty"`
	// Dealer Currency -  DEFAULT 'USD'
	Currency string `bson:"currency" json:"currency,omitempty"`
	// Tenant identification number
	TenantID string `bson:"tenantID" json:"tenantID,omitempty"`
	// Dealership phone contact
	Phone string `bson:"phone" json:"phone,omitempty"`
	// Dealer logos
	Logos []image `bson:"dealerLogos" json:"dealerLogos,omitempty"`
	// Dealer vehicle damage types
	VehicleDamage []vehicleDamage `bson:"vehicleDamage" json:"vehicleDamage,omitempty"`
	// Dealership code
	DealershipCode string `bson:"dealershipCode" json:"dealershipCode,omitempty"` // A dealership can have one or more dealers in it.( Requested to change to dealerCode. But this is not one to one as dealerID, thats the reason we put it as dealershipCode, This is kind of dealer GroupCode)
	// Dealer groups
	Group []string `bson:"dealerGroup" json:"dealerGroup,omitempty"`
	// Dealer addresses
	Address []dealerAddress `bson:"dealerAddress" json:"dealerAddress,omitempty"`
	// Dealer document templates
	DocumentTemplates []dealerDocumentTemplate `bson:"dealerDocumentTemplates" json:"dealerDocumentTemplates,omitempty"`
	// Dealer operation schedules
	OperationSchedule []dealerOperationSchedule `bson:"dealerOperationSchedule" json:"dealerOperationSchedule,omitempty"`
	// Dealer contacts
	Contact []string `bson:"dealerContact" json:"dealerContact,omitempty"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// dealerContact
// swagger:model dealerContact
type dealerContact struct {
	// Dealer contact unique identifier
	ID string `bson:"_id" json:"dealerContactID,omitempty"`
	// Dealer identification number
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"`
	// Dealer operation type - Fixed Operations-Services, Sales, Parts, Management etc
	DealerOperationType string `bson:"dealerOperationType" json:"dealerOperationType,omitempty"`
	// User id or login name ( ex: sig@tekion.com )
	User string `bson:"user" json:"user,omitempty"`
	// Dealer contact/user display name like "Scott Hertler"
	UserDisplayName string `bson:"userDisplayName" json:"userDisplayName,omitempty"`
	// Dealer contact or User title like "General Manager" or "Parts Clerk"
	UserDisplayTitle string `bson:"userDisplayTitle" json:"userDisplayTitle,omitempty"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// Embedded objects in dealer -- start

// image struct contains details of the image stored in S3 bucket, stored as slice of embedded objects in dealer struct
// swagger:model image
type image struct {
	// Width of the stored image in pixels
	Width int32 `bson:"width" json:"width,omitempty"`
	// Height of the stored image in pixels
	Height int32 `bson:"height" json:"height,omitempty"`
	// Image title - e.g Dublin logo
	Title string `bson:"title" json:"title,omitempty"`
	// Image id - unique identifier of the image in S3 bucket
	ImageID string `bson:"imageID" json:"imageID,omitempty"`
}

// dealerAddress struct contains details of the dealer address, stored as embedded objects in dealer struct
// swagger:model dealerAddress
type dealerAddress struct {
	// Dealer address unique identifier
	ID string `bson:"dealerAddressID" json:"dealerAddressID,omitempty"`
	// Dealer address type like Service, Sales, Parts etc
	AddressType string `bson:"addressType" json:"addressType,omitempty"`
	// Dealer address
	StreetAddress1 string `bson:"streetAddress1" json:"streetAddress1,omitempty"`
	// Dealer street address - additional address field
	StreetAddress2 string `bson:"streetAddress2" json:"streetAddress2,omitempty"`
	// Dealer location city
	City string `bson:"city" json:"city,omitempty"`
	// Dealer Location state
	State string `bson:"state" json:"state,omitempty"`
	// Dealer zip code or postal code
	ZipCode string `bson:"zipCode" json:"zipCode,omitempty"`
	// Dealer country
	Country string `bson:"country" json:"country,omitempty"`
	// Dealer location county
	County string `bson:"county" json:"county,omitempty"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
}

// TODO : still unclear where and how to store it

// dealerCommunication struct contains details of the dealer communication
// swagger:model dealerCommunication
type dealerCommunication struct {
	// Dealer communication unique identifier
	ID string `bson:"dealerCommunicationID" json:"dealerCommunicationID,omitempty"`
	// Dealer identification
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"`
	// Customer communication source
	CustomerCommunicationSource string `bson:"customerCommunicationSource" json:"customerCommunicationSource,omitempty"` // Customer Communication Source ? Any example
	// Customer communication out going email
	OutGoingEmail string `bson:"outGoingEmail" json:"outGoingEmail,omitempty"`
	// Customer communication incoming email
	IncomingEmail string `bson:"incomingEmail" json:"incomingEmail,omitempty"`
	// Dealership phone contact
	Phone string `bson:"phone" json:"phone,omitempty"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
}

// dealerDocumentTemplate struct contains details of the templates specific to the dealer, stored as slice of embedded objects in dealer struct
// swagger:model dealerDocumentTemplate
type dealerDocumentTemplate struct {
	// Dealer document template unique identifier
	ID string `bson:"dealerDocumentTemplateID" json:"dealerDocumentTemplateID,omitempty"`
	// Template name
	TemplateName string `bson:"templateName" json:"templateName,omitempty"` // Template Name ( What is the use of this templates)
	// Template type like Appointment, Estimate, Repair Order, Invoice etc
	TemplateType string `bson:"templateType" json:"templateType,omitempty"`
	// Unique identifier of template image stored in S3 bucket
	TemplateImageID string `bson:"templateImageID" json:"templateImageID,omitempty"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
}

// dealerOperationSchedule struct contains details of the dealer operation schedule, stored as slice of embedded objects in dealer struct
// swagger:model dealerOperationSchedule
type dealerOperationSchedule struct {
	// Dealer operation schedule unique identifier
	ID string `bson:"dealerOperationScheduleID" json:"dealerOperationScheduleID,omitempty"` //
	// Dealer operation type - Fixed Operations-Services, Sales, Parts, Management etc
	DealerOperationType string `bson:"dealerOperationType" json:"dealerOperationType,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	SundayOpenHour string `bson:"sundayOpenHour" json:"sundayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	SundayCloseHour string `bson:"sundayCloseHour" json:"sundayCloseHour,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	MondayOpenHour string `bson:"mondayOpenHour" json:"mondayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	MondayCloseHour string `bson:"mondayCloseHour" json:"mondayCloseHour,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	TuesdayOpenHour string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	TuesdayCloseHour string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	WednesdayOpenHour string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	ThursdayOpenHour string `bson:"thursdayOpenHour" json:"thursdayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	ThursdayCloseHour string `bson:"thursdayCloseHour" json:"thursdayCloseHour,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	FridayOpenHour string `bson:"fridayOpenHour" json:"fridayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	FridayCloseHour string `bson:"fridayCloseHour" json:"fridayCloseHour,omitempty"`
	// Business open hour or office start time today : 7:00 AM
	SaturdayOpenHour string `bson:"saturdayOpenHour" json:"saturdayOpenHour,omitempty"`
	// Business close hour or office end time today: 5:00 PM
	SaturdayCloseHour string `bson:"saturdayCloseHour" json:"saturdayCloseHour,omitempty"`
}

// vehicleDamage struct contains details of the dealer vehicle damage types, stored as slice of embedded objects in dealer struct
// swagger:model vehicleDamage
type vehicleDamage struct {
	// Vehicle damage unique identifier
	ID string `bson:"vehicleDamageID" json:"vehicleDamageID"`
	// URL of the damage image
	ImageURL string `bson:"imageURL" json:"imageURL"`
	// Damage type like Scratch, Dent, Chip etc
	DamageType string `bson:"damageType" json:"damageType"`
	// Description of damage type
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
type fixedOperation struct {
	// Fixed operation unique identifier
	ID string `bson:"_id" json:"fixedOperationID,omitempty"`
	// Dealer unique identifier
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"`
	// Environmental Protection Agency Number
	EPANumber string `bson:"EPANumber" json:"EPANumber,omitempty"`
	// Bureau of Automotive Repair Number
	BARNumber string `bson:"BARNumber" json:"BARNumber,omitempty"`
	// List of manufacturer logos
	ManufacturerLogos []image `bson:"manufacturerLogos" json:"manufacturerLogos,omitempty"`
	// List of holidays
	Holidays []holiday `bson:"holidays" json:"holidays,omitempty"`
	// List of service advisors
	ServiceAdvisors []users `bson:"serviceAdvisors" json:"serviceAdvisors,omitempty"` // make sure during insertion that only service advisors are stored in here
	// List of floor capacities
	FloorCapacity []floorCapacity `bson:"floorCapacity" json:"floorCapacity,omitempty"`
	// Fixed operation appointment hrs
	AppointmentHour appointmentHour `bson:"appointmentHour" json:"appointmentHour,omitempty"`
	// List of fixed operation appointment capacities
	AppointmentCapacity []appointmentCapacity `bson:"appointmentCapacity" json:"appointmentCapacity,omitempty"`
	// List of amenities provided by dealer
	Amenities []amenities `bson:"amenities" json:"amenities,omitempty"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive,omitempty"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// Embedded structures in fixed operations-- start

// TODO : need inputs from Venkat on start end and carryover fields

// holiday struct contains details of holidays, stored as slice of embedded objects in fixed operation struct
// swagger:model holiday
type holiday struct {
	// Holiday date like 25th Dec
	Date string `bson:"date" json:"date,omitempty"` // ask for date format used in Tekion
	// Operation hrs start - "7:00"
	StartOperationHour string `bson:"startOperationHour" json:"startOperationHour,omitempty"` // not sure about its use ??
	// Operation hrs end - "5:00"
	EndOperationHour string `bson:"endOperationHour" json:"endOperationHour,omitempty"` // not sure about its use ??
	//
	CarryOver bool `bson:"carryOver" json:"carryOver,omitempty"` // not sure about its use ??
}

// users struct contains details of the users detail specific to the dealer, stored as slice of embedded objects in fixed operation struct
// Using generic user type instead of serviceAdvisor type
// swagger:model users
type users struct {
	// Dealer unique identifier
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"` // maps to Dealer.ID
	// User unique identifier
	UserID string `bson:"userID" json:"userID,omitempty"` // maps to User.ID
	// User job title
	JobTitle string `bson:"jobTitle" json:"jobTitle,omitempty"` // maps to User.jobTitle
}

// floorCapacity struct contains details of dealer floor capacities, , stored as slice of embedded objects in fixed operation struct
// swagger:model floorCapacity
type floorCapacity struct {
	// Skill unique identifier
	SkillCode string `bson:"skillCode" json:"skillCode,omitempty"` // maps to SkillMaster.ID
	// Skill name
	SkillName string `bson:"skillName" json:"skillName,omitempty"` // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call to skillMaster
	// Skill hrs available on sunday
	SundayHours string `bson:"sundayHours" json:"sundayHours,omitempty"`
	// Skill hrs available on monday
	MondayHour string `bson:"mondayHour" json:"mondayHour,omitempty"`
	// Skill hrs available on tuesday
	TuesdayHour string `bson:"tuesdayHour" json:"tuesdayHour,omitempty"`
	// Skill hrs available on wednesday
	WednesdayHour string `bson:"wednesdayHour" json:"wednesdayHour,omitempty"`
	// Skill hrs available on thursday
	ThursdayHour string `bson:"thursdayHour" json:"thursdayHour,omitempty"`
	// Skill hrs available on friday
	FridayHour string `bson:"fridayHour" json:"fridayHour,omitempty"`
	// Skill hrs available on saturday
	SaturdayHour string `bson:"saturdayHour" json:"saturdayHour,omitempty"`
	// Weekly hrs available
	Total string `bson:"total" json:"total"` // added - Can also be calculated by front end
}

// appointmentHour struct contains details appointment hours, stored as embedded object in fixed operation struct
// swagger:model appointmentHour
type appointmentHour struct {
	SundayOpenHour     string `bson:"sundayOpenHour" json:"sundayOpenHour,omitempty"`
	SundayCloseHour    string `bson:"sundayCloseHour" json:"sundayCloseHour,omitempty"`
	MondayOpenHour     string `bson:"mondayOpenHour" json:"mondayOpenHour,omitempty"`
	MondayCloseHour    string `bson:"mondayCloseHour" json:"mondayCloseHour,omitempty"`
	TuesdayOpenHour    string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour,omitempty"`
	TuesdayCloseHour   string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour,omitempty"`
	WednesdayOpenHour  string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour,omitempty"`
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour,omitempty"`
	ThursdayOpenHour   string `bson:"thursdayOpenHour" json:"thursdayOpenHour,omitempty"`
	ThursdayCloseHour  string `bson:"thursdayCloseHour" json:"thursdayCloseHour,omitempty"`
	FridayOpenHour     string `bson:"fridayOpenHour" json:"fridayOpenHour,omitempty"`
	FridayCloseHour    string `bson:"fridayCloseHour" json:"fridayCloseHour,omitempty"`
	SaturdayOpenHour   string `bson:"saturdayOpenHour" json:"saturdayOpenHour,omitempty"`
	SaturdayCloseHour  string `bson:"saturdayCloseHour" json:"saturdayCloseHour,omitempty"`
}

// appointmentCapacity struct contains details of the number of hours present for each skill, stored as slice of embedded objects in fixed operation struct
// swagger:model appointmentCapacity
type appointmentCapacity struct {
	// Skill unique identifier
	SkillCode string `bson:"skillCode" json:"skillCode,omitempty"` // maps to SkillMaster._id
	// Skill name
	SkillName string `bson:"skillName" json:"skillName,omitempty"` // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call
	// Number of service advisors assigned
	NumberOfServiceAdvisors int16 `bson:"numberOfServiceAdvisors" json:"numberOfServiceAdvisors,omitempty"` // data type changed to int16 on Prameet's input
	// Number of appointments per slot
	AppointmentsPerSlot int16 `bson:"appointmentsPerSlot" json:"appointmentsPerSlot,omitempty"` // data type changed to int16 on Prameet's input
	// Slot duration
	AppointmentSlotDuration int16 `bson:"appointmentSlotDuration" json:"appointmentSlotDuration,omitempty"` // data type changed to int16 on Prameet's input
	// Available hrs on sunday
	Sunday string `bson:"sunday" json:"sunday,omitempty"`
	// Available hrs on monday
	Monday string `bson:"monday" json:"monday,omitempty"`
	// Available hrs on tuesday
	Tuesday string `bson:"tuesday" json:"tuesday,omitempty"`
	// Available hrs on wednesday
	Wednesday string `bson:"wednesday" json:"wednesday,omitempty"`
	// Available hrs on thursday
	Thursday string `bson:"thursday" json:"thursday,omitempty"`
	// Available hrs on friday
	Friday string `bson:"friday" json:"friday,omitempty"`
	// Available hrs on saturday
	Saturday string `bson:"saturday" json:"saturday,omitempty"`
}

// amenities struct contains list of dealer amenities, stored as slice of embedded objects in fixed operation struct
// swagger:model amenities
type amenities struct {
	// Am
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
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion,omitempty"`
}

// dealerGroup
// swagger:model dealerGroup
type dealerGroup struct {
	// Unique identifier of dealer group
	ID string `bson:"_id" json:"dealerGroupID,omitempty"`
	// Name of group
	Name string `bson:"dealerGroupName" json:"dealerGroupName,omitempty"`
	// List of dealer id's
	Dealers []string `bson:"dealers" json:"dealers,omitempty"`
	// Description of group
	Desc string `bson:"description" json:"description,omitempty"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`
	// When was this last updated Date and Time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`
	// Document version to keep track of the changes - DEFAULT 1.0
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
