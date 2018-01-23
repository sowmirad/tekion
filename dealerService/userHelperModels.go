package dealerService

import "bitbucket.org/tekion/tbaas/tapi"

/********************** user response models - starts *************************/

// swagger:model userByUserNameRes
type userByUserNameRes struct {
	Meta tapi.MetaData `json:"meta"`
	Data user          `json:"data"`
}

// userByIDRes stores user by id response
// swagger:model userByUserNameRes
type userByIDRes struct {
	Meta tapi.MetaData `json:"meta"`
	Data user          `json:"data"`
}

/********************** user response models - ends ***************************/

/**************************************** user helper models - ends *****************************************/

// userDetails - user information
// swagger:model userDetails
type userDetails struct {
	Name        string `json:"username"`
	DisplayName string `json:"displayName"`
}

/**************************************** user helper models - ends *****************************************/

/**************************************** user helper models - ends *****************************************/

// userDetails - user information
// swagger:model userDetails
type user struct {
	ID               string `bson:"_id" json:"id"`
	TenantID         string `bson:"tenantID" json:"tenantID"`
	TenantName       string `bson:"tenantName" json:"tenantName"`
	HomeDealerID     string `bson:"homeDealerID" json:"homeDealerID"`
	Fname            string `bson:"fname" json:"fname"`
	Lname            string `bson:"lname" json:"lname"`
	Username         string `bson:"username" json:"username"`
	DisplayName      string `bson:"displayName" json:"displayName"`
	Phone            string `bson:"phone" json:"phone"`
	Mobile           string `bson:"mobile" json:"mobile"`
	Email            string `bson:"email" json:"email"`
	ProfilePicture   string `bson:"profilePicture" json:"profilePicture,omitempty"`
	JobTitle         string `bson:"jobTitle" json:"jobTitle"`
	ExternalSourceID string `bson:"externalSourceId" json:"externalSourceId"`
}

/**************************************** user helper models - ends *****************************************/
