package email

// GetEmailTemplatesParams defines parameters for GetEmailTemplates.
type GetEmailTemplatesParams struct {
	// Page Identifies a specific page number from which to retrieve the email template details.
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// Size Identifies the maximum number of email templates returned in the list.
	Size *int32 `form:"size,omitempty" json:"size,omitempty"`
}
