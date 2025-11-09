package email

import "time"

// EmailTemplate defines model for EmailTemplate.
type EmailTemplate struct {
	// Attachments List of attachments.
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// CreatedAt Date and time when the email template was created. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// From Email address with optional sender name.
	From *string `json:"from,omitempty"`

	// Html HTML content of the email template.
	Html string `json:"html"`

	// Id Unique identifier of the email template.
	Id int64 `json:"id"`

	// ImagePreviewUrl URL of the image preview.
	ImagePreviewUrl *string `json:"imagePreviewUrl,omitempty"`

	// IsHtmlEditable Flag indicating if the HTML content is editable.
	IsHtmlEditable bool `json:"isHtmlEditable"`

	// LandingPageId Unique identifier of the landing page.
	LandingPageId *string `json:"landingPageId,omitempty"`

	// Name Name of the email template.
	Name *string `json:"name,omitempty"`

	// Preheader Preheader of the email template.
	Preheader *string `json:"preheader,omitempty"`

	// ReplyTo Email address to which recipients of the email can reply.
	ReplyTo *string `json:"replyTo,omitempty"`

	// Subject Subject of the email template.
	Subject *string `json:"subject,omitempty"`

	// UpdatedAt Date and time when the email template was last updated. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
