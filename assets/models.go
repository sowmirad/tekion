package assets

type assets struct {
	ID string `bson:"_id" json:"-"`

	Source    []string `bson:"sources" json:"sources,omitempty"`
	OEMs      []string `bson:"OEMs" json:"OEMs,omitempty"`
	DealerIDs []string `bson:"dealerIDs" json:"dealerIDs,omitempty"`

	Scheduling     []asset `bson:"scheduling" json:"scheduling,omitempty"`
	Appointment    []asset `bson:"appointment" json:"appointment,omitempty"`
	Estimate       []asset `bson:"estimate" json:"estimate,omitempty"`
	CustomerPortal []asset `bson:"customerPortal" json:"customerPortal,omitempty"`
	CDMSWeb        []asset `bson:"cdmsWeb" json:"cdmsWeb,omitempty"`
	CDMSMobile     []asset `bson:"cdmsMobile" json:"cdmsMobile,omitempty"`
}

// swagger:asset image
type asset struct {
	// Width of the stored image in pixels
	Width int32 `bson:"width" json:"width"`
	// Height of the stored image in pixels
	Height int32 `bson:"height" json:"height"`
	// Title image title - e.g Dublin logo
	Title string `bson:"title" json:"title"`
	// ImageID image id - unique identifier of the image in S3 bucket
	ImageID string `bson:"imageID" json:"imageID"`
}

type assetsReqBody struct {
	Source    []string `json:"sources"`
	OEMs      []string `json:"OEMs"`
	DealerIDs []string `json:"dealerIDs"`
	Modules   []string `json:"modules"`
}
