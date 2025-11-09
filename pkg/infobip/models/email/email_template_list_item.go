package email

import "time"

// EmailTemplateListItem defines model for EmailTemplateListItem.
type EmailTemplateListItem struct {
	// CreatedAt Date and time when the email template was created. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Id Unique identifier of the email template.
	Id int64 `json:"id"`

	// ImagePreviewUrl Image preview URL
	ImagePreviewUrl *string `json:"imagePreviewUrl,omitempty"`

	// Name Name of the email template.
	Name *string `json:"name,omitempty"`

	// UpdatedAt Date and time when the email template was last updated. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
