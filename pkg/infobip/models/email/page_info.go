package email

// PageInfo defines model for PageInfo.
type PageInfo struct {
	// Page Requested page number.
	Page int32 `json:"page"`

	// Size Requested page size.
	Size int32 `json:"size"`

	// TotalPages The total number of pages of the results matching the requested parameters.
	TotalPages int32 `json:"totalPages"`

	// TotalResults The total number of the results matching the requested parameters.
	TotalResults int64 `json:"totalResults"`
}
