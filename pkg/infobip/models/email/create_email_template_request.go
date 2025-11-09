package email

// CreateEmailTemplateRequest defines model for CreateEmailTemplateRequest.
type CreateEmailTemplateRequest struct {
	// Attachments JSON string of attachments to be sent with the email template.
	Attachments *string `json:"attachments,omitempty"`

	// From Email address with optional sender name.
	From *string `json:"from,omitempty"`

	// Html HTML content of the email template.
	Html string `json:"html"`

	// LandingPage The identifier of an opt out landing late to be used and displayed when an end user clicks the unsubscribe link. Create a landing page in your Infobip account and use the ID number. For example, 1_23456. If not present, the default opt out landing page is used.
	LandingPage *string `json:"landingPage,omitempty"`

	// Name Name of the email template.
	Name *string `json:"name,omitempty"`

	// Preheader Preheader of the email template.
	Preheader *string `json:"preheader,omitempty"`

	// ReplyTo Email address to which recipients of the email can reply.
	ReplyTo *string `json:"replyTo,omitempty"`

	// Subject Subject of the email template.
	Subject *string `json:"subject,omitempty"`
}

// CreateEmailTemplateMultipartRequestBody defines body for CreateEmailTemplate for multipart/form-data ContentType.
type CreateEmailTemplateMultipartRequestBody = CreateEmailTemplateRequest

// PatchEmailTemplateMultipartRequestBody defines body for PatchEmailTemplate for multipart/form-data ContentType.
type PatchEmailTemplateMultipartRequestBody = CreateEmailTemplateRequest

// UpdateEmailTemplateMultipartRequestBody defines body for UpdateEmailTemplate for multipart/form-data ContentType.
type UpdateEmailTemplateMultipartRequestBody = CreateEmailTemplateRequest
