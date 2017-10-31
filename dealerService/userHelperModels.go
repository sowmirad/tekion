package dealerService

import "bitbucket.org/tekion/tbaas/tapi"

/********************** user response models - starts *************************/

// swagger:model getUserByUserNameResp
type getUserByUserNameResp struct {
	Meta tapi.MetaData `json:"meta"`
	Data userDetails   `json:"data"`
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
