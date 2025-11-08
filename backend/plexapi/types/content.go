package types

type ContentBodyStrGuid struct {
	MediaContainer MediaContainerStrGuid `json:"MediaContainer"`
}

type MediaContainerStrGuid struct {
	Identifier string            `json:"identifier"`
	Size       int               `json:"size"`
	TotalSize  int               `json:"totalSize"`
	Offset     int               `json:"offset"`
	Metadata   []MetadataStrGuid `json:"Metadata"`
}

type MetadataStrGuid struct {
	RatingKey        string `json:"ratingKey"`
	Title            string `json:"title"`
	GrandparentTitle string `json:"grandparentTitle"`
	Duration         int    `json:"duration"`
	GUID             string `json:"Guid"`
}

type ContentBody struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
}

type MediaContainer struct {
	Identifier string     `json:"identifier"`
	Size       int        `json:"size"`
	TotalSize  int        `json:"totalSize"`
	Offset     int        `json:"offset"`
	Metadata   []Metadata `json:"Metadata"`
}

type Metadata struct {
	RatingKey        string `json:"ratingKey"`
	Title            string `json:"title"`
	GrandparentTitle string `json:"grandparentTitle"`
	Duration         int    `json:"duration"`
	GUIDstr          string `json:"guid"`
	GUID             []GUID `json:"Guid"`
}

type GUID struct {
	ID string `json:"id"`
}
