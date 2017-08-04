package dealerService

// This file contains all the models related to dealer

import (
	"bitbucket.org/tekion/tenums/constants"
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
	ID string `bson:"_id" json:"dealerID"`
	// Dealer name
	Name string `bson:"dealerName" json:"dealerName"`
	// Make code - Car manufacturer code
	MakeCode []string `bson:"makeCode" json:"makeCode"`
	// Dealer doing business as name. This may or may not be government registered Name for the business
	DoingBusinessAsName string `bson:"dealerDoingBusinessAsName" json:"dealerDoingBusinessAsName"`
	// State government registered or issued Number
	StateIssuedNumber string `bson:"stateIssuedNumber" json:"stateIssuedNumber"`
	// Manufacturer (OEM) registered-issued Number
	ManufacturerIssuedNumber string `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber"`
	// Dealer website URL
	Website string `bson:"website" json:"website"`
	// Dealer's timezone like PST (Pacific standard Time)
	TimeZone string `bson:"timeZone" json:"timeZone"`
	// Dealer Currency -  DEFAULT 'USD'
	Currency string `bson:"currency" json:"currency"`
	// Tenant identification number
	TenantID string `bson:"tenantID" json:"tenantID"`
	// Dealership phone contact
	Phone string `bson:"phone" json:"phone"`
	// Dealer logos
	Logos []image `bson:"dealerLogos" json:"dealerLogos"`
	// Dealer vehicle damage types
	VehicleDamage []vehicleDamage `bson:"vehicleDamage" json:"vehicleDamage"`
	// Dealership code
	DealershipCode string `bson:"dealershipCode" json:"dealershipCode"` // A dealership can have one or more dealers in it.( Requested to change to dealerCode. But this is not one to one as dealerID, thats the reason we put it as dealershipCode, This is kind of dealer GroupCode)
	// Dealer groups
	Group []string `bson:"dealerGroup" json:"dealerGroup"`
	// Dealer addresses
	Address []dealerAddress `bson:"dealerAddress" json:"dealerAddress"`
	// Dealer document templates
	DocumentTemplates []dealerDocumentTemplate `bson:"dealerDocumentTemplates" json:"dealerDocumentTemplates"`
	// Dealer operation schedules
	OperationSchedule []dealerOperationSchedule `bson:"dealerOperationSchedule" json:"dealerOperationSchedule"`
	// Dealer contacts
	Contact []string `bson:"dealerContact" json:"dealerContact"`
	// Dealer banner image
	BannerImages []image `bson:"bannerImages" json:"bannerImages"`
	// Dealer video url
	VideoURL string `bson:"videoURL" json:"videoURL"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// This is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// When was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion"`
}

