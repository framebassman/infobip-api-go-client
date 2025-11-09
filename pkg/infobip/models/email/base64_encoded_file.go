package email

// Base64EncodedFile defines model for Base64EncodedFile.
type Base64EncodedFile struct {
	// ContentType Content type of the file.
	ContentType string `json:"contentType"`

	// Data Base64 encoded file data.
	Data string `json:"data"`

	// FileName Name of the file.
	FileName string `json:"fileName"`
}

// UploadEmailTemplateAttachmentMultipartRequestBody defines body for UploadEmailTemplateAttachment for multipart/form-data ContentType.
type UploadEmailTemplateAttachmentMultipartRequestBody = Base64EncodedFile
