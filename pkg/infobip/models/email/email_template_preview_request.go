package email

// EmailTemplatePreviewRequest defines model for EmailTemplatePreviewRequest.
type EmailTemplatePreviewRequest struct {
	// Placeholders A map of placeholder names and their replacement values.
	Placeholders *map[string]map[string]interface{} `json:"placeholders,omitempty"`
}

// GenerateEmailTemplatePreviewMultipartRequestBody defines body for GenerateEmailTemplatePreview for multipart/form-data ContentType.
type GenerateEmailTemplatePreviewMultipartRequestBody = EmailTemplatePreviewRequest