// dealerContact
// swagger:model dealerContact
type dealerContact struct {
	// Dealer contact unique identifier
	ID string `bson:"_id" json:"dealerContactID"`
	// Dealer identification number
	DealerID string `bson:"dealerID" json:"dealerID"`
	// Dealer operation type - Fixed Operations-Services, Sales, Parts, Management etc
	DealerOperationType constants.DealerOperationType `bson:"dealerOperationType" json:"dealerOperationType"`
	// User id or login name ( ex: sig@tekion.com )
	User string `bson:"user" json:"user"`
	// Dealer contact/user display name like "Scott Hertler"
	UserDisplayName string `bson:"userDisplayName" json:"userDisplayName"`
	// Dealer contact or User title like "General Manager" or "Parts Clerk"
	UserDisplayTitle string `bson:"userDisplayTitle" json:"userDisplayTitle"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// This is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// When was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// Document version to keep track of the changes - DEFAULT 1.0
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
	// Image title - e.g Dublin logo
	Title string `bson:"title" json:"title"`
	// Image id - unique identifier of the image in S3 bucket
	ImageID string `bson:"imageID" json:"imageID"`
}

// dealerAddress struct contains details of the dealer address, stored as embedded objects in dealer struct
// swagger:model dealerAddress
type dealerAddress struct {
	// Dealer address unique identifier
	ID string `bson:"dealerAddressID" json:"dealerAddressID"`
	// Dealer address type like Service, Sales, Parts etc
	AddressType constants.DealerOperationType `bson:"addressType" json:"addressType"`
	// Dealer address
	StreetAddress1 string `bson:"streetAddress1" json:"streetAddress1"`
	// Dealer street address - additional address field
	StreetAddress2 string `bson:"streetAddress2" json:"streetAddress2"`
	// Dealer location city
	City string `bson:"city" json:"city"`
	// Dealer Location state
	State string `bson:"state" json:"state"`
	// Dealer zip code or postal code
	ZipCode string `bson:"zipCode" json:"zipCode"`
	// Dealer country
	Country string `bson:"country" json:"country"`
	// Dealer location county
	County string `bson:"county" json:"county"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// TODO : still unclear where and how to store it

// dealerCommunication struct contains details of the dealer communication
// swagger:model dealerCommunication
type dealerCommunication struct {
	// Dealer communication unique identifier
	ID string `bson:"dealerCommunicationID" json:"dealerCommunicationID"`
	// Dealer identification
	DealerID string `bson:"dealerID" json:"dealerID"`
	// Customer communication source
	CustomerCommunicationSource string `bson:"customerCommunicationSource" json:"customerCommunicationSource"` // Customer Communication Source ? Any example
	// Customer communication out going email
	OutGoingEmail string `bson:"outGoingEmail" json:"outGoingEmail"`
	// Customer communication incoming email
	IncomingEmail string `bson:"incomingEmail" json:"incomingEmail"`
	// Dealership phone contact
	Phone string `bson:"phone" json:"phone"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// dealerDocumentTemplate struct contains details of the templates specific to the dealer, stored as slice of embedded objects in dealer struct
// swagger:model dealerDocumentTemplate
type dealerDocumentTemplate struct {
	// Dealer document template unique identifier
	ID string `bson:"dealerDocumentTemplateID" json:"dealerDocumentTemplateID"`
	// Template name
	TemplateName string `bson:"templateName" json:"templateName"` // Template Name ( What is the use of this templates)
	// Template type like Appointment, Estimate, Repair Order, Invoice etc
	TemplateType constants.DealerDocumentTemplateType `bson:"templateType" json:"templateType"`
	// Unique identifier of template image stored in S3 bucket
	TemplateImageID string `bson:"templateImageID" json:"templateImageID"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
}

// dealerOperationSchedule struct contains details of the dealer operation schedule, stored as slice of embedded objects in dealer struct
// swagger:model dealerOperationSchedule
type dealerOperationSchedule struct {
	// Dealer operation schedule unique identifier
	ID string `bson:"dealerOperationScheduleID" json:"dealerOperationScheduleID"` //
	// Dealer operation type - Fixed Operations-Services, Sales, Parts, Management etc
	DealerOperationType constants.DealerOperationType `bson:"dealerOperationType" json:"dealerOperationType"`
	// Business open hour or office start time today : 7:00 AM
	SundayOpenHour string `bson:"sundayOpenHour" json:"sundayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	SundayCloseHour string `bson:"sundayCloseHour" json:"sundayCloseHour"`
	// Business open hour or office start time today : 7:00 AM
	MondayOpenHour string `bson:"mondayOpenHour" json:"mondayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	MondayCloseHour string `bson:"mondayCloseHour" json:"mondayCloseHour"`
	// Business open hour or office start time today : 7:00 AM
	TuesdayOpenHour string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	TuesdayCloseHour string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour"`
	// Business open hour or office start time today : 7:00 AM
	WednesdayOpenHour string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	WednesdayCloseHour string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour"`
	// Business open hour or office start time today : 7:00 AM
	ThursdayOpenHour string `bson:"thursdayOpenHour" json:"thursdayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	ThursdayCloseHour string `bson:"thursdayCloseHour" json:"thursdayCloseHour"`
	// Business open hour or office start time today : 7:00 AM
	FridayOpenHour string `bson:"fridayOpenHour" json:"fridayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	FridayCloseHour string `bson:"fridayCloseHour" json:"fridayCloseHour"`
	// Business open hour or office start time today : 7:00 AM
	SaturdayOpenHour string `bson:"saturdayOpenHour" json:"saturdayOpenHour"`
	// Business close hour or office end time today: 5:00 PM
	SaturdayCloseHour string `bson:"saturdayCloseHour" json:"saturdayCloseHour"`
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
	ID string `bson:"_id" json:"fixedOperationID"`
	// Dealer unique identifier
	DealerID string `bson:"dealerID" json:"dealerID"`
	// Environmental Protection Agency Number
	EPANumber string `bson:"EPANumber" json:"EPANumber"`
	// Bureau of Automotive Repair Number
	BARNumber string `bson:"BARNumber" json:"BARNumber"`
	// List of manufacturer logos
	ManufacturerLogos []image `bson:"manufacturerLogos" json:"manufacturerLogos"`
	// List of holidays
	Holidays []holiday `bson:"holidays" json:"holidays"`
	// List of service advisors
	ServiceAdvisors []users `bson:"serviceAdvisors" json:"serviceAdvisors"` // make sure during insertion that only service advisors are stored in here
	// List of floor capacities
	FloorCapacity []floorCapacity `bson:"floorCapacity" json:"floorCapacity"`
	// Fixed operation appointment hrs
	AppointmentHour appointmentHour `bson:"appointmentHour" json:"appointmentHour"`
	// List of fixed operation appointment capacities
	AppointmentCapacity []appointmentCapacity `bson:"appointmentCapacity" json:"appointmentCapacity"`
	// List of amenities provided by dealer
	Amenities []amenities `bson:"amenities" json:"amenities"`
	// dealer disclaimer message
	Disclaimer string `bson:"disclaimer" json:"disclaimer"`
	// Is active T or F (TRUE or FALSE) -- DEFAULT 'T'
	IsActive bool `bson:"isActive" json:"isActive"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// This is to display the name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// When was this last updated date and time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion"`
}

// Embedded structures in fixed operations-- start

// TODO : need inputs from Venkat on start end and carryover fields

// holiday struct contains details of holidays, stored as slice of embedded objects in fixed operation struct
// swagger:model holiday
type holiday struct {
	// Holiday date like 25th Dec
	Date string `bson:"date" json:"date"` // ask for date format used in Tekion
	// Operation hrs start - "7:00"
	StartOperationHour string `bson:"startOperationHour" json:"startOperationHour"` // not sure about its use ??
	// Operation hrs end - "5:00"
	EndOperationHour string `bson:"endOperationHour" json:"endOperationHour"` // not sure about its use ??
	//
	CarryOver bool `bson:"carryOver" json:"carryOver"` // not sure about its use ??
}

// users struct contains details of the users detail specific to the dealer, stored as slice of embedded objects in fixed operation struct
// Using generic user type instead of serviceAdvisor type
// swagger:model users
type users struct {
	// Dealer unique identifier
	DealerID string `bson:"dealerID" json:"dealerID"` // maps to Dealer.ID
	// User unique identifier
	UserID string `bson:"userID" json:"userID"` // maps to User.ID
	// User job title
	JobTitle string `bson:"jobTitle" json:"jobTitle"` // maps to User.jobTitle
}

// floorCapacity struct contains details of dealer floor capacities, , stored as slice of embedded objects in fixed operation struct
// swagger:model floorCapacity
type floorCapacity struct {
	// Skill unique identifier
	SkillCode string `bson:"skillCode" json:"skillCode"` // maps to SkillMaster.ID
	// Skill name
	SkillName string `bson:"skillName" json:"skillName"` // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call to skillMaster
	// Skill hrs available on sunday
	SundayHours string `bson:"sundayHours" json:"sundayHours"`
	// Skill hrs available on monday
	MondayHour string `bson:"mondayHour" json:"mondayHour"`
	// Skill hrs available on tuesday
	TuesdayHour string `bson:"tuesdayHour" json:"tuesdayHour"`
	// Skill hrs available on wednesday
	WednesdayHour string `bson:"wednesdayHour" json:"wednesdayHour"`
	// Skill hrs available on thursday
	ThursdayHour string `bson:"thursdayHour" json:"thursdayHour"`
	// Skill hrs available on friday
	FridayHour string `bson:"fridayHour" json:"fridayHour"`
	// Skill hrs available on saturday
	SaturdayHour string `bson:"saturdayHour" json:"saturdayHour"`
	// Weekly hrs available
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

// appointmentCapacity struct contains details of the number of hours present for each skill, stored as slice of embedded objects in fixed operation struct
// swagger:model appointmentCapacity
type appointmentCapacity struct {
	// Skill unique identifier
	SkillCode string `bson:"skillCode" json:"skillCode"` // maps to SkillMaster._id
	// Skill name
	SkillName string `bson:"skillName" json:"skillName"` // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call
	// Number of service advisors assigned
	NumberOfServiceAdvisors int16 `bson:"numberOfServiceAdvisors" json:"numberOfServiceAdvisors"` // data type changed to int16 on Prameet's input
	// Number of appointments per slot
	AppointmentsPerSlot int16 `bson:"appointmentsPerSlot" json:"appointmentsPerSlot"` // data type changed to int16 on Prameet's input
	// Slot duration
	AppointmentSlotDuration int16 `bson:"appointmentSlotDuration" json:"appointmentSlotDuration"` // data type changed to int16 on Prameet's input
	// Available hrs on sunday
	Sunday string `bson:"sunday" json:"sunday"`
	// Available hrs on monday
	Monday string `bson:"monday" json:"monday"`
	// Available hrs on tuesday
	Tuesday string `bson:"tuesday" json:"tuesday"`
	// Available hrs on wednesday
	Wednesday string `bson:"wednesday" json:"wednesday"`
	// Available hrs on thursday
	Thursday string `bson:"thursday" json:"thursday"`
	// Available hrs on friday
	Friday string `bson:"friday" json:"friday"`
	// Available hrs on saturday
	Saturday string `bson:"saturday" json:"saturday"`
}

// amenities struct contains list of dealer amenities, stored as slice of embedded objects in fixed operation struct
// swagger:model amenities
type amenities struct {
	ID   string `bson:"amenityID" json:"amenityID"` // maps to AmenitiesMaster._id
	Name string `bson:"name" json:"name"`           // maps to AmenitiesMaster.name
}

// Embedded structures in fixed operations-- end

// dealerGoal
// swagger:model dealerGoal
type dealerGoal struct {
	// dealerGoalID
	ID string `bson:"_id" json:"dealerGoalID"`
	// dealerID
	DealerID string `bson:"dealerID" json:"dealerID"`
	// hoursPerRepairOrderAdvisorGoal
	HoursPerRepairOrderAdvisorGoal string `bson:"hoursPerRepairOrderAdvisorGoal" json:"hoursPerRepairOrderAdvisorGoal"`
	// totalHoursAdvisorGoal
	TotalHoursAdvisorGoal string `bson:"totalHoursAdvisorGoal" json:"totalHoursAdvisorGoal"`
	// averageLaborRateAdvisorGoal
	AverageLaborRateAdvisorGoal string `bson:"averageLaborRateAdvisorGoal" json:"averageLaborRateAdvisorGoal"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// When was this last updated Date and Time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion"`
}

// dealerGroup
// swagger:model dealerGroup
type dealerGroup struct {
	// Unique identifier of dealer group
	ID string `bson:"_id" json:"dealerGroupID"`
	// Name of group
	Name string `bson:"dealerGroupName" json:"dealerGroupName"`
	// List of dealer id's
	Dealers []string `bson:"dealers" json:"dealers"`
	// Description of group
	Desc string `bson:"description" json:"description"`
	// Data updated by who
	LastUpdatedByUser string `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	// This is to Display the Name in the application
	LastUpdatedByDisplayName string `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	// When was this last updated Date and Time - type: datetime - DEFAULT CURRENT_TIMESTAMP
	LastUpdatedDateTime time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	// Document version to keep track of the changes - DEFAULT 1.0
	DocumentVersion float32 `bson:"documentVersion" json:"documentVersion"`
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
*/
//Missing collections in mongo -- end
