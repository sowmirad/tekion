package assets

type assets struct {
	ID string `bson:"_id" json:"-"`

	Source    []string `bson:"sources" json:"sources,omitempty"`
	OEMs      []string `bson:"OEMs" json:"OEMs,omitempty"`
	DealerIDs []string `bson:"dealerIDs" json:"dealerIDs,omitempty"`

	Scheduling     map[string]asset `bson:"scheduling" json:"scheduling,omitempty"`
	Appointment    map[string]asset `bson:"appointment" json:"appointment,omitempty"`
	Estimate       map[string]asset `bson:"estimate" json:"estimate,omitempty"`
	CustomerPortal map[string]asset `bson:"customerPortal" json:"customerPortal,omitempty"`
	CDMSWeb        map[string]asset `bson:"cdmsWeb" json:"cdmsWeb,omitempty"`
	CDMSMobile     map[string]asset `bson:"cdmsMobile" json:"cdmsMobile,omitempty"`
}

// swagger:asset image
type asset struct {
	Width    int32  `bson:"width" json:"width"`   // Width of the stored image in pixels
	Height   int32  `bson:"height" json:"height"` // Height of the stored image in pixels
	MIMEType string `bson:"MIMEType" json:"MIMEType"`
	URL      string `bson:"URL" json:"URL"` // S3 bucket url
	FileSize int    `bson:"fileSize" json:"fileSize"`
	IsActive bool   `bson:"isActive" json:"isActive"`
}

type assetsReqBody struct {
	Source    []string `json:"sources"`
	OEMs      []string `json:"OEMs"`
	DealerIDs []string `json:"dealerIDs"`
	Modules   []string `json:"modules"`
}
