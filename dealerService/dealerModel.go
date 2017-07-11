package dealerService

// This file contains all the models related to dealer

import (
	"time"
)

// Collection names used by dealer service
const (
	dealerCollectionName               = "DealerMaster"
	dealerFixedOperationCollectionName = "FixedOperation"
	dealerContactCollectionName        = "DealerContact"
	dealerGoalCollectionName           = "DealerGoal"
	dealerGroupCollectionName          = "DealerGroup"
)

// getDealerCollectionName returns dealer collection name
func getDealerCollectionName() string {
	return dealerCollectionName
}

// getDealerFixedOperationCollectionName returns dealer fixed operation collection name
func getDealerFixedOperationCollectionName() string {
	return dealerFixedOperationCollectionName
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

// Dealer struct contains details of the dealer
type Dealer struct {
	ID                       string                    `bson:"_id" json:"dealerID,omitempty"`                                        //Dealer Identification  - Keep it Unique across the tenant
	Name                     string                    `bson:"dealerName" json:"dealerName,omitempty"`                               //Dealer Name
	MakeCode                 []string                  `bson:"makeCode" json:"makeCode,omitempty"`                                   //Make Code - Car manufacturer code // made slice of string on Prameet's recommendation
	DoingBusinessAsName      string                    `bson:"dealerDoingBusinessAsName" json:"dealerDoingBusinessAsName,omitempty"` //Dealer Doing Business As Name. This may or may not be government registered Name for the business
	StateIssuedNumber        string                    `bson:"stateIssuedNumber" json:"stateIssuedNumber,omitempty"`                 //State Government registered or issued Number
	ManufacturerIssuedNumber string                    `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber,omitempty"`   // Manufacturer (OEM) registered-issued Number.
	Website                  string                    `bson:"website" json:"website,omitempty"`                                     //Dealer website URL
	Timezone                 string                    `bson:"timezone" json:"timezone,omitempty"`                                   // Dealer's timezone like PST ( Pacific standard Time)
	Currency                 string                    `bson:"currency" json:"currency,omitempty"`                                   // Dealer Currency --  DEFAULT 'USD'
	TenantID                 string                    `bson:"tenantID" json:"tenantID,omitempty"`                                   // Tenant Identification Number
	Phone                    string                    `bson:"phone" json:"phone,omitempty"`                                         // Dealership phone Contact
	Logos                    []image                   `bson:"dealerLogos" json:"dealerLogos,omitempty"`                             // Store Logos in slice of image // changed name from dealerLogoURL to dealerLogos
	VehicleDamage            []vehicleDamage           `bson:"vehicleDamage" json:"vehicleDamage,omitempty"`                         //
	DealershipCode           string                    `bson:"dealershipCode" json:"dealershipCode,omitempty"`                       // Dealership Code.  A dealership can have one or more dealers in it.( Requested to change to dealerCode. But this is not one to one as dealerID, thats the reason we put it as dealershipCode, This is kind of dealerGroupCode)
	Group                    []string                  `bson:"dealerGroup" json:"dealerGroup,omitempty"`                             //
	Address                  []dealerAddress           `bson:"dealerAddress" json:"dealerAddress,omitempty"`                         //
	DocumentTemplates        []dealerDocumentTemplate  `bson:"dealerDocumentTemplates" json:"dealerDocumentTemplates,omitempty"`     //
	OperationSchedule        []dealerOperationSchedule `bson:"dealerOperationSchedule" json:"dealerOperationSchedule,omitempty"`     //
	Contact                  []string                  `bson:"dealerContact" json:"dealerContact,omitempty"`                         //
	IsActive                 bool                      `bson:"isActive" json:"isActive,omitempty"`                                   // Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	LastUpdatedByUser        string                    `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`                 // Last Updated By User
	LastUpdatedByDisplayName string                    `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`   // This is to Display the Name in the application
	LastUpdatedDateTime      time.Time                 `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`             // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion          float32                   `bson:"documentVersion" json:"documentVersion,omitempty"`                     // Document version to keep track of the changes -- DEFAULT 1.0
}

// DealerContact struct contains details of the primary contacts to the dealer
type DealerContact struct {
	ID                       string    `bson:"_id" json:"dealerContactID,omitempty"`                               //
	DealerID                 string    `bson:"dealerID" json:"dealerID,omitempty"`                                 // Dealer Identification Number
	DealerOperationType      string    `bson:"dealerOperationType" json:"dealerOperationType,omitempty"`           // Dealer Operation Type - // Fixed Operations-Services, Sales, Parts, Management etc
	User                     string    `bson:"user" json:"user,omitempty"`                                         // User Id or Login name ( ex: sig@tekion.com )
	UserDisplayName          string    `bson:"userDisplayName" json:"userDisplayName,omitempty"`                   // Dealer Contact/User Display Name like "Scott Hertler "
	UserDisplayTitle         string    `bson:"userDisplayTitle" json:"userDisplayTitle,omitempty"`                 // Dealer Contact or User Title like "General Manager" or "Parts Clerk"
	IsActive                 bool      `bson:"isActive" json:"isActive,omitempty"`                                 // Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	LastUpdatedByUser        string    `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`               // Last Updated By User
	LastUpdatedByDisplayName string    `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"` // This is to Display the Name in the application
	LastUpdatedDateTime      time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`           // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion          float32   `bson:"documentVersion" json:"documentVersion,omitempty"`                   // Document version to keep track of the changes -- DEFAULT 1.0
}

// Embedded objects in Dealer -- start

// image struct contains details of the image stored in S3 bucket, stored as slice of embedded objects in dealer struct
type image struct {
	Width   string `bson:"width" json:"width,omitempty"`     //required
	Height  string `bson:"height" json:"height,omitempty"`   //required
	Title   string `bson:"title" json:"title,omitempty"`     //
	ImageID string `bson:"imageID" json:"imageID,omitempty"` //UUID => name of the image stored in S3 bucket
}

// dealerAddress struct contains details of the dealer address, stored as embedded objects in dealer struct
type dealerAddress struct {
	ID             string `bson:"dealerAddressID" json:"dealerAddressID,omitempty"` // Dealer Identification
	AddressType    string `bson:"addressType" json:"addressType,omitempty"`         // Dealer Address Type like Service, Sales, Parts etc
	StreetAddress1 string `bson:"streetAddress1" json:"streetAddress1,omitempty"`   // Dealer Address1
	StreetAddress2 string `bson:"streetAddress2" json:"streetAddress2,omitempty"`   // Dealer Street Address2
	City           string `bson:"city" json:"city,omitempty"`                       // Dealer location City
	State          string `bson:"state" json:"state,omitempty"`                     // Dealer Location State
	ZipCode        string `bson:"zipCode" json:"zipCode,omitempty"`                 // Dealer Zip Code - Postal Code
	Country        string `bson:"country" json:"country,omitempty"`                 // Dealer Country
	County         string `bson:"county" json:"county,omitempty"`                   // Dealer Location County
	IsActive       bool   `bson:"isActive" json:"isActive,omitempty"`               // Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
}

// TODO : still unclear where and how to store it
// dealerCommunication struct contains details of the dealer communication
type dealerCommunication struct {
	ID                          string `bson:"dealerCommunicationID" json:"dealerCommunicationID,omitempty"`             //
	DealerID                    string `bson:"dealerID" json:"dealerID,omitempty"`                                       // Dealer Identification  - Keep it Unique across the tenant
	CustomerCommunicationSource string `bson:"customerCommunicationSource" json:"customerCommunicationSource,omitempty"` // Customer Communication Source ? Any example
	OutGoingEmail               string `bson:"outGoingEmail" json:"outGoingEmail,omitempty"`                             // Customer communication Out going email ?
	IncomingEmail               string `bson:"incomingEmail" json:"incomingEmail,omitempty"`                             // Customer communication incoming email ?
	Phone                       string `bson:"phone" json:"phone,omitempty"`                                             // Dealership phone Contact
	IsActive                    bool   `bson:"isActive" json:"isActive,omitempty"`                                       // Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
}

// dealerDocumentTemplate struct contains details of the templates specific to the dealer, stored as slice of embedded objects in dealer struct
type dealerDocumentTemplate struct {
	ID              string `bson:"dealerDocumentTemplateID" json:"dealerDocumentTemplateID,omitempty"` //
	TemplateName    string `bson:"templateName" json:"templateName,omitempty"`                         // Template Name ( What is the use of this templates)
	TemplateType    string `bson:"templateType" json:"templateType,omitempty"`                         // Template Type like Appointment, Estimate, Repair Order, Invoice etc
	TemplateImageID string `bson:"templateImageID" json:"templateImageID,omitempty"`                   // Changed from Template URL to TemplateImageID stored in S3 bucket
	IsActive        bool   `bson:"isActive" json:"isActive,omitempty"`                                 // Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
}

// dealerOperationSchedule struct contains details of the dealer operation schedule, stored as slice of embedded objects in dealer struct
type dealerOperationSchedule struct {
	ID                  string `bson:"dealerOperationScheduleID" json:"dealerOperationScheduleID,omitempty"` //
	DealerOperationType string `bson:"dealerOperationType" json:"dealerOperationType,omitempty"`             //Dealer Operation Type - // Fixed Operations-Services, Sales, Parts, Management etc
	SundayOpenHour      string `bson:"sundayOpenHour" json:"sundayOpenHour,omitempty"`                       // Business Open Hour or Office start time today : 7:00 AM
	SundayCloseHour     string `bson:"sundayCloseHour" json:"sundayCloseHour,omitempty"`                     // Business Close Hour or Office end time today: 5:00 PM
	MondayOpenHour      string `bson:"mondayOpenHour" json:"mondayOpenHour,omitempty"`                       // Business Open Hour or Office start time today : 7:00 AM
	MondayCloseHour     string `bson:"mondayCloseHour" json:"mondayCloseHour,omitempty"`                     // Business Close Hour or Office end time today: 5:00 PM
	TuesdayOpenHour     string `bson:"tuesdayOpenHour" json:"tuesdayOpenHour,omitempty"`                     // Business Open Hour or Office start time today : 7:00 AM
	TuesdayCloseHour    string `bson:"tuesdayCloseHour" json:"tuesdayCloseHour,omitempty"`                   // Business Close Hour or Office end time today: 5:00 PM
	WednesdayOpenHour   string `bson:"wednesdayOpenHour" json:"wednesdayOpenHour,omitempty"`                 // Business Open Hour or Office start time today : 7:00 AM
	WednesdayCloseHour  string `bson:"wednesdayCloseHour" json:"wednesdayCloseHour,omitempty"`               // Business Close Hour or Office end time today: 5:00 PM
	ThursdayOpenHour    string `bson:"thursdayOpenHour" json:"thursdayOpenHour,omitempty"`                   // Business Open Hour or Office start time today : 7:00 AM
	ThursdayCloseHour   string `bson:"thursdayCloseHour" json:"thursdayCloseHour,omitempty"`                 // Business Close Hour or Office end time today: 5:00 PM
	FridayOpenHour      string `bson:"fridayOpenHour" json:"fridayOpenHour,omitempty"`                       // Business Open Hour or Office start time today : 7:00 AM
	FridayCloseHour     string `bson:"fridayCloseHour" json:"fridayCloseHour,omitempty"`                     // Business Close Hour or Office end time today: 5:00 PM
	SaturdayOpenHour    string `bson:"saturdayOpenHour" json:"saturdayOpenHour,omitempty"`                   // Business Open Hour or Office start time today : 7:00 AM
	SaturdayCloseHour   string `bson:"saturdayCloseHour" json:"saturdayCloseHour,omitempty"`                 // Business Close Hour or Office end time today: 5:00 PM
}

// vehicleDamage struct contains details of the dealer vehicle damage types, stored as slice of embedded objects in dealer struct
type vehicleDamage struct {
	ID          string `bson:"vehicleDamageID" json:"vehicleDamageID"` //
	ImageURL    string `bson:"imageURL" json:"imageURL"`               //
	DamageType  string `bson:"damageType" json:"damageType"`           //
	Description string `bson:"description" json:"description"`         //
	Priority    int    `bson:"priority" json:"priority"`               // Decided the sequence in which damage types will be displayed on UI
}

// Embedded objects in Dealer -- end

// When get appointment hrs check it its a holiday or not
// date is stored in utc
// date should be converted to dealer time zone

// FixedOperation struct contains dealer fixed operation details
type FixedOperation struct {
	ID                       string                `bson:"_id" json:"fixedOperationID,omitempty"`                              //
	DealerID                 string                `bson:"dealerID" json:"dealerID,omitempty"`                                 // Dealer Identification
	EPANumber                string                `bson:"EPANumber" json:"EPANumber,omitempty"`                               // Name correction -- capitalize abbreviations
	BARNumber                string                `bson:"BARNumber" json:"BARNumber,omitempty"`                               // Name correction -- capitalize abbreviations
	ManufacturerLogos        []image               `bson:"manufacturerLogos" json:"manufacturerLogos,omitempty"`               // Store Logos in slice of image // changed name from manufacturerLogo to manufacturerLogos
	Holidays                 []holiday             `bson:"holidays" json:"holidays,omitempty"`                                 //
	ServiceAdvisors          []users               `bson:"serviceAdvisors" json:"serviceAdvisors,omitempty"`                   // make sure during insertion that only service advisors are stored in here
	FloorCapacity            []floorCapacity       `bson:"floorCapacity" json:"floorCapacity,omitempty"`                       //
	AppointmentHour          appointmentHour       `bson:"appointmentHour" json:"appointmentHour,omitempty"`                   //
	AppointmentCapacity      []appointmentCapacity `bson:"appointmentCapacity" json:"appointmentCapacity,omitempty"`           //
	Amenities                []amenities           `bson:"amenities" json:"amenities,omitempty"`                               //
	IsActive                 bool                  `bson:"isActive" json:"isActive,omitempty"`                                 // Is Active T or F (TRUE or FALSE) -- DEFAULT 'T'
	LastUpdatedByUser        string                `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`               // Last Updated By User
	LastUpdatedByDisplayName string                `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"` // This is to Display the Name in the application
	LastUpdatedDateTime      time.Time             `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`           // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion          float32               `bson:"documentVersion" json:"documentVersion,omitempty"`                   // Document version to keep track of the changes -- DEFAULT 1.0
	//Moved to DealerGoal hoursPerRepairOrderAdvisorGoal, totalHoursAdvisorGoal, averageLaborRateAdvisorGoal
	//Moved AppointmentsPerSlot, AppointmentSlotDuration to appointmentCapacity
}

// Embedded structures in fixed operations-- start

// TODO : need inputs from Venkat on start end and carryover fields
// holiday struct contains details of holidays, stored as slice of embedded objects in fixed operation struct
type holiday struct {
	Date               string `bson:"date" json:"date,omitempty"`                             // ask for date format used in Tekion
	StartOperationHour string `bson:"startOperationHour" json:"startOperationHour,omitempty"` // not sure about its use ??
	EndOperationHour   string `bson:"endOperationHour" json:"endOperationHour,omitempty"`     // not sure about its use ??
	CarryOver          bool   `bson:"carryOver" json:"carryOver,omitempty"`                   // not sure about its use ??
}

// users struct contains details of the users detail specific to the dealer, stored as slice of embedded objects in fixed operation struct
// Using generic user type instead of serviceAdvisor type
type users struct {
	DealerID string `bson:"dealerID" json:"dealerID,omitempty"` // maps to Dealer.ID
	UserID   string `bson:"userID" json:"userID,omitempty"`     // maps to User.ID
	JobTitle string `bson:"jobTitle" json:"jobTitle,omitempty"` // maps to User.jobTitle
}

// floorCapacity struct contains details of dealer floor capacities, , stored as slice of embedded objects in fixed operation struct
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
type appointmentCapacity struct {
	SkillCode               string    `bson:"skillCode" json:"skillCode,omitempty"`                             // maps to SkillMaster._id
	SkillName               string    `bson:"skillName" json:"skillName,omitempty"`                             // maps to skillMaster.Name added // Name is rarely going to change, keeping a copy here to avoid extra call
	NumberOfServiceAdvisors string    `bson:"numberOfServiceAdvisors" json:"numberOfServiceAdvisors,omitempty"` //
	AppointmentsPerSlot     string    `bson:"appointmentsPerSlot" json:"appointmentsPerSlot,omitempty"`         //
	AppointmentSlotDuration time.Time `bson:"appointmentSlotDuration" json:"appointmentSlotDuration,omitempty"` // check for data type ?
	Sunday                  string    `bson:"sunday" json:"sunday,omitempty"`                                   //
	Monday                  string    `bson:"monday" json:"monday,omitempty"`                                   //
	Tuesday                 string    `bson:"tuesday" json:"tuesday,omitempty"`                                 //
	Wednesday               string    `bson:"wednesday" json:"wednesday,omitempty"`                             //
	Thursday                string    `bson:"thursday" json:"thursday,omitempty"`                               //
	Friday                  string    `bson:"friday" json:"friday,omitempty"`                                   //
	Saturday                string    `bson:"saturday" json:"saturday,omitempty"`                               //
}

// amenities struct contains list of dealer amenities, stored as slice of embedded objects in fixed operation struct
type amenities struct {
	ID   string `bson:"amenityID" json:"amenityID,omitempty"` //maps to AmenitiesMaster._id
	Name string `bson:"name" json:"name,omitempty"`           // maps to AmenitiesMaster.name
}

// Embedded structures in fixed operations-- end

// DealerGoal struct contains dealer goals
type DealerGoal struct {
	ID                             string    `bson:"_id" json:"dealerGoalID,omitempty"`                                              //
	DealerID                       string    `bson:"dealerID" json:"dealerID,omitempty"`                                             //
	HoursPerRepairOrderAdvisorGoal string    `bson:"hoursPerRepairOrderAdvisorGoal" json:"hoursPerRepairOrderAdvisorGoal,omitempty"` //
	TotalHoursAdvisorGoal          string    `bson:"totalHoursAdvisorGoal" json:"totalHoursAdvisorGoal,omitempty"`                   //
	AverageLaborRateAdvisorGoal    string    `bson:"averageLaborRateAdvisorGoal" json:"averageLaborRateAdvisorGoal,omitempty"`       //
	LastUpdatedByUser              string    `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`                           // Last Updated By User
	LastUpdatedByDisplayName       string    `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"`             // This is to Display the Name in the application
	LastUpdatedDateTime            time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`                       // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion                float32   `bson:"documentVersion" json:"documentVersion,omitempty"`                               // Document version to keep track of the changes -- DEFAULT 1.0
}

// DealerGroup struct contains groups to which dealer belong
type DealerGroup struct {
	ID                       string    `bson:"_id" json:"dealerGroupID,omitempty"`                                 //
	Name                     string    `bson:"dealerGroupName" json:"dealerGroupName,omitempty"`                   //
	Dealers                  []string  `bson:"dealers" json:"dealers,omitempty"`                                   //
	Desc                     string    `bson:"description" json:"description,omitempty"`                           //
	LastUpdatedByUser        string    `bson:"lastUpdatedByUser" json:"lastUpdatedByUser,omitempty"`               // Last Updated By User
	LastUpdatedByDisplayName string    `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName,omitempty"` // This is to Display the Name in the application
	LastUpdatedDateTime      time.Time `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime,omitempty"`           // When was this last updated Date and Time -- DEFAULT CURRENT_TIMESTAMP
	DocumentVersion          float32   `bson:"documentVersion" json:"documentVersion,omitempty"`                   // Document version to keep track of the changes -- DEFAULT 1.0
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
